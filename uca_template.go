package uca

import (
	`net/url`
)

type ucaTemplate struct {
	changcache ucaInternal
	tencentyun ucaInternal
}

func (t *ucaTemplate) Sign(original url.URL, opts ...signOption) (url url.URL, err error) {
	options := defaultSignOptions()
	for _, opt := range opts {
		opt.applySign(options)
	}

	switch options.ucaType {
	case TypeChuangcache:
		url, err = t.changcache.sign(original, options)
	case TypeTencentyun:
		url, err = t.tencentyun.sign(original, options)
	}

	return
}
