<img src="logo.png" style="zoom:25%;" />

### box介绍

这是一款轻量级的Go语言ORM框架，内置sqlite3驱动 ，还在快速迭代中，相信你能3分钟内上手 

目前实现了简单的增删查改，大部分代码学习至geektutu的系列教程

### 快速上手

```go
package main

import (
	"box"
	"box/log"
	"fmt"
)

type User struct {
	Name string `box:"PRIMARY KEY"`
	Age  int
}

func main() {
	engine, _ := box.NewEngine("sqlite3", "test.db")
	defer engine.Close()

	session := engine.NewSession().Model(&User{})
	//增删表
	session.CreateTable()
	defer session.DropTable()
	//判断表存在
	if session.IsExistTable() {
		log.Info("表存在")
	}

	//插入操作
	if affect, err := session.Insert(
		&User{Name: "tomygin", Age: 20},
		&User{Name: "ice", Age: 19},
		&User{Name: "test", Age: 18}); err == nil {
		fmt.Println("成功插入", affect, "条数据")
	}

	//单条查询
	tmp := User{}
	if err := session.Where("Name = ?", "tomygin").First(&tmp); err != nil {
		fmt.Println(err)
	}
	//多条查询
	tmps := []User{}
	if err := session.Where("Age > 10").Find(&tmps); err == nil {
		fmt.Println("拿到数据", tmps)
	}

	//删除
	if _, err := session.Where("Age = ?", 18).Limit(1).Delete(); err != nil {
		fmt.Println(err)
	}

	//更新
	session.Where("Name = ?", "tomygin").Update("Age", 18)

	//查看更新
	session.Where("Name = ?", "tomygin").First(&tmp)
	fmt.Println(tmp)

	//执行原生SQL
	session.Raw("INSERT INTO User (`Name`)  VALUES (?) ", "RAW").Exec()
}

```

### 未来计划

- [ ] 支持钩子函数
- [ ] 事务提交
- [ ] 选项初始化
- [ ] 异步插入
