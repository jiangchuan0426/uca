package uca

import (
	`net/url`
)

type executor interface {
	sign(url *url.URL, options *signOptions) (err error)
}
