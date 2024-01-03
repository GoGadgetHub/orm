package schema

type Field struct {
	Name string // 字段名
	Type string // 类型
	Tag  string // 约束条件
}

type Schema struct {
	Model      interface{}
	Name       string
	Fields     []*Field
	FieldNames []string
	filedMap   map[string]*Field
}

func (s *Schema) GetField(name string) *Field {
	return s.filedMap[name]
}
