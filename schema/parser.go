// Copyright (c) 2022 Xelaj Software
//
// This file is a part of tl package.
// See https://github.com/xelaj/tl/blob/master/LICENSE_README.md for details.

package schema

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/alecthomas/participle/v2"
	"github.com/pkg/errors"

	"github.com/xelaj/tl/schema/internal/declaration"
	"github.com/xelaj/tl/schema/internal/lexer"
)

//nolint:gochecknoglobals // obviously parser must be global
var parser = participle.MustBuild(&declaration.Program{},
	participle.Lexer(lexer.NewDefinition()),
)

func ParseFile(filename, content string) (*Schema, error) {
	res := &declaration.Program{}
	err := parser.ParseString(filename, content, res)
	if err != nil {
		return nil, err //nolint:wrapcheck // it's important to keep error
	}

	normalized, err := normalizeProgram(res)
	if err != nil {
		return nil, err
	}

	return normalized, nil
}

func normalizeIdent(i *declaration.Ident) (Type, error) {
	if i.Extension != nil {
		if len(i.Extension.Inner) > 1 {
			return nil, errors.New(i.String() + ": too many modificators")
		}
		if i.Simple.String() != "Vector" {
			return nil, errors.New(i.String() + ": only Vector is allowed to modify")
		}

		return TypeVector(i.Extension.Inner[0].String()), nil
	}

	return TypeCommon(i.Simple.String()), nil
}

func normalizeArgument(arg *declaration.Argument, comment string) (Parameter, error) {
	typ, err := normalizeIdent(&arg.Term)
	if err != nil {
		return nil, errors.Wrap(err, arg.Ident.String())
	}

	if arg.Conditional == nil {
		return RequiredParameter{
			Comment: comment,
			Name:    arg.Ident.String(),
			Type:    typ,
		}, nil
	}

	if v, ok := typ.(TypeCommon); ok && v == "true" {
		return TriggerParameter{
			Comment:     comment,
			Name:        arg.Ident.String(),
			FlagTrigger: arg.Conditional.Ident.Value,
			BitTrigger:  arg.Conditional.Index,
		}, nil
	}

	return OptionalParameter{
		Comment:     comment,
		Name:        arg.Ident.String(),
		Type:        typ,
		FlagTrigger: arg.Conditional.Ident.Value,
		BitTrigger:  arg.Conditional.Index,
	}, nil
}

func normalizeCombinator(
	decl *declaration.CombinatorDecl,
	constructorComment string,
	argsComments map[string]string,
	functionsMode bool,
) (*Object, error) {
	parts := strings.Split(decl.ID, "#") // guaranteed to split by two parts, lexer handles it
	name := parts[0]
	crcStr := parts[1]

	// same: lexer handles everything already
	crc, err := strconv.ParseUint(crcStr, 16, 32) //nolint:gomnd // obvious
	if err != nil {
		panic(err)
	}

	params := make([]Parameter, len(decl.Args))
	for i, arg := range decl.Args {
		arg := arg

		var comment string
		if !arg.Ident.Empty {
			comment = argsComments[arg.Ident.String()]
			delete(argsComments, arg.Ident.String())
		}

		var argErr error
		params[i], argErr = normalizeArgument(&arg, comment)
		if argErr != nil {
			return nil, errors.Wrap(argErr, decl.ID)
		}
	}

	typ, err := normalizeIdent(&declaration.Ident{
		Simple:    declaration.SimpleIdent{Type: &declaration.TypeIdent{Ident: decl.Result.Simple}},
		Extension: decl.Result.Expr,
	})
	if err != nil {
		return nil, errors.Wrap(err, decl.ID+": parsing return type")
	}

	if len(argsComments) != 0 {
		keys := make([]string, 0, len(argsComments))
		for k := range argsComments {
			keys = append(keys, k)
		}
		keysStr := strings.Join(keys, ", ")
		return nil, fmt.Errorf("%v: unknown params in comment tags: %v", decl.ID, keysStr)
	}

	return &Object{
		Comment:  constructorComment,
		Name:     name,
		CRC:      uint32(crc),
		Fields:   params,
		Type:     typ,
		isMethod: functionsMode,
	}, nil
}

