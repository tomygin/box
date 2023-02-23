package box

import (
	"box/dialect"
	"box/log"
	"box/session"
	"database/sql"
)

type Engine struct {
	db      *sql.DB
	dialect dialect.Dialect
}

func NewEngine(driver, source string) (e *Engine, err error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Error(err)
		return
	}

	//测试连接
	if err = db.Ping(); err != nil {
		log.Error(err)
		return
	}

	//获取sql方言
	dial, ok := dialect.GetDialect(driver)
	if !ok {
		log.Error("dialect %s Not Found ", driver)
		return
	}

	e = &Engine{db: db, dialect: dial}

	log.Infof("Connect %s success \n", source)
	return
}

func (e *Engine) Close() {
	if err := e.db.Close(); err != nil {
		log.Error("Failed to close database ")
	}
	log.Info("Close database success ")
}

func (e *Engine) NewSession() *session.Session {
	return session.New(e.db, e.dialect)
}
