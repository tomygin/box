package clause

import "strings"

type Type int

const (
	INSERT Type = iota + 1
	VALUES
	SELECT
	LIMIT
	OFFSET
	WHERE
	ORDERBY
	UPDATE
	DELETE
	COUNT
)

type Clause struct {
	sql     map[Type]string
	sqlVars map[Type][]interface{}
}

// 生成子句
func (c *Clause) Set(name Type, vars ...interface{}) {
	if c.sql == nil {
		c.sql = make(map[Type]string)
		c.sqlVars = make(map[Type][]interface{})
	}
	sql, vars := generators[name](vars...)
	c.sql[name] = sql
	c.sqlVars[name] = vars
}

// 构造sql语句
// 构造完成后清空构造的sql语句
func (c *Clause) Build(orders ...Type) (string, []interface{}) {

	defer func() {
		c.sql = nil
		c.sqlVars = nil
	}()

	var sqls []string
	var vars []interface{}
	for _, order := range orders {
		if sql, ok := c.sql[order]; ok {
			sqls = append(sqls, sql)
			vars = append(vars, c.sqlVars[order]...)
		}
	}
	return strings.Join(sqls, " "), vars
}
