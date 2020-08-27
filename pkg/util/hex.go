/**
 * @Title  md5
 * @description  上传文件格式化
 * @Author  沈来
 * @Update  2020/8/6 19:13
 **/
package util

import (
	"encoding/hex"
)

func Encode(value string) string {
	maxEnLen := hex.EncodedLen(len([]byte(value)))
	dst1 := make([]byte, maxEnLen)
	n := hex.Encode(dst1, []byte(value))

	return string(dst1[:n])
}

func Decode(value string) string {
	maxDeLen := hex.DecodedLen(len([]byte(value)))
	dst1 := make([]byte, maxDeLen)
	n, _ := hex.Decode(dst1, []byte(value))

	return string(dst1[:n])
}
