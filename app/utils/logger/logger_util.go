package logger

import (
	"fmt"
	"go.uber.org/zap"
	"goskeleton/app/global/variable"
	"goskeleton/app/http/middleware/request_context"
	"runtime"
)

type RequestLogger struct {
	Request       *interface{}
	Context       *request_context.RequestContext
	SLogger       *zap.SugaredLogger
	requestString string
}

func GetLogger(request *interface{}, context *request_context.RequestContext) *RequestLogger {
	requestLogger := RequestLogger{
		Request:      request,
		Context:      context,
		SLogger:      variable.ZapSugarLog,
	}
	if request == nil {
		requestLogger.requestString = "&{ }"
	} else {
		requestLogger.requestString = fmt.Sprintf("%v", *request)
	}
	return &requestLogger
}

func (l *RequestLogger) GetCallerPrefix() string {
	// 获取上层调用者PC，文件名，所在行git
	prefixTemplate := "%s:%d"
	pc, _, codeLine, ok := runtime.Caller(2)
	if !ok{
		return fmt.Sprintf(prefixTemplate, "UnknownFunction", 0)
	} else {
		return fmt.Sprintf(prefixTemplate, runtime.FuncForPC(pc).Name(), codeLine)
	}
}

func (l *RequestLogger) Debug(err error, template string, args ...interface{}) {
	l.SLogger.With(zap.String("trace_id", l.Context.TraceId), zap.String("authorization", l.Context.Authorization),
		zap.Error(err), zap.String("request", l.requestString), zap.String("breakpoint", l.GetCallerPrefix())).
		Debugf( template, args...)
}

func (l *RequestLogger) Info(err error, template string, args ...interface{}) {
	l.SLogger.With(zap.String("trace_id", l.Context.TraceId), zap.String("authorization", l.Context.Authorization),
		zap.Error(err), zap.String("request", l.requestString), zap.String("breakpoint", l.GetCallerPrefix())).
		Infof( template, args...)
}

func (l *RequestLogger) Warn(err error, template string, args ...interface{}) {
	l.SLogger.With(zap.String("trace_id", l.Context.TraceId), zap.String("authorization", l.Context.Authorization),
		zap.Error(err), zap.String("request", l.requestString), zap.String("breakpoint", l.GetCallerPrefix())).
		Warnf( template, args...)
}

func (l *RequestLogger) Error(err error, template string, args ...interface{}) {
	l.SLogger.With(zap.String("trace_id", l.Context.TraceId), zap.String("authorization", l.Context.Authorization),
		zap.Error(err), zap.String("request", l.requestString), zap.String("breakpoint", l.GetCallerPrefix())).
		Errorf( template, args...)
}

func (l *RequestLogger) Panic(err error, template string, args ...interface{}) {
	l.SLogger.With(zap.String("trace_id", l.Context.TraceId), zap.String("authorization", l.Context.Authorization),
		zap.Error(err), zap.String("request", l.requestString), zap.String("breakpoint", l.GetCallerPrefix())).
		Panicf( template, args...)
}
