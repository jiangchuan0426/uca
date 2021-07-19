package uca

import (
	`fmt`
	`net/url`
	`time`

	`github.com/storezhang/gox`
)

const chuangcacheSignPattern = "%s%s%d"

var _ ucaInternal = (*chuangcache)(nil)

type chuangcache struct{}

func (c *chuangcache) sign(original url.URL, options *signOptions) (url url.URL, err error) {
	now := time.Now().Unix()
	key := fmt.Sprintf(chuangcacheSignPattern, options.secret.Key, original.Path, now)

	url = original
	query := url.Query()

	var sign string
	if sign, err = gox.Md5(key); nil != err {
		return
	}
	query.Add("KEY1", sign)
	query.Add("KEY2", "timestamp")

	return
}
