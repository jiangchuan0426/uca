package uca

import (
	`net/url`
)

type template struct {
	changcache executor
	tencentyun executor
}

func (t *template) Sign(original *url.URL, opts ...signOption) (err error) {
	options := defaultSignOptions()
	for _, opt := range opts {
		opt.applySign(options)
	}

	// 去掉请求参数，原因：
	// 比如使用腾讯云CDN，原始存储使用COS，如果带上COS的请求参数，会导致报SecretKey不可用的错误
	if options.removeQuery {
		original.RawQuery = ""
	}

	switch options.ucaType {
	case TypeChuangcache:
		err = t.changcache.sign(original, options)
	case TypeTencentyun:
		err = t.tencentyun.sign(original, options)
	}
	original.Scheme = string(options.scheme)
	original.Host = options.domain

	return
}
