/**
 * @Title  seek
 * @description  断点续传
 * @Author  沈来
 * @Update  2020/8/28 15:34
 **/
package upload

import (
	"CloudDisk/global"
	"fmt"
	"io"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

func SeekDown(max string, min string, rate int, fileurl string, name string) (string, error) {
	u := strings.Split(fileurl, "/")
	f, err := os.Open("./" + GetSavePath() + "/" + u[4])
	if err != nil {
		return "", err
	}
	defer f.Close()

	start := time.Now()
	filename := url.QueryEscape(name) // 防止中文乱码
	out, err := os.OpenFile(global.FileSetting.DownPath+filename, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return "", err
	}
	defer out.Close()

	temp, err := os.OpenFile("temp.txt", os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		return "", err
	}
	_, _ = temp.Seek(0, io.SeekStart)

	count, _ := strconv.ParseInt(min, 10, 64)
	//step2:设置读，写的位置
	_, _ = f.Seek(count, io.SeekStart)
	_, _ = out.Seek(count, io.SeekStart)

	data := make([]byte, rate, rate)

	n2 := -1            //读取的数据量
	n3 := -1            //写出的数据量
	total := int(count) //读取的总量

	maxx, _ := strconv.Atoi(max)
	for {
		n2, err = f.Read(data)
		if err == io.EOF || n2 == 0 {
			fmt.Println("文件复制完毕。。。")
			temp.Close()
			_ = os.Remove("temp.txt")
			break
		}
		n3, err = out.Write(data[:n2])
		total += n3

		//将复制的总量，存储到临时文件中
		_, _ = temp.Seek(0, io.SeekStart)
		_, _ = temp.WriteString(strconv.Itoa(total))

		fmt.Println("total 总量：", total)
		if total > maxx {
			break
		}
	}

	return fmt.Sprintf("Copied %d bytes in %s\n", total, time.Since(start)), nil
}
