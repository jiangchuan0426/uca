package uca

import (
	`net/url`
)

type template struct {
	changcache executor
	tencentyun executor
}

func (t *template) Sign(url *url.URL, opts ...signOption) (err error) {
	options := defaultSignOptions()
	for _, opt := range opts {
		opt.applySign(options)
	}

	// 去掉请求参数，原因：
	// 比如使用腾讯云CDN，原始存储使用COS，如果带上COS的请求参数，会导致报SecretKey不可用的错误
	if options.removeQuery {
		url.RawQuery = ""
	}

	switch options.ucaType {
	case TypeChuangcache:
		err = t.changcache.sign(url, options)
	case TypeTencentyun:
		err = t.tencentyun.sign(url, options)
	}
	url.Scheme = string(options.scheme)
	url.Host = options.domain

	return
}
