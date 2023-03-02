package session

import (
	"reflect"

	"github.com/tomygin/box/log"
)

const (
	BeforeQuery  = "BeforeQuery"
	AfterQuery   = "AfterQuery"
	BeforeUpdate = "BeforeUpdate"
	AfterUpdate  = "AfterUpdate"
	BeforeDelete = "BeforeDelete"
	AfterDelete  = "AfterDelete"
	BeforeInsert = "BeforeInsert"
	AfterInsert  = "AfterInsert"
)

// 反射

func (s *Session) CallMethod(method string, value interface{}) {
	//找到当前表 结构体 的 method 方法
	fm := reflect.ValueOf(s.RefTable().Model).MethodByName(method)

	//如果有自定义结构体就不用表结构体
	if value != nil {
		fm = reflect.ValueOf(value).MethodByName(method)
	}

	param := []reflect.Value{reflect.ValueOf(s)}

	if fm.IsValid() {
		if v := fm.Call(param); len(v) > 0 {
			if err, ok := v[0].Interface().(error); ok {
				// panic(err)
				log.Error(err)
			}
		}
	}

}

// 接口

// func (s *Session) CallMethod(method string, value interface{}) {
// 	//当前操作的表对象
// 	// o := s.RefTable().Model
// 	o := reflect.ValueOf(value)
// 	switch method {
// 	case BeforeQuery:
// 		if i, ok := o.Interface().(IBeforeQuery); ok {
// 			i.BeforeQuery(s)
// 		}
// 	default:
// 		panic("Unsupported hook method")
// 	}

// }

// type IBeforeQuery interface {
// 	BeforeQuery(s *Session) error
// }

// type IAfterQuery interface {
// 	AfterQuery(*Session) error
// }

// type IBeforeUpdate interface {
// 	BeforeUpdate(*Session) error
// }
