package uca

import (
	`github.com/storezhang/gox`
)

var _ signOption = (*optionDomain)(nil)

type optionDomain struct {
	scheme gox.URIScheme
	domain string
	key    string
}

// Url 配置地址
func Url(scheme gox.URIScheme, domain string, key string) *optionDomain {
	return &optionDomain{
		scheme: scheme,
		domain: domain,
		key:    key,
	}
}

func (d *optionDomain) applySign(options *signOptions) {
	options.domain = d.domain
	options.key = d.key
}

func (d *optionDomain) apply(options *options) {
	options.domain = d.domain
	options.key = d.key
}
