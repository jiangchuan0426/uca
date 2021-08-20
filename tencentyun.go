package uca

import (
	`fmt`
	`net/url`
	`strconv`
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

var _ executor = (*tencentyun)(nil)

type (
	tencentyunType uint8
	tencentyun     struct{}
)

func (t *tencentyun) sign(original *url.URL, options *signOptions) (err error) {
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

func (t *tencentyun) signA(url *url.URL, options *signOptions) (err error) {
	nowHex := time.Now().Unix()
	key := fmt.Sprintf(tencentyunSignPatternA, url.EscapedPath(), nowHex, xid.New().String(), options.key)

	query := url.Query()
	var sign string
	if sign, err = gox.Md5(key); nil != err {
		return
	}
	query.Add("sign", sign)
	url.RawQuery = query.Encode()

	return
}

func (t *tencentyun) signB(url *url.URL, options *signOptions) (err error) {
	nowHex := time.Now().Format("20060102150405")
	key := fmt.Sprintf(tencentyunSignPatternB, options.key, nowHex, url.EscapedPath())

	var sign string
	if sign, err = gox.Md5(key); nil != err {
		return
	}
	url.RawPath = fmt.Sprintf("%s/%s%s", nowHex, sign, url.EscapedPath())
	url.Path = fmt.Sprintf("%s/%s%s", nowHex, sign, url.Path)

	return
}

func (t *tencentyun) signC(url *url.URL, options *signOptions) (err error) {
	nowHex := strconv.FormatInt(time.Now().Unix(), 16)
	key := fmt.Sprintf(tencentyunSignPatternC, options.key, url.EscapedPath(), nowHex)

	var sign string
	if sign, err = gox.Md5(key); nil != err {
		return
	}
	url.RawPath = fmt.Sprintf("/%s/%s%s", sign, nowHex, url.EscapedPath())
	url.Path = fmt.Sprintf("/%s/%s%s", sign, nowHex, url.Path)

	return
}

func (t *tencentyun) signD(url *url.URL, options *signOptions) (err error) {
	nowHex := strconv.FormatInt(time.Now().Unix(), 16)
	key := fmt.Sprintf(tencentyunSignPatternD, options.key, url.EscapedPath(), nowHex)

	var sign string
	if sign, err = gox.Md5(key); nil != err {
		return
	}

	query := url.Query()
	query.Add(options.signParam, sign)
	query.Add(options.timestampParam, nowHex)
	url.RawQuery = query.Encode()

	return
}
