package uca

type signOptions struct {
	*options

	tencentType    tencentyunType
	signParam      string
	timestampParam string
}

func defaultSignOptions() *signOptions {
	return &signOptions{
		options: defaultOptions,

		tencentType:    tencentyunTypeA,
		signParam:      "sign",
		timestampParam: "timestamp",
	}
}
