package uca

type signOptions struct {
	*options
}

func defaultSignOptions() *signOptions {
	return &signOptions{
		options: defaultOptions,
	}
}
