/**
 * @Title  md5
 * @description  #
 * @Author  沈来
 * @Update  2020/8/26 21:25
 **/
package util

import (
	"crypto/md5"
	"encoding/hex"
)

func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}
