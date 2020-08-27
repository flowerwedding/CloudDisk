/**
 * @Title  setting
 * @Description  配置管理
 * @Author  沈来
 * @Update  2020/8/3 20:22
 **/
package global

import (
	"CloudDisk/pkg/logger"
	"CloudDisk/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	JWTSetting      *setting.JWTSettingS
	EmailSetting    *setting.EmailSettingS
	FileSetting     *setting.FileSettingS

	Logger *logger.Logger
)
