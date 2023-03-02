package session

// 选项卡初始化
type options struct {
	notNeedOpenHook bool
}

// 用于选项卡初始化的函数
type modify func(options)

// 对于选项卡初始化的支持
func (s *Session) OPtions(opts ...modify) {
	for _, use := range opts {
		use(s.ops)
	}
}

func CloseHook() modify {
	return func(o options) {
		o.notNeedOpenHook = true
	}
}
