/**
 * @Title  auth
 * @description  JWT
 * @Author  沈来
 * @Update  2020/8/6 21:34
 **/
package model

import "github.com/jinzhu/gorm"

type Auth struct {
	*Model
	Username string `json:"username"`
	Password string `json:"password"`
}

func (a Auth) TableName() string {
	return "cloud_auth"
}

func (a Auth) Get(db *gorm.DB) (Auth, error) {
	var auth Auth
	db = db.Where("username = ? AND password = ? AND is_del = ?", a.Username, a.Password, 0)
	err := db.First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return auth, err
	}

	return auth, nil
}
