<img src="logo.png" style="zoom:15%;" />

### box介绍

这是一款轻量级的Go语言ORM框架，学习于Geektutu的系列教程，内置sqlite3驱动 ，还在递归更新中，相信你能3分钟内上手，目前实现了增删查改,事务，hook

### 更新下载

```go
//在当前工作区下输入
go get -u github.com/tomygin/box@latest
```

### 快速上手

```go
package main

import (
	"github.com/tomygin/box"
	"github.com/tomygin/box/log"
	"github.com/tomygin/box/session"
)

type User struct {
	Name string `box:"PRIMARY KEY"`
	Age  int
}

func main() {
	engine, _ := box.NewEngine("sqlite3", "test.db")
	defer engine.Close()

	s := engine.NewSession().Model(&User{})

	// 增删表
	s.CreateTable()
	defer s.DropTable()

	// 判断表存在
	if s.IsExistTable() {
		log.Info("表存在")
	}

	// 插入操作
	if affect, err := s.Insert(
		&User{Name: "tomygin", Age: 20},
		&User{Name: "ice", Age: 19},
		&User{Name: "test", Age: 18}); err == nil {
		log.Info("成功插入", affect, "条数据")
	}

	// 单条查询
	tmp := User{}
	if err := s.Where("Name = ?", "tomygin").First(&tmp); err != nil {
		log.Error(err)
	}

	// 多条查询
	tmps := []User{}
	if err := s.Where("Age > 10").Find(&tmps); err == nil {
		log.Info("拿到数据", tmps)
	}

	// 删除
	if _, err := s.Where("Age = ?", 18).Limit(1).Delete(); err != nil {
		log.Error(err)
	}

	// 更新
	s.Where("Name = ?", "tomygin").Update("Age", 18)

	// 查看更新
	s.Where("Name = ?", "tomygin").First(&tmp)
	log.Info(tmp)

	// 执行原生SQL
	s.Raw("INSERT INTO User (`Name`)  VALUES (?) ", "RAW").Exec()

	// 一键事务，失败自动回滚
	r, err := engine.Transaction(func(s *session.Session) (interface{}, error) {
		s.Model(&User{})
		s.CreateTable()
		s.Insert(&User{Name: "tomygin"})
		t := User{}
		err := s.Where("Name = ?", "tomygin").First(&t)
		return t, err
	})
	log.Info(r, err)

	//日志分级
	log.SetLevel(log.ErrorLevel)

}

// 钩子函数
func (t *User) BeforeQuery(s *session.Session) error {
	log.Info("---钩子函数运行成功---")
	return nil
}

```

```go
// 可用的钩子函数
BeforeQuery  
AfterQuery   
BeforeUpdate 
AfterUpdate  
BeforeDelete 
AfterDelete  
BeforeInsert 
AfterInsert  
```



### 未来计划

- [x] 支持钩子函数
- [x] 事务提交
- [x] 选项初始化
- [ ] 分页
- [ ] 异步插入

