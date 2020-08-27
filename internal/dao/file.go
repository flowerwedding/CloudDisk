/**
 * @Title  file
 * @description  #
 * @Author  沈来
 * @Update  2020/8/26 16:05
 **/
package dao

import (
	"CloudDisk/internal/model"
	"errors"
)

/*
func (d *Dao) CountFile(name string, state uint8) (int, error) {
	file := model.File{Name: name, State: state}
	return file.Count(d.engine)
}*/

func (d *Dao) GetFile(name string, createdBy string, state uint8) (model.File, error) {
	file := model.File{Name: name, Model: &model.Model{CreatedBy: createdBy}, State: state}
	return file.Get(d.engine)
}

/*
func (d *Dao) GetFileList(name string, state uint8, page, pageSize int) ([]*model.File, error) {
	file := model.File{Name: name, State: state}
	pageOffset := app.GetPageOffset(page, pageSize)
	return file.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) GetFileListByIDs(ids []uint32, state uint8) ([]*model.File, error) {
	file := model.File{State: state}
	return file.ListByIDs(d.engine, ids)
}
*/
//param.Name, param.State, param.CreatedBy, param.Url, param.Type
func (d *Dao) CreateFile(name string, state uint8, createdBy string, url string, a string) error {
	file := model.File{
		Name:  name,
		State: state,
		Model: &model.Model{CreatedBy: createdBy},
		Url:   url,
		Type:  a,
	}

	tt, _ := file.Get(d.engine)
	if tt.ID != 0 {
		return errors.New("this file is exist")
	}

	return file.Create(d.engine)
}

/*
func (d *Dao) DeleteFile(id uint32) error {
	file := model.File{Model: &model.Model{ID: id}}
	return file.Delete(d.engine)
}*/
