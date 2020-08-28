/**
 * @Title  file
 * @description  文件上传下载分享
 * @Author  沈来
 * @Update  2020/8/25 22:08
 **/
package v1

import (
	"CloudDisk/global"
	"CloudDisk/internal/service"
	"CloudDisk/pkg/app"
	"CloudDisk/pkg/convert"
	"CloudDisk/pkg/errcode"
	"CloudDisk/pkg/qtcode"
	"CloudDisk/pkg/upload"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type File struct{}

func NewFile() File {
	return File{}
}

// @Summary  加密分享链接
// @Produce  json
// @Param  name query string true "文件名称"
// @Param  created_by query string true "文件上传者"
// @Param  state query int false "状态" Enums(0, 1) default(1)
// @Success  200 {object} model.FileSwagger "成功"
// @Failure  400 {object} errcode.Error "请求错误"
// @Failure  500 {object} errcode.Error "内部错误"
// @Router  /api/v1/file/link [get]
func (f File) GetByLink(c *gin.Context) {
	param := service.FileRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	file, err := svc.GetFile(&param)
	if err != nil {
		global.Logger.Errorf("svc.GetFile err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetFileFail)
		return
	}

	//文件上传时链接已经MD5加密
	response.ToResponse(file)
	return
}

// @Summary  二维码分享链接
// @Produce  json
// @Param  name query string true "文件名称"
// @Param  created_by query string true "文件上传者"
// @Param  state query int false "状态" Enums(0, 1) default(1)
// @Success  200 {object} model.FileSwagger "成功"
// @Failure  400 {object} errcode.Error "请求错误"
// @Failure  500 {object} errcode.Error "内部错误"
// @Router  /api/v1/file/Qt [get]
func (f File) GetByQt(c *gin.Context) {
	param := service.FileRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	file, err := svc.GetFile(&param)
	if err != nil {
		global.Logger.Errorf("svc.GetFile err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetFileFail)
		return
	}

	t := strconv.FormatInt(time.Now().Unix(), 10)
	add := upload.GetSavePath() + "/" + t + ".png"
	err = qtcode.Qt(file.Url, add)
	if err != nil {
		global.Logger.Errorf("svc.GetFile err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateFileQtFail)
		return
	}

	qt := "http://localhost:8000/static/" + t + ".png"
	response.ToResponse(gin.H{"二维码": qt})
	return
}

// @Summary  文件列表
// @Produce  json
// @Param  type query string false "文件类型"
// @Param  state query int false "状态" Enums(0, 1) default(1)
// @Param  page query int false "页码"
// @Param  page_size query int false "每页数量"
// @Success  200 {object} model.FileSwagger "成功"
// @Failure  400 {object} errcode.Error "请求错误"
// @Failure  500 {object} errcode.Error "内部错误"
// @Router  /api/v1/list [get]
func (f File) List(c *gin.Context) {
	param := service.FileListRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	user, _ := c.Get("user")
	svc := service.New(c.Request.Context())
	rows, err := svc.CountFile(user.(string), &param, 0)
	if err != nil {
		global.Logger.Errorf("svc.GetArticleList err: %v", err)
		response.ToErrorResponse(errcode.ErrorCountFileFail)
		return
	}

	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	files, err := svc.GetFileList(user.(string), &param, &pager, 0)
	if err != nil {
		global.Logger.Errorf("svc.GetArticleList err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetFileListFail)
		return
	}

	response.ToResponseList(files, rows)
	return
}

// @Summary  回收站
// @Produce  json
// @Param  type query string false "文章类型"
// @Param  state query int false "状态" Enums(0, 1) default(1)
// @Param  page query int false "页码"
// @Param  page_size query int false "每页数量"
// @Success  200 {object} model.FileSwagger "成功"
// @Failure  400 {object} errcode.Error "请求错误"
// @Failure  500 {object} errcode.Error "内部错误"
// @Router  /api/v1/list/delete [get]
func (f File) ListInDelete(c *gin.Context) {
	param := service.FileListRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	user, _ := c.Get("user")
	svc := service.New(c.Request.Context())
	rows, err := svc.CountFile(user.(string), &param, 1)
	if err != nil {
		global.Logger.Errorf("svc.GetArticleList err: %v", err)
		response.ToErrorResponse(errcode.ErrorCountFileFail)
		return
	}

	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	files, err := svc.GetFileList(user.(string), &param, &pager, 1)
	if err != nil {
		global.Logger.Errorf("svc.GetArticleList err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetFileListFail)
		return
	}

	response.ToResponseList(files, rows)
	return
}

// @Summary  一次性快传
// @Produce  json
// @Param  file body string true "文件"
// @Param  name body string true "文件名"
// @Param  type body string true "类型"
// @Success  200 {object} app.Response "成功"
// @Failure  400 {object} errcode.Error "请求错误"
// @Failure  500 {object} errcode.Error "内部错误"
// @Router  /api/v1/file [post]
func (f File) Create(c *gin.Context) {
	response := app.NewResponse(c)
	file, fileHeader, err := c.Request.FormFile("file")
	fileType := convert.StrTo(c.PostForm("type")).MustInt()
	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}
	if fileHeader == nil || fileType <= 0 {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}

	user, _ := c.Get("user")
	svc := service.New(c.Request.Context())
	fileInfo, err := svc.UploadFile(upload.FileType(fileType), user.(string), file, fileHeader)
	if err != nil {
		global.Logger.Errorf("svc.UploadFile err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateFileFail.WithDetails(err.Error()))
		return
	}

	param := service.CreateFileRequest{
		Name:      c.PostForm("name"),
		CreatedBy: user.(string),
		Url:       fileInfo.AccessUrl,
		Type:      c.PostForm("type"),
	}
	err = svc.CreateFile(&param)
	if err != nil {
		global.Logger.Errorf("app.CreateFile errs: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateFileFail)
		return
	}

	response.ToResponse(gin.H{
		"message":         "上传成功",
		"file_access_url": fileInfo.AccessUrl,
	})
}

// @Summary  下载限速
// @Produce  json
// @Param  name query string true "文件名称"
// @Param  created_by query string true "文件上传者"
// @Param  state query int false "状态" Enums(0, 1) default(1)
// @Success  200 {array} model.File "成功"
// @Failure  400 {object} errcode.Error "请求错误"
// @Failure  500 {object} errcode.Error "内部错误"
// @Router  /api/v1/file/limit [put]
func (f File) DownloadWithLimit(c *gin.Context) {
	param := service.FileRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	file, err := svc.GetFile(&param)
	if err != nil {
		global.Logger.Errorf("app.UpdateTag errs: %v", errs)
		response.ToErrorResponse(errcode.ErrorDownFileFail)
		return
	}

	user, _ := c.Get("user")
	err = svc.CreateFileWithUser(user.(string), &file)
	if err != nil {
		global.Logger.Errorf("app.UpdateTag errs: %v", errs)
		response.ToErrorResponse(errcode.ErrorDownFileFail)
		return
	}

	message, err := svc.FileDownload(file)
	if err != nil {
		global.Logger.Errorf("app.UpdateTag errs: %v", errs)
		response.ToErrorResponse(errcode.ErrorDownFileFail)
		return
	}

	response.ToResponse(gin.H{"message": message})
	return
}

// @Summary  断点续传
// @Produce  json
// @Param  name query string true "文件名称"
// @Param  created_by query string true "文件上传者"
// @Param  state query int false "状态" Enums(0, 1) default(1)
// @Param  max query string true "文件读取结束位置"
// @Param  min query string false "文件读取起始位置" default(0)
// @Param  rate query int false "文件读取速率" default(100)
// @Success  200 {array} model.File "成功"
// @Failure  400 {object} errcode.Error "请求错误"
// @Failure  500 {object} errcode.Error "内部错误"
// @Router  /api/v1/file/seek [put]
func (f File) DownloadWithSeek(c *gin.Context) {
	param := service.FileSeekRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	file, err := svc.GetFile(&service.FileRequest{
		Name:      param.Name,
		CreatedBy: param.CreatedBy,
		State:     param.State,
	})
	if err != nil {
		global.Logger.Errorf("app.UpdateTag errs: %v", errs)
		response.ToErrorResponse(errcode.ErrorDownFileFail)
		return
	}

	user, _ := c.Get("user")
	err = svc.CreateFileWithUser(user.(string), &file)
	if err != nil {
		global.Logger.Errorf("app.UpdateTag errs: %v", errs)
		response.ToErrorResponse(errcode.ErrorDownFileFail)
		return
	}

	message, err := svc.FileSeekDown(param.Max, param.Min, param.Rate, file)
	if err != nil {
		global.Logger.Errorf("app.UpdateTag errs: %v", errs)
		response.ToErrorResponse(errcode.ErrorDownFileFail)
		return
	}

	response.ToResponse(gin.H{"message": message, "status": "下载成功"})
	return
}

// @Summary  文件删除
// @Produce  json
// @Param  name query string true "文件名称"
// @Param  created_by query string true "文件上传者"
// @Success  200 {string} string "成功"
// @Failure  400 {object} errcode.Error "请求错误"
// @Failure  500 {object} errcode.Error "内部错误""
// @Router  /api/v1/file [delete]
func (f File) Delete(c *gin.Context) {
	param := service.DeleteFileRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	user, _ := c.Get("user")
	err := svc.DeleteFile(user.(string), &param)
	if err != nil {
		global.Logger.Errorf("svc.DeleteArticle err: %v", err)
		response.ToErrorResponse(errcode.ErrorDeleteFileFail)
		return
	}

	response.ToResponse(gin.H{"message": "删除成功"})
	return
}
