package schema

import "github.com/quenbyako/ext/slices"

type DiffObjects struct {
	Added   []Object
	Removed []ObjName
}

func (d DiffObjects) Patch(a []Object) []Object {
	for _, add := range d.Added {
		a = append(a, add)
	}

	for _, rem := range d.Removed {
		for i, obj := range a {
			if cmpObjName(obj.Name, rem) == 0 {
				a = append(a[:i], a[i+1:]...)
				break
			}
		}
	}

	return slices.SortFunc(a, sortObject)
}

type DiffEnum struct {
	Comment string
	Changes DiffObjects
}

func (a EnumObjects) Diff(b EnumObjects) (res DiffEnum) {
	if a.Comment != b.Comment {
		res.Comment = b.Comment
	}

	panic("unimplemented")
}

func (d DiffEnum) Patch(a EnumObjects) EnumObjects {
	if d.Comment != "" {
		a.Comment = d.Comment
	}

	a.Objects = d.Changes.Patch(a.Objects)

	return a
}
