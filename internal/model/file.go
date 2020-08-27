/**
 * @Title  file
 * @description  #
 * @Author  沈来
 * @Update  2020/8/26 16:13
 **/
package model

import (
	"CloudDisk/pkg/app"
	"github.com/jinzhu/gorm"
)

type File struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
	Url   string `json:"url"`
	Type  string `json:"type"`
}

func (f File) TableName() string {
	return "cloud_file"
}

type FileSwagger struct {
	List  []*File
	Pager *app.Pager
}

func (f File) Count(db *gorm.DB) (int, error) {
	var count int
	if f.Name != "" {
		//别忘 ? 不然sql语句会拼接错误
		db = db.Where("name = ?", f.Name)
	}
	db = db.Where("state = ?", f.State)
	if err := db.Model(&f).Where("is_del = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (f File) List(db *gorm.DB, pageOffSet, pageSize int) ([]*File, error) {
	var files []*File
	var err error
	if pageOffSet >= 0 && pageSize > 0 {
		db = db.Offset(pageOffSet).Limit(pageSize)
	}
	if f.Name != "" {
		db = db.Where("name = ?", f.Name)
	}
	db = db.Where("state = ?", f.State)
	if err = db.Where("is_del = ?", 0).Find(&files).Error; err != nil {
		return nil, err
	}
	return files, nil
}

func (f File) ListByIDs(db *gorm.DB, ids []uint32) ([]*File, error) {
	var files []*File
	err := db.Where("state = ? AND is_del = ?", f.State, 0).Where("id IN (?)", ids).Find(&files).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return files, nil
}

func (f File) Get(db *gorm.DB) (File, error) {
	var file File
	err := db.Where("name = ? AND created_by = ? AND is_del = ? AND state = ?", f.Name, f.CreatedBy, 0, f.State).First(&file).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return file, err
	}

	return file, nil
}

func (f File) Create(db *gorm.DB) error {
	return db.Create(&f).Error
}

func (f File) Update(db *gorm.DB, values interface{}) error {
	if err := db.Model(f).Updates(values).Where("id = ? AND is_del = ?", f.ID, 0).Error; err != nil {
		return err
	}
	return nil
}

func (f File) Delete(db *gorm.DB) error {
	return db.Where("id = ? AND is_del = ?", f.Model.ID, 0).Delete(&f).Error
}
