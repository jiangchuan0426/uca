package uca

var _ signOption = (*optionType)(nil)

type optionType struct {
	ucaType     Type
	tencentType tencentyunType

	signParam      string
	timestampParam string
}

// Chuangcache 配置创世云授权
func Chuangcache() *optionType {
	return &optionType{
		ucaType: TypeChuangcache,

		signParam:      "sign",
		timestampParam: "t",
	}
}

// TencentyunTypeA 配置腾讯云授权
func TencentyunTypeA() *optionType {
	return &optionType{
		ucaType:     TypeTencentyun,
		tencentType: tencentyunTypeA,

		signParam: "sign",
	}
}

// TencentyunTypeAWithConfig 配置腾讯云授权
func TencentyunTypeAWithConfig(sign string) *optionType {
	return &optionType{
		ucaType:     TypeTencentyun,
		tencentType: tencentyunTypeA,
		signParam:   sign,
	}
}

// TencentyunTypeB 配置腾讯云授权
func TencentyunTypeB() *optionType {
	return &optionType{
		ucaType:     TypeTencentyun,
		tencentType: tencentyunTypeB,
	}
}

// TencentyunTypeC 配置腾讯云授权
func TencentyunTypeC() *optionType {
	return &optionType{
		ucaType:     TypeTencentyun,
		tencentType: tencentyunTypeC,
	}
}

// TencentyunTypeD 配置腾讯云授权
func TencentyunTypeD() *optionType {
	return &optionType{
		ucaType:     TypeTencentyun,
		tencentType: tencentyunTypeD,
	}
}

// TencentyunTypeDWithConfig 配置腾讯云授权
func TencentyunTypeDWithConfig(sign string, timestamp string) *optionType {
	return &optionType{
		ucaType:     TypeTencentyun,
		tencentType: tencentyunTypeD,

		signParam:      sign,
		timestampParam: timestamp,
	}
}

func (t *optionType) applySign(options *signOptions) {
	options.ucaType = t.ucaType
	options.tencentType = t.tencentType

	if "" != t.signParam {
		options.signParam = t.signParam
	}
	if "" != t.timestampParam {
		options.timestampParam = t.timestampParam
	}
}

func (t *optionType) apply(options *options) {
	options.ucaType = t.ucaType
	options.tencentType = t.tencentType

	if "" != t.signParam {
		options.signParam = t.signParam
	}
	if "" != t.timestampParam {
		options.timestampParam = t.timestampParam
	}
}
