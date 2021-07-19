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

var _ ucaInternal = (*tencentyunInternal)(nil)

type (
	tencentyunType     uint8
	tencentyunInternal struct{}
)

func (t *tencentyunInternal) sign(original *url.URL, options *signOptions) (err error) {
	switch options.tencentType {
	case tencentyunTypeA:
		err = t.signA(original, options)
	case tencentyunTypeB:
		err = t.signB(original, options)
	case tencentyunTypeC:
		err = t.signC(original, options)
	case tencentyunTypeD:
		err = t.signD(original, options)
	default:
		err = t.signA(original, options)
	}

	return
}

func (t *tencentyunInternal) signA(url *url.URL, options *signOptions) (err error) {
	now := time.Now().Unix()
	key := fmt.Sprintf(tencentyunSignPatternA, url.Path, now, xid.New().String(), options.key)

	query := url.Query()
	var sign string
	if sign, err = gox.Md5(key); nil != err {
		return
	}
	query.Add("sign", sign)

	return
}

func (t *tencentyunInternal) signB(url *url.URL, options *signOptions) (err error) {
	now := time.Now().Format("20060102150405")
	key := fmt.Sprintf(tencentyunSignPatternB, options.key, now, url.Path)

	var sign string
	if sign, err = gox.Md5(key); nil != err {
		return
	}
	url.Path = fmt.Sprintf("%s%s%s", now, sign, url.Path)

	return
}

func (t *tencentyunInternal) signC(url *url.URL, options *signOptions) (err error) {
	now := strings.ToUpper(strconv.FormatInt(time.Now().Unix(), 16))
	key := fmt.Sprintf(tencentyunSignPatternC, options.key, url.Path, now)

	var sign string
	if sign, err = gox.Md5(key); nil != err {
		return
	}
	url.Path = fmt.Sprintf("%s%s%s", sign, now, url.Path)

	return
}

func (t *tencentyunInternal) signD(url *url.URL, options *signOptions) (err error) {
	now := strings.ToUpper(strconv.FormatInt(time.Now().Unix(), 16))
	key := fmt.Sprintf(tencentyunSignPatternD, options.key, url.Path, now)

	var sign string
	if sign, err = gox.Md5(key); nil != err {
		return
	}

	query := url.Query()
	query.Add(options.signParam, sign)
	query.Add(options.timestampParam, now)

	return
}
