package uca

type option interface {
	apply(options *options)
}
