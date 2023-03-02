package session

type options struct {
	notNeedHook bool
}

type optionFunc func(*options)

func CloseHook() optionFunc {
	return func(o *options) {
		o.notNeedHook = true
	}
}

func (s *Session) Options(opts ...optionFunc) {
	for _, opt := range opts {
		opt(&s.options)
	}
}
