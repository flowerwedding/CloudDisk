/**
 * @Title  service
 * @description  公共service
 * @Author  沈来
 * @Update  2020/8/5 15:52
 **/
package service

import (
	"CloudDisk/global"
	"CloudDisk/internal/dao"
	"context"
	otgorm "github.com/EDDYCJY/opentracing-gorm"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.New(otgorm.WithContext(svc.ctx, global.DBEngine))
	return svc
}
