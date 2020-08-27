/**
 * @Title  enforcer
 * @description  #
 * @Author  沈来
 * @Update  2020/8/25 19:46
 **/
package model

import (
	"CloudDisk/pkg/setting"
	"fmt"
	"github.com/casbin/casbin"
	xormadapter "github.com/casbin/xorm-adapter"
	_ "github.com/go-sql-driver/mysql"
)

func NewEnforcer(databaseSetting *setting.DatabaseSettingS) (*casbin.Enforcer, error) {
	a := xormadapter.NewAdapter(databaseSetting.DBType,
		fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
			databaseSetting.UserName,
			databaseSetting.Password,
			databaseSetting.Host,
			databaseSetting.DBName,
			databaseSetting.Charset,
			databaseSetting.ParseTime,
		),
		true,
	)

	e := casbin.NewEnforcer("./configs/rbac_models.conf", a)

	err := e.LoadPolicy()

	return e, err
}
