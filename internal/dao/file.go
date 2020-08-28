/**
 * @Title  file
 * @description  #
 * @Author  沈来
 * @Update  2020/8/26 16:05
 **/
package dao

import (
	"CloudDisk/internal/model"
	"CloudDisk/pkg/app"
	"errors"
)

func (d *Dao) CountFile(user string, a string, state uint8, delete uint32) (int, error) {
	file := model.File{
		Model: &model.Model{ModifiedBy: user, DeletedOn: delete},
		Type:  a,
		State: state,
	}
	return file.Count(d.engine)
}

func (d *Dao) GetFile(name string, createdBy string, state uint8) (model.File, error) {
	file := model.File{
		Name:  name,
		Model: &model.Model{CreatedBy: createdBy},
		State: state,
	}
	return file.Get(d.engine)
}

func (d *Dao) GetFileList(user string, a string, state uint8, page, pageSize int, delete uint32) ([]*model.File, error) {
	file := model.File{
		Model: &model.Model{ModifiedBy: user, DeletedOn: delete},
		Type:  a,
		State: state,
	}
	pageOffset := app.GetPageOffset(page, pageSize)
	return file.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) CreateFile(name string, state uint8, createdBy string, url string, a string) error {
	file := model.File{
		Name:  name,
		State: state,
		Model: &model.Model{CreatedBy: createdBy, ModifiedBy: createdBy},
		Url:   url,
		Type:  a,
	}

	tt, _ := file.Get(d.engine)
	if tt.ID != 0 {
		return errors.New("this file is exist")
	}

	return file.Create(d.engine)
}

func (d *Dao) CreateFileWithUser(user string, name string, state uint8, createdBy string, url string, a string) error {
	file := model.File{
		Name:  name,
		State: state,
		Model: &model.Model{CreatedBy: createdBy, ModifiedBy: user},
		Url:   url,
		Type:  a,
	}

	return file.Create(d.engine)
}

func (d *Dao) DeleteFile(user string, name string, createdBy string) error {
	file := model.File{
		Name:  name,
		Model: &model.Model{CreatedBy: createdBy, ModifiedBy: user},
	}
	return file.Delete(d.engine)
}
