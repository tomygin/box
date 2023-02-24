package session

import "github.com/tomygin/box/log"

func (s *Session) Begin() (err error) {
	log.Info("transaction begin")
	s.tx, err = s.db.Begin()
	if err != nil {
		log.Error(err)
	}

	return
}

func (s *Session) Commit() (err error) {
	log.Info("transaction commit")
	err = s.tx.Commit()
	if err != nil {
		log.Error(err)
	}

	return
}

func (s *Session) RollBack() (err error) {
	log.Info("transaction rollback")
	err = s.tx.Rollback()
	if err != nil {
		log.Error(err)
	}

	return
}
