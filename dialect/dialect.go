package dialect

import "reflect"

var dialectsMap = map[string]Dialect{}

type Dialect interface {
	// DataTypeOf 转换对应数据库的类型
	DataTypeOf(typ reflect.Value) string
	// TableExistSQL 是否存在SQL语句
	TableExistSQL(tableName string) (string, []interface{})
}

// RegisterDialect 注册方言
func RegisterDialect(name string, dialect Dialect) {
	dialectsMap[name] = dialect
}

// GetDialect 获取方言
func GetDialect(name string) (dialect Dialect, ok bool) {
	dialect, ok = dialectsMap[name]
	return
}
