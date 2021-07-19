package uca

import (
	`net/url`
)

// Uca 统一CDN接口
type Uca interface {
	Sign(original url.URL, opts ...signOption) (url url.URL, err error)
}
