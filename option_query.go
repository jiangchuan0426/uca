package uca

var _ signOption = (*optionQuery)(nil)

type optionQuery struct {
	remove bool
}

// Query 配置是否去掉请求参数
func Query(remove bool) *optionQuery {
	return &optionQuery{
		remove: remove,
	}
}

// RetainQuery 保留请求参数
func RetainQuery() *optionQuery {
	return Query(false)
}

// RemoveQuery 去掉请求参数
func RemoveQuery() *optionQuery {
	return Query(true)
}

func (t *optionQuery) applySign(options *signOptions) {
	options.removeQuery = t.remove
}

func (t *optionQuery) apply(options *options) {
	options.removeQuery = t.remove
}
