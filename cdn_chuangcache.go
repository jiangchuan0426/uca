package uoa

import (
	`fmt`
	`time`

	`github.com/storezhang/gox`
)

const chuangcacheSignPattern = "%s%s%d"

type chuangcache struct {
	cdnBase
}

func (c *chuangcache) sign(original string) (url string, err error) {
	now := time.Now().Unix()
	key := fmt.Sprintf(chuangcacheSignPattern, c.secret, original, now)
	url, err = gox.Md5(key)

	return
}
