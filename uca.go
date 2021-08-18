package uca

import (
	`net/url`
)

// Uca 统一CDN接口
type Uca interface {
	Sign(url *url.URL, opts ...signOption) (err error)
}

// New 创建统一CDN接口
func New(opts ...option) Uca {
	for _, opt := range opts {
		opt.apply(defaultOptions)
	}

	return &template{
		changcache: &chuangcache{},
		tencentyun: &tencentyun{},
	}
}
