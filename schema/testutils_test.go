package schema

func Method(o *Object) *Object {
	o.isMethod = true
	return o
}
