/**
 * @Title  file
 * @description  #
 * @Author  沈来
 * @Update  2020/8/26 15:37
 **/
package service

import (
	"CloudDisk/internal/model"
	"CloudDisk/pkg/app"
)

type FileRequest struct {
	Name      string `form:"name" binding:"required,min=1,max=100"`
	CreatedBy string `form:"created_by" binding:"required,min=3,max=100"`
	State     uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type FileSeekRequest struct {
	Name      string `form:"name" binding:"required,min=1,max=100"`
	CreatedBy string `form:"created_by" binding:"required,min=3,max=100"`
	State     uint8  `form:"state,default=1" binding:"oneof=0 1"`
	Max       string `form:"max" binding:"required"`
	Min       string `form:"min,default=0"`
	Rate      int    `form:"rate,default=100" binding:"max=1024"` //这里有限速，注意
}

type FileListRequest struct {
	Type  string `form:"type" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateFileRequest struct {
	Name      string `form:"name" binding:"required,min=1,max=100"`
	CreatedBy string `form:"created_by" binding:"required,min=3,max=100"`
	State     uint8  `form:"state,default=1" binding:"oneof=0 1"`
	Url       string `form:"name" binding:"max=100"`
	Type      string `form:"name" binding:"max=100"`
}

type DeleteFileRequest struct {
	Name      string `form:"name" binding:"required,min=1,max=100"`
	CreatedBy string `form:"created_by" binding:"required,min=3,max=100"`
}

func (svc *Service) GetFile(param *FileRequest) (model.File, error) {
	return svc.dao.GetFile(param.Name, param.CreatedBy, param.State)
}

func (svc *Service) CountFile(user string, param *FileListRequest, delete uint32) (int, error) {
	return svc.dao.CountFile(user, param.Type, param.State, delete)
}

func (svc *Service) GetFileList(user string, param *FileListRequest, pager *app.Pager, delete uint32) ([]*model.File, error) {
	return svc.dao.GetFileList(user, param.Type, param.State, pager.Page, pager.PageSize, delete)
}

func (svc *Service) CreateFile(param *CreateFileRequest) error {
	return svc.dao.CreateFile(param.Name, param.State, param.CreatedBy, param.Url, param.Type)
}

func (svc *Service) CreateFileWithUser(user string, param *model.File) error {
	return svc.dao.CreateFileWithUser(user, param.Name, param.State, param.CreatedBy, param.Url, param.Type)
}

func (svc *Service) DeleteFile(user string, param *DeleteFileRequest) error {
	return svc.dao.DeleteFile(user, param.Name, param.CreatedBy)
}
