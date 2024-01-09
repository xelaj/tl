package schema

import (
	"cmp"
	"fmt"
	"hash/crc32"
	"strings"
)

type Object struct {
	Comment string
	Name    ObjName
	CRC     uint32
	Fields  Parameters
	Type    Type
}

type ObjName struct {
	Group string
	Name  string
}

func objNameFromString(s string) ObjName {
	groupname := strings.Split(s, ".")
	var group string
	name := groupname[0]
	if len(groupname) > 1 {
		group = groupname[0]
		name = groupname[1]
	}

	return ObjName{Group: group, Name: name}
}

func (o ObjName) String() string {
	if o.Group != "" {
		return o.Group + "." + o.Name
	}

	return o.Name
}

func (o ObjName) IsInterface() bool { return isFirstRuneUpper(o.Name) }

func (o ObjName) Cmp(b ObjName) int { return cmpObjName(o, b) }

func cmpObjName(a, b ObjName) int {
	if c := cmp.Compare(a.Group, b.Group); c != 0 {
		return c
	} else if c := cmp.Compare(a.Name, b.Name); c != 0 {
		return c
	} else {
		return 0
	}
}

func sortObject(a, b Object) int { return cmpObjName(a.Name, b.Name) }

type ObjectType uint8

const (
	ObjectTypeUnknown ObjectType = iota
	ObjectTypeConstructor
	ObjectTypeEnum
	ObjectTypeMethod
)

func (o ObjectType) String() string {
	switch o {
	case ObjectTypeConstructor:
		return "constructor"
	case ObjectTypeEnum:
		return "enum"
	case ObjectTypeMethod:
		return "method"
	default:
		return "<UNKNOWN>"
	}
}

// how CRC is calculated:
// first of all, constructor formats to canonical state, e.g.
// `user#12325 id:int         first_name:string    last_name:string = User`
// ↓↓↓
// `user id:int first_name:string last_name:string = User`
//
// all bit triggers are removing, e.g.
// `messages.clearRecentStickers#8999602d flags:# attached:flags.0?true = Bool`
// ↓↓↓
// `messages.clearRecentStickers flags:# = Bool`
//
// For vectors like `getSmthn items:Vector<int> = Bool` i still don't understand
// how to generate, cause it fails in real mtproto schema
func (o *Object) getCRC() uint32 {
	if o.CRC != 0 {
		return o.CRC
	}

	filtered := make(Parameters, 0, len(o.Fields))
	for _, item := range o.Fields {
		if _, ok := item.(TriggerParameter); !ok {
			filtered = append(filtered, item)
		}
	}

	fieldsStr := ""
	if len(filtered) > 0 {
		fieldsStr = " " + filtered.String()
	}

	return crc32.ChecksumIEEE([]byte(fmt.Sprintf("%v%v = %v;", o.Name.String(), fieldsStr, o.Type)))
}

func (o *Object) String() string {
	fields := ""
	if len(o.Fields) > 0 {
		fields = " " + o.Fields.String()
	}

	return fmt.Sprintf("%v#%08x%v = %v;", o.Name.String(), o.getCRC(), fields, o.Type)
}

func (o *Object) Comments(typ ObjectType) []string {
	var res []string
	if o.Comment != "" {
		res = append(res, "// @"+typ.String()+" "+o.Comment)
	}

	return append(res, o.Fields.Comments()...)
}
