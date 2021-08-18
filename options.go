package uca

import (
	`github.com/storezhang/gox`
)

var defaultOptions = &options{
	tencentType: tencentyunTypeD,
	scheme:      gox.URISchemeHttps,

	signParam:      "sign",
	timestampParam: "t",
	removeQuery:    false,
}

type options struct {
	ucaType     Type
	tencentType tencentyunType

	scheme gox.URIScheme
	domain string
	key    string

	signParam      string
	timestampParam string
	removeQuery    bool
}
