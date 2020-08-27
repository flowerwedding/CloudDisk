/**
 * @Title  dao
 * @description  公共dao
 * @Author  沈来
 * @Update  2020/8/5 15:13
 **/
package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Dao struct {
	engine *gorm.DB
}

func New(engine *gorm.DB) *Dao {
	return &Dao{engine: engine}
}

type Article struct {
	ID            uint32 `json:"id"`
	TagID         uint32 `json:"tag_id"`
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	CreatedBy     string `json:"created_by"`
	ModifiedBy    string `json:"modified_by"`
	State         uint8  `json:"state"`
}
