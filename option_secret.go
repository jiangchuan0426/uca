package uca

var _ signOption = (*optionSecret)(nil)

type optionSecret struct {
	ucaType Type
	id      string
	key     string

	tencentType tencentyunType
}

// Chuangcache 配置创世云授权
func Chuangcache(ak string, sk string) *optionSecret {
	return &optionSecret{
		ucaType: TypeChuangcache,
		id:      ak,
		key:     sk,
	}
}

// TencentyunTypeA 配置腾讯云授权
func TencentyunTypeA(secretId string, secretKey string) *optionSecret {
	return &optionSecret{
		ucaType:     TypeTencentyun,
		id:          secretId,
		key:         secretKey,
		tencentType: tencentyunTypeA,
	}
}

// TencentyunTypeB 配置腾讯云授权
func TencentyunTypeB(secretId string, secretKey string) *optionSecret {
	return &optionSecret{
		ucaType:     TypeTencentyun,
		id:          secretId,
		key:         secretKey,
		tencentType: tencentyunTypeB,
	}
}

// TencentyunTypeC 配置腾讯云授权
func TencentyunTypeC(secretId string, secretKey string) *optionSecret {
	return &optionSecret{
		ucaType:     TypeTencentyun,
		id:          secretId,
		key:         secretKey,
		tencentType: tencentyunTypeC,
	}
}

// TencentyunTypeD 配置腾讯云授权
func TencentyunTypeD(secretId string, secretKey string) *optionSecret {
	return &optionSecret{
		ucaType:     TypeTencentyun,
		id:          secretId,
		key:         secretKey,
		tencentType: tencentyunTypeD,
	}
}

func (s *optionSecret) applySign(options *signOptions) {
	options.ucaType = s.ucaType
	options.secret.Id = s.id
	options.secret.Key = s.key
}

func (s *optionSecret) apply(options *options) {
	options.ucaType = s.ucaType
	options.secret.Id = s.id
	options.secret.Key = s.key
}
