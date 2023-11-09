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
		Request: request,
		Context: context,
		SLogger: variable.ZapSugarLog,
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
	if !ok {
		return fmt.Sprintf(prefixTemplate, "UnknownFunction", 0)
	} else {
		return fmt.Sprintf(prefixTemplate, runtime.FuncForPC(pc).Name(), codeLine)
	}
}

func (l *RequestLogger) InterfaceLog(api string, cost int64, success bool, internal bool) {
	successInt := 0
	internalInt := 0
	if success {
		successInt = 1
	}
	if internal {
		internalInt = 0
	}
	l.SLogger.With(zap.String("trace_id", l.Context.TraceId), zap.String("authorization", l.Context.Authorization),
		zap.Error(nil), zap.String("request", l.requestString), zap.String("breakpoint", l.GetCallerPrefix()),
		zap.String("api", api), zap.String("log_type", "interface_log"), zap.Int("cost", int(cost)),
		zap.Int("success", successInt), zap.Int("internal", internalInt)).
		Infof("Api")
}

func (l *RequestLogger) Debug(err error, template string, args ...interface{}) {
	l.SLogger.With(zap.String("trace_id", l.Context.TraceId), zap.String("authorization", l.Context.Authorization),
		zap.Error(err), zap.String("request", l.requestString), zap.String("breakpoint", l.GetCallerPrefix())).
		Debugf(template, args...)
}

func (l *RequestLogger) Info(err error, template string, args ...interface{}) {
	l.SLogger.With(zap.String("trace_id", l.Context.TraceId), zap.String("authorization", l.Context.Authorization),
		zap.Error(err), zap.String("request", l.requestString), zap.String("breakpoint", l.GetCallerPrefix())).
		Infof(template, args...)
}

func (l *RequestLogger) Warn(err error, template string, args ...interface{}) {
	l.SLogger.With(zap.String("trace_id", l.Context.TraceId), zap.String("authorization", l.Context.Authorization),
		zap.Error(err), zap.String("request", l.requestString), zap.String("breakpoint", l.GetCallerPrefix())).
		Warnf(template, args...)
}

func (l *RequestLogger) Error(err error, template string, args ...interface{}) {
	l.SLogger.With(zap.String("trace_id", l.Context.TraceId), zap.String("authorization", l.Context.Authorization),
		zap.Error(err), zap.String("request", l.requestString), zap.String("breakpoint", l.GetCallerPrefix())).
		Errorf(template, args...)
}

func (l *RequestLogger) Panic(err error, template string, args ...interface{}) {
	l.SLogger.With(zap.String("trace_id", l.Context.TraceId), zap.String("authorization", l.Context.Authorization),
		zap.Error(err), zap.String("request", l.requestString), zap.String("breakpoint", l.GetCallerPrefix())).
		Panicf(template, args...)
}
