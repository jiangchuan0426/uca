package uca

import (
	`github.com/storezhang/gox`
)

var defaultOptions = &options{
	tencentType: tencentyunTypeD,
	scheme:      gox.URISchemeHttps,
	removeQuery: false,
}

type options struct {
	ucaType     Type
	tencentType tencentyunType

	scheme      gox.URIScheme
	domain      string
	key         string
	removeQuery bool
}
