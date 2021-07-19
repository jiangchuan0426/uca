package uca

import (
	`fmt`
	`net/url`
	`strconv`
	`strings`
	`time`

	`github.com/rs/xid`
	`github.com/storezhang/gox`
)

const (
	tencentyunSignPatternA = "%s%d0%s%s"
	tencentyunSignPatternB = "%s%s%s"
	tencentyunSignPatternC = "%s%s%s"
	tencentyunSignPatternD = "%s%s%s"
)

const (
	tencentyunTypeA tencentyunType = 1
	tencentyunTypeB tencentyunType = 2
	tencentyunTypeC tencentyunType = 3
	tencentyunTypeD tencentyunType = 4
)

var (
	_ ucaInternal = (*tencentyun)(nil)
	_ Uca         = (*tencentyun)(nil)
)

type (
	tencentyunType uint8
	tencentyun     struct {
		template ucaTemplate
	}
)

func (t *tencentyun) Sign(original url.URL, opts ...signOption) (url url.URL, err error) {
	return t.template.Sign(original, opts...)
}

func (t *tencentyun) sign(original url.URL, options *signOptions) (url url.URL, err error) {
	switch options.tencentType {
	case tencentyunTypeA:
		url, err = t.signA(original, options)
	case tencentyunTypeB:
		url, err = t.signB(original, options)
	case tencentyunTypeC:
		url, err = t.signC(original, options)
	case tencentyunTypeD:
		url, err = t.signD(original, options)
	default:
		url, err = t.signA(original, options)
	}

	return
}

func (t *tencentyun) signA(original url.URL, options *signOptions) (url url.URL, err error) {
	now := time.Now().Unix()
	key := fmt.Sprintf(tencentyunSignPatternA, original.Path, now, xid.New().String(), options.secret.Key)

	url = original
	query := url.Query()

	var sign string
	if sign, err = gox.Md5(key); nil != err {
		return
	}
	query.Add("sign", sign)

	return
}

func (t *tencentyun) signB(original url.URL, options *signOptions) (url url.URL, err error) {
	now := time.Now().Format("20060102150405")
	key := fmt.Sprintf(tencentyunSignPatternB, options.secret.Key, now, original.Path)

	url = original
	var sign string
	if sign, err = gox.Md5(key); nil != err {
		return
	}
	url.Path = fmt.Sprintf("%s%s%s", now, sign, original.Path)

	return
}

func (t *tencentyun) signC(original url.URL, options *signOptions) (url url.URL, err error) {
	now := strings.ToUpper(strconv.FormatInt(time.Now().Unix(), 16))
	key := fmt.Sprintf(tencentyunSignPatternC, options.secret.Key, original.Path, now)

	url = original
	var sign string
	if sign, err = gox.Md5(key); nil != err {
		return
	}
	url.Path = fmt.Sprintf("%s%s%s", sign, now, original.Path)

	return
}

func (t *tencentyun) signD(original url.URL, options *signOptions) (url url.URL, err error) {
	now := strings.ToUpper(strconv.FormatInt(time.Now().Unix(), 16))
	key := fmt.Sprintf(tencentyunSignPatternD, options.secret.Key, original.Path, now)

	var sign string
	if sign, err = gox.Md5(key); nil != err {
		return
	}

	url = original
	query := url.Query()
	query.Add(options.signParam, sign)
	query.Add(options.timestampParam, now)

	return
}
