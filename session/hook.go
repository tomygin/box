package session

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

// func (s *Session) CallMethod(method string, value any) {
// 	//找到当前表 结构体 的 method 方法
// 	fm := reflect.ValueOf(s.RefTable().Model).MethodByName(method)

// 	//如果有自定义结构体就不用表结构体
// 	if value != nil {
// 		fm = reflect.ValueOf(value).MethodByName(method)
// 	}

// 	// param := []reflect.Value{reflect.ValueOf(s)}
// 	param := []reflect.Value{}

// 	if fm.IsValid() {
// 		if v := fm.Call(param); len(v) > 0 {
// 			if err, ok := v[0].Interface().(error); ok {
// 				log.Error(err)
// 			}
// 		}
// 	}

// }

// 接口

// type IBeforeQuery interface {
// 	BeforeQuery(*Session) error
// }

// type IAfterQuery interface {
// 	AfterQuery(*Session) error
// }

// type IBeforeUpdate interface {
// 	BeforeUpdate(*Session) error
// }
