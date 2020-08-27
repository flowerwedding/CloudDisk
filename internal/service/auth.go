/**
 * @Title  auth
 * @description  JWT
 * @Author  沈来
 * @Update  2020/8/6 22:08
 **/
package service

import "errors"

type AuthRequest struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func (svc *Service) CheckAuth(param *AuthRequest) error {
	auth, err := svc.dao.GetAuth(
		param.Username,
		param.Password,
	)
	if err != nil {
		return nil
	}
	if auth.ID > 0 {
		return nil
	}

	return errors.New("auth info does not exist")
}
