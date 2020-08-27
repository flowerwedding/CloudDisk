/**
 * @Title  qtcode
 * @description  生成二维码
 * @Author  沈来
 * @Update  2020/8/27 16:24
 **/
package qtcode

import (
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"image/png"
	"os"
)

func Qt(url string, add string) error {

	qrCode, err := qr.Encode(url, qr.M, qr.Auto)
	if err != nil {
		return err
	}

	qrCode, err = barcode.Scale(qrCode, 256, 256)
	if err != nil {
		return err
	}

	file, err := os.Create(add) //保存路径和名字
	if err != nil {
		return err
	}

	defer file.Close()

	err = png.Encode(file, qrCode)
	if err != nil {
		return err
	}

	return nil
}
