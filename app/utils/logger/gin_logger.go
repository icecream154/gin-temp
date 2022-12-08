package logger

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"goskeleton/app/global/variable"
	"strings"
)

type AccessLog struct {
	AppName      string `json:"app_name"`
	LogType      string `json:"log_type"`
	ClientIP     string `json:"client_ip"`
	CreatedAt    string `json:"created_at"`
	Method       string `json:"method"`
	Path         string `json:"path"`
	Proto        string `json:"proto"`
	StatusCode   int    `json:"status_code"`
	Latency      int64  `json:"latency"` // in microseconds
	ClientAgent  string `json:"client_agent"`
	ErrorMessage string `json:"error_message"`
}

const accessLogType = "access"

func FormatGinLogs(param gin.LogFormatterParams) string {
	rawPath := param.Path
	questionMarkIndex := strings.Index(rawPath, "?")
	if questionMarkIndex != -1 {
		rawPath = rawPath[0:questionMarkIndex]
	}

	beatLog := AccessLog{
		AppName:      variable.AppName,
		LogType:      accessLogType,
		ClientIP:     param.ClientIP,
		CreatedAt:    param.TimeStamp.Format(variable.DateFormat),
		Method:       param.Method,
		Path:         rawPath,
		Proto:        param.Request.Proto,
		StatusCode:   param.StatusCode,
		Latency:      param.Latency.Microseconds(),
		ClientAgent:  param.Request.UserAgent(),
		ErrorMessage: param.ErrorMessage,
	}

	bytes, _ := json.Marshal(beatLog)
	return string(bytes) + "\n"
}
