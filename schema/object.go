package schema

import (
	"fmt"
	"hash/crc32"
)

type Object struct {
	Comment string
	Name    string
	CRC     uint32
	Fields  Parameters
	Type    Type

	isMethod bool
}

func (o *Object) ObjType() ObjectType {
	if o.isMethod {
		return ObjectTypeMethod
	}

	if len(o.Fields) == 0 {
		return ObjectTypeEnum
	}

	return ObjectTypeConstructor
}

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

	return crc32.ChecksumIEEE([]byte(fmt.Sprintf("%v%v = %v;", o.Name, fieldsStr, o.Type)))
}

func (o *Object) String() string {
	fields := ""
	if len(o.Fields) > 0 {
		fields = " " + o.Fields.String()
	}

	return fmt.Sprintf("%v#%08x%v = %v;", o.Name, o.getCRC(), fields, o.Type)
}

func (o *Object) Comments() []string {
	res := []string{}
	if o.Comment != "" {
		res = append(res, "// @"+o.ObjType().String()+" "+o.Comment)
	}

	return append(res, o.Fields.Comments()...)
}
