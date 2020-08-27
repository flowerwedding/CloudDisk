/**
 * @Title  upload
 * @description  文件上传
 * @Author  沈来
 * @Update  2020/8/6 19:53
 **/
package service

import (
	"CloudDisk/global"
	"CloudDisk/internal/model"
	"CloudDisk/pkg/limiter"
	"CloudDisk/pkg/upload"
	"errors"
	"fmt"
	"github.com/juju/ratelimit"
	"io"
	"mime/multipart"
	"net/url"
	"os"
	"strings"
	"time"
)

type FileInfo struct {
	Name      string
	AccessUrl string
}

func (svc *Service) UploadFile(fileType upload.FileType, user string, file multipart.File, fileHeader *multipart.FileHeader) (*FileInfo, error) {
	fileName := upload.GetFileName(user, fileHeader.Filename)
	uploadSavePath := upload.GetSavePath()
	dst := uploadSavePath + "/" + fileName

	if !upload.CheckContainExt(fileType, fileName) {
		return nil, errors.New("file suffix is not supported")
	}
	if upload.CheckSavePath(uploadSavePath) {
		if err := upload.CreateSavePath(uploadSavePath, os.ModePerm); err != nil {
			return nil, errors.New("fail to create save directory")
		}
	}
	if upload.CheckMaxSize(fileType, file) {
		return nil, errors.New("exceeded maximum file limit")
	}
	if upload.CheckPermission(uploadSavePath) {
		return nil, errors.New("insufficient file permissions")
	}
	if err := upload.SaveFile(fileHeader, dst); err != nil {
		return nil, err
	}

	accessUrl := global.AppSetting.UploadServerUrl + "/" + fileName
	return &FileInfo{Name: fileName, AccessUrl: accessUrl}, nil
}

func (svc *Service) FileDownload(file model.File) (string, error) {
	u := strings.Split(file.Url, "/")
	f, err := os.Open("./" + upload.GetSavePath() + "/" + u[4])
	if err != nil {
		return "", err
	}
	defer f.Close()

	var l = limiter.NewMethodLimiter().AddBuckets(limiter.LimiterBucketRule{
		Key:          "download",
		FillInterval: time.Duration(global.FileSetting.LimitRate * 1024),
		Capacity:     1,
		Quantum:      1,
	})
	bucket, _ := l.GetBucket("download")

	start := time.Now()
	filename := url.QueryEscape(file.Name) // 防止中文乱码
	out, err := os.Create(global.FileSetting.DownPath + filename)
	if err != nil {
		return "", err
	}
	defer out.Close()
	n, err := io.Copy(ratelimit.Writer(out, bucket), f)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Copied %d bytes in %s\n", n, time.Since(start)), nil
}
