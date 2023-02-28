package dialect

import "reflect"

var dialectsMap = map[string]Dialect{}

type Dialect interface {
	DataType(typ reflect.Value) string
	TableExistSql(tableName string) (string, []interface{})
}

func RegisterDialect(name string, dialet Dialect) {
	dialectsMap[name] = dialet
}

func GetDialect(name string) (dialect Dialect, ok bool) {
	dialect, ok = dialectsMap[name]
	return
}
