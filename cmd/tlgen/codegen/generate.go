package codegen

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/quenbyako/ext/maps"
	"github.com/quenbyako/ext/slices"
	"github.com/xelaj/tl/schema"
)

func Generate(s *schema.Schema) (string, error) {
	f := jen.NewFile("main")

	isInterface := func(name schema.ObjName) bool {
		_, ok := s.Objects[name]
		return ok
	}

	for _, name := range slices.SortFunc(maps.Keys(s.Objects), func(a, b schema.ObjName) int { return a.Cmp(b) }) {
		f.Add(generateObjects(name, s.Objects[name], isInterface))
	}

	for _, name := range slices.SortFunc(maps.Keys(s.Enums), func(a, b schema.ObjName) int { return a.Cmp(b) }) {
		f.Add(generateEnum(name, s.Enums[name]))
	}

	f.Add(generateRequestBareFunction())

	for _, methodGroup := range slices.Sort(maps.Keys(s.MethodsGroups)) {
		methods := s.MethodsGroups[methodGroup]
		for _, method := range methods {
			obj := generateRequestType(methodGroup, method, isInterface)
			f.Add(obj, jen.Line())
		}
	}

	return fmt.Sprintf("%#v", f), nil
}
