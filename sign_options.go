package uca

type signOptions struct {
	*options

	signParam      string
	timestampParam string
}

func defaultSignOptions() *signOptions {
	return &signOptions{
		options: defaultOptions,

		signParam:      "sign",
		timestampParam: "t",
	}
}
