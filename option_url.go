package uca

var _ signOption = (*optionUrl)(nil)

type optionUrl struct {
	url string
}

// Url 配置授权
func Url(url string) *optionUrl {
	return &optionUrl{
		url: url,
	}
}

func (u *optionUrl) applySign(options *signOptions) {
	options.url = u.url
}

func (u *optionUrl) apply(options *options) {
	options.url = u.url
}