const (
	tagType        = "type"
	tagEnum        = "enum"
	tagConstructor = "constructor"
	tagMethod      = "method"
	tagParam       = "param"
)

func normalizeEntries(items []declaration.ProgramEntry, functionsMode bool) ([]*Object, map[string]string, error) {
	var (
		objects      = []*Object{}
		typeComments = map[string]string{}

		currentTypeComment     string
		constructorComment     string
		constructorCommentType string
		argumentComments       = map[string]string{}
	)
	for _, item := range items {
		switch {
		case item.Newline:
			constructorComment = ""
			constructorCommentType = ""
			argumentComments = map[string]string{}

		case item.Comment != nil:
			comment := strings.TrimSpace(strings.TrimPrefix(*item.Comment, "//"))
			var commentTag string
			if strings.HasPrefix(comment, "@") {
				const commentParts = 2

				comment = strings.TrimPrefix(comment, "@")
				parts := strings.SplitN(comment, " ", commentParts)

				commentTag = parts[0]
				if len(parts) >= commentParts {
					comment = strings.TrimSpace(parts[1])
				} else {
					comment = ""
				}
			}

			switch tag := commentTag; tag {
			case "":
				// pass

			case tagType:
				if functionsMode {
					err := fmt.Errorf("@%v %v: impossible on functions", tag, comment)
					return nil, nil, err
				}

				// special case, declaration of type apply to first available type (= SomeType;).
				// when type comment applied to type, currentTypeComment is clearing, so we checking
				// is this type comment declared twice.
				if currentTypeComment != "" {
					err := fmt.Errorf("@%v %v: type comment declared twice", tag, comment)
					return nil, nil, err
				}
				currentTypeComment = comment

			case tagEnum, tagConstructor, tagMethod:
				if functionsMode && tag != tagMethod {
					return nil, nil, errors.New("@" + tag + ": works only in type definitions")
				} else if !functionsMode && comment == tagMethod {
					return nil, nil, errors.New("@" + tag + ": works only in functions definitions")
				}

				if constructorCommentType != "" {
					err := fmt.Errorf("@%v %v: constructor comment declared twice", tag, comment)
					return nil, nil, err
				}
				constructorCommentType = tag
				constructorComment = comment

			case tagParam:
				const commentParts = 2

				parts := strings.SplitN(comment, " ", commentParts)
				realComment := ""
				paramName := parts[0]
				if len(parts) >= commentParts {
					realComment = strings.TrimSpace(parts[1])
				}

				if _, ok := argumentComments[paramName]; ok {
					return nil, nil, errors.New("@param " + paramName + ": comment declared twice")
				}
				if realComment != "" {
					argumentComments[paramName] = realComment
				}

			default:
				return nil, nil, errors.New("@" + commentTag + ": invalid comment tag")
			}

		case item.Decl != nil:
			obj, err := normalizeCombinator(item.Decl, constructorComment, argumentComments, functionsMode)
			if err != nil {
				return nil, nil, err
			}
			objects = append(objects, obj)

			if currentTypeComment != "" {
				v, ok := obj.Type.(TypeCommon)
				if !ok {
					typStr := obj.Type.String()
					err := errors.New("@type: comment set to " + typStr + ", which is impossible")
					return nil, nil, err
				}

				if _, ok := typeComments[string(v)]; ok {
					err := errors.New("@type: for " + string(v) + ", type comment defined twice")
					return nil, nil, err
				}

				typeComments[string(v)] = currentTypeComment
				currentTypeComment = ""
			}

			constructorComment = ""
			constructorCommentType = ""
			argumentComments = map[string]string{}
		}
	}

	return objects, typeComments, nil
}

func normalizeProgram(program *declaration.Program) (*Schema, error) {
	types, comments, err := normalizeEntries(program.Constraints, false)
	if err != nil {
		return nil, err
	}

	methods, _, err := normalizeEntries(program.Methods, true)
	if err != nil {
		return nil, err
	}

	return &Schema{
		Objects: append(types, methods...),
		//	Methods:      methods,
		TypeComments: comments,
	}, nil
}
