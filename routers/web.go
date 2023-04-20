package routers

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"goskeleton/app/global/variable"
	"goskeleton/app/http/middleware/cors"
	"goskeleton/app/http/middleware/request_context"
	validatorFactory "goskeleton/app/http/validator/core/factory"
	"goskeleton/app/utils/logger"
	"io"
	"os"
)

// 该路由主要设置 后台管理系统等后端应用路由

func InitWebRouter() *gin.Engine {
	var router *gin.Engine
	// 非调试模式（生产模式） 日志写到日志文件
	if variable.ConfigYml.GetBool("AppDebug") == false {
		//1.将日志写入日志文件
		gin.DisableConsoleColor()
		f, _ := os.Create(variable.BasePath + variable.ConfigYml.GetString("Logs.GinLogName"))
		gin.DefaultWriter = io.MultiWriter(f)
		// 2.如果是有nginx前置做代理，基本不需要gin框架记录访问日志，开启下面一行代码，屏蔽上面的三行代码，性能提升 5%
		//gin.SetMode(gin.ReleaseMode)
		router = gin.New()
		router.Use(gin.Recovery())
		router.Use(gin.LoggerWithFormatter(logger.FormatGinLogs))
	} else {
		// 调试模式，开启 pprof 包，便于开发阶段分析程序性能
		router = gin.Default()
		pprof.Register(router)
	}

	//根据配置进行设置跨域
	if variable.ConfigYml.GetBool("HttpServer.AllowCrossDomain") {
		router.Use(cors.Next())
	}

	//  创建一个后端接口路由组
	backend := router.Group("/yi")
	{
		backend.Use(request_context.CheckRequestContext())
		{
			sysNoAuth := backend.Group("/common")
			{
				sysNoAuth.POST("/login", validatorFactory.Create("AccountLogin"))
				sysNoAuth.POST("/register", validatorFactory.Create("AccountRegister"))
			}

			qaNoAuth := backend.Group("/qa")
			{
				qaNoAuth.POST("/submitInput", validatorFactory.Create("QASubmitInput"))
			}

			//backend.Use(authorization.CheckTokenAuth())
			//{
			//	// 通用接口
			//	commonAuth := backend.Group("/common")
			//	{
			//		opinionAuth := commonAuth.Group("/opinion")
			//		{
			//			opinionAuth.POST("/submitOpinion", validatorFactory.Create("SysSubmitOpinion"))
			//			opinionAuth.GET("/queryOpinion", validatorFactory.Create("SysQueryOpinion"))
			//		}
			//	}
			//}
		}
	}
	return router
}
