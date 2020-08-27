/**
 * @Title  standard
 * @Description  日志标准化
 * @Author  沈来
 * @Update  2020/8/3 21:50
 **/
package logger

import (
	"context"
	"fmt"
	"io"
	"log"
	"runtime"
)

type Logger struct{
	newLogger *log.Logger
	ctx       context.Context
	level     Level
	fields    Fields
	callers   []string
}

func NewLogger(w io.Writer,prefix string,flag int) *Logger{
	l := log.New(w, prefix, flag)
	return &Logger{newLogger: l}
}

func (l *Logger) clone() *Logger{
	nl := *l
	return &nl
}

func (l *Logger) WithLevel(lvl Level) *Logger{
	ll := l.clone()
	ll.level = lvl
	return ll
}

func (l *Logger) WithFields(f Fields) *Logger{
	ll := l.clone()
	if ll.fields == nil{
		ll.fields = make(Fields)
	}
	for k, v := range f{
		ll.fields[k] = v
	}
	return ll
}

func (l *Logger) WithContext(ctx context.Context) *Logger{
	ll := l.clone()
	ll.ctx = ctx
	return ll
}

func (l *Logger) WithCaller(skip int) *Logger{
	ll := l.clone()
	pc, file, line, ok := runtime.Caller(skip)
	if ok {
		f := runtime.FuncForPC(pc)
		ll.callers = []string{ fmt.Sprintf("%s: %d %s", file, line, f.Name())}
	}
	return ll
}

func (l *Logger) WithCallersFrames() *Logger{
	maxCallerDepth := 25
	minCallerDepth := 1
	var callers []string
	pcs := make([]uintptr,maxCallerDepth)
	depth := runtime.Callers(minCallerDepth,pcs)
	frames := runtime.CallersFrames(pcs[:depth])
	for frame, more := frames.Next(); more; frame, more = frames.Next(){
		callers = append(callers, fmt.Sprintf("%s: %d %s",frame.File, frame.Line, frame.Function))
		if !more{
			break
		}
	}

	ll := l.clone()
	ll.callers = callers
	return ll
}