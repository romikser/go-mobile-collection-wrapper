package sliceWrapper

const objectTemplate = `{{range .Types}}
type {{.NameTitle}}Slice struct {
	s []{{.Name}}
}

func New{{.NameTitle}}Slice() *{{.NameTitle}}Slice {
	return &{{.NameTitle}}Slice{}
}

func (v *{{.NameTitle}}Slice) Clear() {
	v.s = v.s[:0]
}

func (v *{{.NameTitle}}Slice) Equal(rhs *{{.NameTitle}}Slice) bool {
	if rhs == nil {
		return false
	}

	if len(v.s) != len(rhs.s) {
		return false
	}

	for i := range v.s {
		if v.s[i] != rhs.s[i] {
			return false
		}
	}

	return true
}

func (v *{{.NameTitle}}Slice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.s)
}

func (v *{{.NameTitle}}Slice) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &v.s)
}

func (v *{{.NameTitle}}Slice) Copy(rhs *{{.NameTitle}}Slice) {
	v.s = make([]{{.Name}}, len(rhs.s))
	copy(v.s, rhs.s)
}

func (v *{{.NameTitle}}Slice) Clone() *{{.NameTitle}}Slice {
	return &{{.NameTitle}}Slice{
		s: v.s[:],
	}
}

func (v *{{.NameTitle}}Slice) Index(rhs {{.Name}}) int {
	for i, lhs := range v.s {
		if lhs == rhs {
			return i
		}
	}
	return -1
}

func (v *{{.NameTitle}}Slice) Append(n {{.Name}}) {
	v.s = append(v.s, n)
}

func (v *{{.NameTitle}}Slice) Insert(i int, n {{.Name}}) {
	if i < 0 || i > len(v.s) {
		fmt.Printf("Vapi::{{.Name}}Slice field_values.go error trying to insert at index %d\n", i)
		return
	}
	v.s = append(v.s, reflect.Zero(reflect.TypeOf(n)).Interface().({{.Name}}))
	copy(v.s[i+1:], v.s[i:])
	v.s[i] = n
}

func (v *{{.NameTitle}}Slice) Remove(i int) {
	if i < 0 || i >= len(v.s) {
		fmt.Printf("Vapi::{{.Name}}Slice field_values.go error trying to remove bad index %d\n", i)
		return
	}
	copy(v.s[i:], v.s[i+1:])
	v.s[len(v.s)-1] = reflect.Zero(reflect.TypeOf(v.s[i])).Interface().({{.Name}})
	v.s = v.s[:len(v.s)-1]
}

func (v *{{.NameTitle}}Slice) Count() int {
	return len(v.s)
}

func (v *{{.NameTitle}}Slice) At(i int) {{.Name}} {
	if i < 0 || i >= len(v.s) {
		fmt.Printf("Vapi::{{.Name}}Slice field_values.go invalid index %d\n", i)
	}
	return v.s[i]
}

func new{{.NameTitle}}SliceWithObjects(objects []{{.Name}}) *{{.NameTitle}}Slice {
	pSlice := make([]{{.Name}}, 0, len(objects))
	for _, i := range objects {
		pSlice = append(pSlice, i)
	}
	return &{{.NameTitle}}Slice{pSlice}
}

{{end}}`