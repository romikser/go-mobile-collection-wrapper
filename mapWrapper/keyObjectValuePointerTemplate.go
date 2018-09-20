package mapWrapper

const keyObjectValuePointerTemplate = `{{range .Types}}
type {{.KeyTitle}}{{.Value}}Map struct {
	s map[{{.Key}}]*{{.Value}}
}

func New{{.KeyTitle}}{{.Value}}Map() *{{.KeyTitle}}{{.Value}}Map {
	return &{{.KeyTitle}}{{.Value}}Map{}
}

func (m *{{.KeyTitle}}{{.Value}}Map) Clear() {
	m.s = make(map[{{.Key}}]*{{.Value}})
}

func (m *{{.KeyTitle}}{{.Value}}Map) Equal(rhs *{{.KeyTitle}}{{.Value}}Map) bool {
	if rhs == nil {
		return false
	}

	if len(m.s) != len(rhs.s) {
		return false
	}

	for k := range m.s {
		if m.s[k] != rhs.s[k] {
			return false
		}
	}

	return true
}

func (m *{{.KeyTitle}}{{.Value}}Map) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.s)
}

func (m *{{.KeyTitle}}{{.Value}}Map) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &m.s)
}

func (m *{{.KeyTitle}}{{.Value}}Map) Copy(rhs *{{.KeyTitle}}{{.Value}}Map) {
	m.s = make(map[{{.Key}}]*{{.Value}})
	for k, v := range rhs.s {
		m.s[k] = v
	}
}

func (m *{{.KeyTitle}}{{.Value}}Map) Clone() *{{.KeyTitle}}{{.Value}}Map {
	nS := make(map[{{.Key}}]*{{.Value}})
	for k, v := range m.s {
		nS[k] = v
	}
	return &{{.KeyTitle}}{{.Value}}Map{
		s: nS,
	}
}

func (m *{{.KeyTitle}}{{.Value}}Map) Key(rhs *{{.Value}}) {{.Key}} {
	for i, lhs := range m.s {
		if lhs == rhs {
			return i
		}
	}
	return reflect.Zero(reflect.TypeOf(m.s).Key()).Interface().({{.Key}})
}

func (m *{{.KeyTitle}}{{.Value}}Map) Set(key {{.Key}}, n *{{.Value}}) {
	m.s[key] = n
}

func (m *{{.KeyTitle}}{{.Value}}Map) Remove(key {{.Key}}) {
	delete(m.s, key)
}

func (m *{{.KeyTitle}}{{.Value}}Map) Count() int {
	return len(m.s)
}

func (m *{{.KeyTitle}}{{.Value}}Map) At(key {{.Key}}) *{{.Value}} {
	u, ok := m.s[key]
	if !ok {
		return nil
	}
	return u
}

func (m *{{.KeyTitle}}{{.Value}}Map) Keys() *{{.KeyTitle}}Slice {
	keys := New{{.KeyTitle}}Slice()
	for k := range m.s {
		keys.Append(k)
	}
	return keys
}

func (m *{{.KeyTitle}}{{.Value}}Map) objectsMap() map[{{.Key}}]{{.Value}} {
	res := make(map[{{.Key}}]{{.Value}})
	for k, v:= range m.s {
		if v == nil {
			continue
		}
		res[k] = *v
	}
	return res
}

func new{{.KeyTitle}}{{.Value}}MapWithObjects(objects map[{{.Key}}]{{.Value}}) *{{.KeyTitle}}{{.Value}}Map {
	pMap := make(map[{{.Key}}]*{{.Value}})
	for k, v := range objects {
		v1 := v
		pMap[k] = &v1
	}
	return &{{.KeyTitle}}{{.Value}}Map{pMap}
}

{{end}}`
