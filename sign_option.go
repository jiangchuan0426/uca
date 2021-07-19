package uca

type signOption interface {
	applySign(options *signOptions)
}
