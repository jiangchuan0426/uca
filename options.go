package uca

import (
	`github.com/storezhang/gox`
)

var defaultOptions = &options{}

type options struct {
	ucaType Type
	secret  gox.Secret

	domain string
	key    string
}
