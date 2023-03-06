package session

import (
	"errors"
	"reflect"

	"github.com/tomygin/box/clause"
)

func (s *Session) Insert(values ...interface{}) (int64, error) {

	s.CallMethod(BeforeInsert, nil)
	defer s.CallMethod(AfterInsert, nil)

	recordValues := make([]interface{}, 0)
	for _, value := range values {
		table := s.Model(value).RefTable()
		s.clause.Set(clause.INSERT, table.Name, table.FieldNames)
		recordValues = append(recordValues, table.RecordValues(value))
	}

	s.clause.Set(clause.VALUES, recordValues...)
	sql, vars := s.clause.Build(clause.INSERT, clause.VALUES)
	resout, err := s.Raw(sql, vars...).Exec()
	if err != nil {
		return 0, err
	}
	return resout.RowsAffected()
}

func (s *Session) Find(values interface{}) error {

	s.CallMethod(BeforeQuery, nil)
	defer s.CallMethod(AfterQuery, nil)

	destSlice := reflect.Indirect(reflect.ValueOf(values))
	destType := destSlice.Type().Elem()
	table := s.Model(reflect.New(destType).Elem().Interface()).RefTable()

	s.clause.Set(clause.SELECT, table.Name, table.FieldNames)
	sql, vars := s.clause.Build(clause.SELECT, clause.WHERE, clause.ORDERBY, clause.LIMIT)
	rows, err := s.Raw(sql, vars...).QueryRows()
	if err != nil {
		return err
	}

	for rows.Next() {
		dest := reflect.New(destType).Elem()
		var values []interface{}
		for _, name := range table.FieldNames {
			values = append(values, dest.FieldByName(name).Addr().Interface())
		}
		if err := rows.Scan(values...); err != nil {
			return err
		}
		destSlice.Set(reflect.Append(destSlice, dest))
	}
	return rows.Close()
}

func (s *Session) Update(kv ...interface{}) (int64, error) {

	s.CallMethod(BeforeUpdate, nil)
	defer s.CallMethod(AfterUpdate, nil)

	m, ok := kv[0].(map[string]interface{})
	if !ok {
		m = make(map[string]interface{})
		for i := 0; i < len(kv); i += 2 {
			m[kv[i].(string)] = kv[i+1]
		}
	}
	s.clause.Set(clause.UPDATE, s.RefTable().Name, m)
	sql, vars := s.clause.Build(clause.UPDATE, clause.WHERE)
	result, err := s.Raw(sql, vars...).Exec()
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (s *Session) Delete() (int64, error) {

	s.CallMethod(BeforeDelete, nil)
	defer s.CallMethod(AfterDelete, nil)

	s.clause.Set(clause.DELETE, s.RefTable().Name)
	sql, vars := s.clause.Build(clause.DELETE, clause.WHERE)
	result, err := s.Raw(sql, vars...).Exec()
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (s *Session) Count() (int64, error) {
	s.clause.Set(clause.COUNT, s.RefTable().Name)
	sql, vars := s.clause.Build(clause.COUNT, clause.WHERE)
	row := s.Raw(sql, vars...).QueryRow()

	var tmp int64
	err := row.Scan(&tmp)
	return tmp, err

}

func (s *Session) Limit(num int) *Session {
	s.clause.Set(clause.LIMIT, num)
	return s
}

// 用于跳过多少条数据
func (s *Session) OFFSET(num int) *Session {
	s.clause.Set(clause.OFFSET, num)
	return s
}

// 指定pageId第几页,启始页第0页,pageCount一页有多少数据
func (s *Session) Page(pageId, pageCount int) *Session {
	offsetCount := pageId * pageCount
	s.Limit(pageCount)
	s.OFFSET(offsetCount)
	return s
}

func (s *Session) Where(desc string, args ...interface{}) *Session {
	var vars []interface{}
	s.clause.Set(clause.WHERE, append(append(vars, desc), args...)...)
	return s
}

func (s *Session) OrderBy(desc string) *Session {
	s.clause.Set(clause.ORDERBY, desc)
	return s
}

func (s *Session) First(value interface{}) error {

	s.CallMethod(BeforeQuery, nil)
	defer s.CallMethod(AfterQuery, nil)

	dest := reflect.Indirect(reflect.ValueOf(value))
	destSlice := reflect.New(reflect.SliceOf(dest.Type())).Elem()
	if err := s.Limit(1).Find(destSlice.Addr().Interface()); err != nil {
		return err
	}
	if destSlice.Len() == 0 {
		return errors.New("NOT FOUND")
	}
	dest.Set(destSlice.Index(0))
	return nil
}
