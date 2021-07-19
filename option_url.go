package uca

var _ signOption = (*optionDomain)(nil)

type optionDomain struct {
	domain string
	key    string
}

// Domain 配置授权
func Domain(domain string, key string) *optionDomain {
	return &optionDomain{
		domain: domain,
		key:    key,
	}
}

func (u *optionDomain) applySign(options *signOptions) {
	options.domain = u.domain
	options.key = u.key
}

func (u *optionDomain) apply(options *options) {
	options.domain = u.domain
	options.key = u.key
}
