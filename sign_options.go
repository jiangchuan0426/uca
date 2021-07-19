package uca

type signOptions struct {
	options

	tencentType    tencentyunType
	signParam      string
	timestampParam string
}

func defaultSignOptions() *signOptions {
	return &signOptions{
		tencentType:    tencentyunTypeA,
		signParam:      "sign",
		timestampParam: "timestamp",
	}
}
