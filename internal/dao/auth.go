/**
 * @Title  auth
 * @description  JWT
 * @Author  沈来
 * @Update  2020/8/6 22:05
 **/
package dao

import "CloudDisk/internal/model"

func (d *Dao) GetAuth(username, password string) (model.Auth, error) {
	auth := model.Auth{Username: username, Password: password}
	return auth.Get(d.engine)
}
