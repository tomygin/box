package session

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/tomygin/box/log"
	"github.com/tomygin/box/schema"
)

// 如果当前对象没有被解析为Schema就解析
func (s *Session) Model(value interface{}) *Session {
	if s.refTable == nil || reflect.TypeOf(value) != reflect.TypeOf(s.refTable.Model) {
		s.refTable = schema.Parse(value, s.dialect)
	}
	return s
}

// 用于返回反射的Schema
func (s *Session) RefTable() *schema.Schema {
	if s.refTable == nil {
		log.Error("Model is not set ")
	}
	return s.refTable
}

func (s *Session) CreateTable() error {
	table := s.RefTable()
	var col []string
	for _, field := range table.Fields {
		col = append(col, fmt.Sprintf("%s %s %s", field.Name, field.Type, field.Tag))
	}

	desc := strings.Join(col, ",")
	_, err := s.Raw(fmt.Sprintf("CREATE TABLE %s (%s);", table.Name, desc)).Exec()
	return err
}

func (s *Session) DropTable() error {
	_, err := s.Raw(fmt.Sprintf("DROP TABLE IF EXISTS %s", s.RefTable().Name)).Exec()
	return err
}

func (s *Session) IsExistTable() bool {
	sql, values := s.dialect.TableExistSql(s.RefTable().Name)
	row := s.Raw(sql, values...).QueryRow()
	var tmp string
	row.Scan(&tmp)
	return tmp == s.RefTable().Name
}
