package uca

import (
	`net/url`
)

type ucaInternal interface {
	sign(url *url.URL, options *signOptions) (err error)
}
