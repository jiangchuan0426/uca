package uca

import (
	`fmt`
	`net/url`
	`time`

	`github.com/storezhang/gox`
)

const chuangcacheSignPattern = "%s%s%d"

var _ ucaInternal = (*chuangcacheInternal)(nil)

type chuangcacheInternal struct{}

func (c *chuangcacheInternal) sign(url *url.URL, options *signOptions) (err error) {
	now := time.Now().Unix()
	key := fmt.Sprintf(chuangcacheSignPattern, options.key, url.Path, now)

	query := url.Query()
	var sign string
	if sign, err = gox.Md5(key); nil != err {
		return
	}
	query.Add("KEY1", sign)
	query.Add("KEY2", "timestamp")

	return
}
