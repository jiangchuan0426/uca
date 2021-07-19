package uca

import (
	`net/url`
)

type ucaTemplate struct {
	changcache ucaInternal
	tencentyun ucaInternal
}

func (t *ucaTemplate) Sign(url *url.URL, opts ...signOption) (err error) {
	options := defaultSignOptions()
	for _, opt := range opts {
		opt.applySign(options)
	}

	switch options.ucaType {
	case TypeChuangcache:
		err = t.changcache.sign(url, options)
	case TypeTencentyun:
		err = t.tencentyun.sign(url, options)
	}

	return
}
