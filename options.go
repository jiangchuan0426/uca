package uca

import (
	`github.com/storezhang/gox`
)

var defaultOptions = &options{}

type options struct {
	ucaType Type
	url     string
	secret  gox.Secret
}
