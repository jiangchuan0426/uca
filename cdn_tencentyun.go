package uoa

import (
	`fmt`
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

const (
	tencentyunTimestampTypeDecimal tencentyunTimestampType = 1
	tencentyunTimestampTypeHex     tencentyunTimestampType = 2
)

type (
	tencentyunType          uint8
	tencentyunTimestampType uint8
	tencentyun              struct {
		cdnBase

		signType       tencentyunType
		signParam      string
		timestampParam string
		timestampType  tencentyunTimestampType
	}
)

func (t *tencentyun) sign(original string) (url string, err error) {
	switch t.signType {
	case tencentyunTypeA:
		url, err = t.signA(original)
	case tencentyunTypeB:
		url, err = t.signB(original)
	case tencentyunTypeC:
		url, err = t.signC(original)
	case tencentyunTypeD:
		url, err = t.signD(original)
	default:
		url, err = t.signA(original)
	}

	return
}

func (t *tencentyun) signA(original string) (url string, err error) {
	now := time.Now().Unix()
	key := fmt.Sprintf(tencentyunSignPatternA, original, now, xid.New().String(), t.secret)
	url, err = t.format(key)

	return
}

func (t *tencentyun) signB(original string) (url string, err error) {
	now := time.Now().Format("20060102150405")
	key := fmt.Sprintf(tencentyunSignPatternB, t.secret, now, original)
	url, err = t.format(key)

	return
}

func (t *tencentyun) signC(original string) (url string, err error) {
	now := strings.ToUpper(strconv.FormatInt(time.Now().Unix(), 16))
	key := fmt.Sprintf(tencentyunSignPatternC, t.secret, original, now)
	url, err = t.format(key)

	return
}

func (t *tencentyun) signD(original string) (url string, err error) {
	if "" == t.signParam {
		t.signParam = "sign"
	}
	if "" == t.timestampParam {
		t.timestampParam = "timestamp"
	}
	if 0 == t.timestampType {
		t.timestampType = tencentyunTimestampTypeHex
	}

	var now string
	switch t.timestampType {
	case tencentyunTimestampTypeDecimal:
		now = strconv.FormatInt(time.Now().Unix(), 10)
	case tencentyunTimestampTypeHex:
		now = strings.ToUpper(strconv.FormatInt(time.Now().Unix(), 16))
	default:
		now = strings.ToUpper(strconv.FormatInt(time.Now().Unix(), 16))
	}
	key := fmt.Sprintf(tencentyunSignPatternD, t.secret, original, now)
	if url, err = t.format(key); nil != err {
		return
	}
	url = fmt.Sprintf("%s&%s=%s", url, t.timestampParam, now)

	return
}

func (t *tencentyun) format(key string) (url string, err error) {
	var sign string
	if sign, err = gox.Md5(key); nil != err {
		return
	}
	url = fmt.Sprintf("sign=%s", sign)

	return
}
