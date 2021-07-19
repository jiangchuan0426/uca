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

func (s *optionUrl) applySign(options *signOptions) {
	options.url = s.url
}
