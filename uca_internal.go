package uca

import (
	`net/url`
)

type ucaInternal interface {
	sign(original url.URL, options *signOptions) (url url.URL, err error)
}
