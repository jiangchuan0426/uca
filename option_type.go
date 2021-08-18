package uca

var _ signOption = (*optionType)(nil)

type optionType struct {
	ucaType     Type
	tencentType tencentyunType
}

// Chuangcache 配置创世云授权
func Chuangcache() *optionType {
	return &optionType{
		ucaType: TypeChuangcache,
	}
}

// TencentyunTypeA 配置腾讯云授权
func TencentyunTypeA() *optionType {
	return &optionType{
		ucaType:     TypeTencentyun,
		tencentType: tencentyunTypeA,
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

func (t *optionType) applySign(options *signOptions) {
	options.ucaType = t.ucaType
	options.tencentType = t.tencentType
}

func (t *optionType) apply(options *options) {
	options.ucaType = t.ucaType
	options.tencentType = t.tencentType
}
