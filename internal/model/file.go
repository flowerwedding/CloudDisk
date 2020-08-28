/**
 * @Title  file
 * @description  #
 * @Author  沈来
 * @Update  2020/8/26 16:13
 **/
package model

import (
	"CloudDisk/pkg/app"
	"fmt"
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
	if f.Type != "" {
		db = db.Where("type = ?", f.Type)
	}
	db = db.Where("modified_by = ?", f.ModifiedBy)
	db = db.Where("state = ?", f.State)
	if err := db.Model(&f).Where("is_del = ?", f.DeletedOn).Count(&count).Error; err != nil {
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
	if f.Type != "" {
		db = db.Where("type = ?", f.Type)
	}
	db = db.Where("modified_by = ?", f.ModifiedBy)
	db = db.Where("state = ?", f.State)
	if err = db.Where("is_del = ?", f.DeletedOn).Find(&files).Error; err != nil {
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

func (f File) Delete(db *gorm.DB) error {
	fmt.Println(f.ModifiedBy)
	return db.Where("name = ? AND created_by = ? AND is_del = ? AND modified_by = ?", f.Name, f.CreatedBy, 0, f.ModifiedBy).Delete(&f).Error
}
