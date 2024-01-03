package parse

import (
	"go/ast"
	"orm/dialect"
	"orm/schema"
	"reflect"
)

// Parse 将对象解析成 Schema 实例
func Parse(dest interface{}, d dialect.Dialect) *schema.Schema {
	modelType := reflect.Indirect(reflect.ValueOf(dest)).Type()
	s := &schema.Schema{
		Model:    dest,
		Name:     modelType.Name(),
		FiledMap: make(map[string]*schema.Field),
	}

	for i := 0; i < modelType.NumField(); i++ {
		p := modelType.Field(i)
		if !p.Anonymous && ast.IsExported(p.Name) {
			field := &schema.Field{
				Name: p.Name,
				Type: d.DataTypeOf(reflect.Indirect(reflect.New(p.Type))),
			}
			if v, ok := p.Tag.Lookup("orm"); ok {
				field.Tag = v
			}
			s.Fields = append(s.Fields, field)
			s.FieldNames = append(s.FieldNames, p.Name)
			s.FiledMap[p.Name] = field
		}
	}
	return s
}
