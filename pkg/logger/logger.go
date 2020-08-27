/**
 * @Title  logger
 * @Description  日志分级和输出
 * @Author  沈来
 * @Update  2020/8/3 21:45
 **/
package logger

import (
	"context"
	"fmt"
)

type Level int8

type Fields map[string]interface{}

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelPanic
)

func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "debug"
	case LevelInfo:
		return "info"
	case LevelWarn:
		return "warn"
	case LevelError:
		return "error"
	case LevelFatal:
		return "fatal"
	case LevelPanic:
		return "panic"
	}
	return ""
}

func (l *Logger) Debug(ctx context.Context, v ...interface{}) {
	l.WithLevel(LevelDebug).WithContext(ctx).WithTrace().Output(fmt.Sprint(v...))
}

func (l *Logger) Debugf(ctx context.Context, format string, v ...interface{}) {
	l.WithLevel(LevelDebug).WithContext(ctx).WithTrace().Output(fmt.Sprintf(format, v...))
}

func (l *Logger) Info(ctx context.Context, v ...interface{}) {
	l.WithLevel(LevelInfo).WithContext(ctx).WithTrace().Output(fmt.Sprint(v...))
}

func (l *Logger) Infof(ctx context.Context, format string, v ...interface{}) {
	l.WithLevel(LevelInfo).WithContext(ctx).WithTrace().Output(fmt.Sprintf(format, v...))
}

func (l *Logger) Warn(v ...interface{}) {
	l.WithLevel(LevelWarn).Output(fmt.Sprint(v...))
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	l.WithLevel(LevelWarn).Output(fmt.Sprintf(format, v...))
}

func (l *Logger) Error(v ...interface{}) {
	l.WithLevel(LevelError).Output(fmt.Sprint(v...))
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.WithLevel(LevelError).Output(fmt.Sprintf(format, v...))
}

func (l *Logger) Fatal(v ...interface{}) {
	l.WithLevel(LevelFatal).Output(fmt.Sprint(v...))
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.WithLevel(LevelFatal).Output(fmt.Sprintf(format, v...))
}

func (l *Logger) Panic(v ...interface{}) {
	l.WithLevel(LevelPanic).Output(fmt.Sprint(v...))
}

func (l *Logger) Panicf(format string, v ...interface{}) {
	l.WithLevel(LevelPanic).Output(fmt.Sprintf(format, v...))
}
