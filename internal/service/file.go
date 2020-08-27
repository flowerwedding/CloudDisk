/**
 * @Title  file
 * @description  #
 * @Author  沈来
 * @Update  2020/8/26 15:37
 **/
package service

import "CloudDisk/internal/model"

type FileRequest struct {
	Name      string `form:"name" binding:"max=100"`
	CreatedBy string `form:"created_by" binding:"required,min=3,max=100"`
	State     uint8  `form:"state,default=0" binding:"oneof=0 1"`
}

/*
type CountFileRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type FileListRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}
*/
type CreateFileRequest struct {
	Name      string `form:"name" binding:"required,min=1,max=100"`
	CreatedBy string `form:"created_by" binding:"required,min=3,max=100"`
	State     uint8  `form:"state,default=0" binding:"oneof=0 1"`
	Url       string `form:"name" binding:"max=100"`
	Type      string `form:"name" binding:"max=100"`
}

/*
type DeleteFileRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}
*/

func (svc *Service) GetFile(param *FileRequest) (model.File, error) {
	return svc.dao.GetFile(param.Name, param.CreatedBy, param.State)
}

/*
func (svc *Service) CountFile(param *CountFileRequest) (int, error) {
	return svc.dao.CountTag(param.Name, param.State)
}

func (svc *Service) GetFileList(param *FileListRequest, pager *app.Pager) ([]*model.File, error) {
	return svc.dao.GetFileList(param.Name, param.State, pager.Page, pager.PageSize)
}*/

func (svc *Service) CreateFile(param *CreateFileRequest) error {
	return svc.dao.CreateFile(param.Name, param.State, param.CreatedBy, param.Url, param.Type)
}

/*
func (svc *Service) DeleteFile(param *DeleteFileRequest) error {
	return svc.dao.DeleteFile(param.ID)
}*/
