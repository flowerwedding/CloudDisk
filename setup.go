/**
 * @Title  test_pkg
 * @Description  公共组件测试
 * @Author  沈来
 * @Update  2020/8/3 21:47
 **/
package main

import (
	"CloudDisk/global"
	"CloudDisk/internal/model"
	"CloudDisk/pkg/logger"
	"CloudDisk/pkg/setting"
	"CloudDisk/pkg/tracer"
	"flag"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"strings"
	"time"
)

func setupSetting() error {
	s, err := setting.NewSetting(strings.Split(config, ",")...)
	//s, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = s.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Email", &global.EmailSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("File", &global.FileSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	global.JWTSetting.Expire *= time.Hour

	if port != "" {
		global.ServerSetting.HttpPort = port
	}
	if runMode != "" {
		global.ServerSetting.RunMode = runMode
	}

	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupEnforcer() error {
	var err error
	global.Enforcer, err = model.NewEnforcer(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupLogger() error {
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		//生成的日志文件目录
		Filename:  global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)

	return nil
}

func setupTracer() error {
	jaegerTracer, _, err := tracer.NewJaegerTracer("CloudDisk", "127.0.0.1:6831")
	if err != nil {
		return err
	}

	global.Tracer = jaegerTracer

	return nil
}

func setupFlag() error {
	flag.StringVar(&port, "port", "", "启动端口")
	flag.StringVar(&runMode, "mode", "", "启动模式")
	flag.StringVar(&config, "config", "", "指定要使用的配置文件路径")
	flag.Parse()

	return nil
}
