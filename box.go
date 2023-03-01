package box

import "github.com/tomygin/box/session"

// 事务的回调函数
type TxFunc func(*session.Session) (interface{}, error)

func (e *Engine) Transaction(f TxFunc) (result interface{}, err error) {
	s := e.NewSession()
	if err := s.Begin(); err != nil {
		return nil, err
	}

	defer func() {
		if p := recover(); p != nil {
			_ = s.RollBack()
		} else if err != nil {
			_ = s.RollBack()
		} else {
			err = s.Commit()
		}
	}()
	return f(s)
}
