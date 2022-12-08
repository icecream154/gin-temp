package bootstrap

import (
	"go.uber.org/zap"
	_ "goskeleton/app/core/destroy" // 监听程序退出信号，用于资源的释放
	"goskeleton/app/global/my_errors"
	"goskeleton/app/global/variable"
	"goskeleton/app/http/validator/common/register_validator"
	"goskeleton/app/service/sys_log_hook"
	"goskeleton/app/utils/casbin_v2"
	"goskeleton/app/utils/gorm_v2"
	"goskeleton/app/utils/snow_flake"
	"goskeleton/app/utils/websocket/core"
	"goskeleton/app/utils/yml_config"
	"goskeleton/app/utils/zap_factory"
	"goskeleton/env"
	"log"
	"os"
)

// 检查项目必须的非编译目录是否存在，避免编译后调用的时候缺失相关目录
func checkRequiredFolders() {
	//1.检查配置文件是否存在
	// app 配置
	allEnvs := []string{env.DevEnv, env.TestEnv, env.ProdEnv}
	for i := 0; i < len(allEnvs); i++ {
		if _, err := os.Stat(variable.BasePath + "/config/app/config_" + allEnvs[i] + ".yml"); err != nil {
			log.Fatal(my_errors.ErrorsConfigYamlNotExists + err.Error())
		}
	}
	// db 配置
	for i := 0; i < len(allEnvs); i++ {
		if _, err := os.Stat(variable.BasePath + "/config/db/gorm_v2_" + allEnvs[i] + ".yml"); err != nil {
			log.Fatal(my_errors.ErrorsConfigYamlNotExists + err.Error())
		}
		if _, err := os.Stat(variable.BasePath + "/config/db/gorm_v2_" + allEnvs[i] + "_docker.yml"); err != nil {
			log.Fatal(my_errors.ErrorsConfigYamlNotExists + err.Error())
		}
	}

	//2.检查storage/logs 目录是否存在
	if _, err := os.Stat(variable.BasePath + "/storage/logs/"); err != nil {
		log.Fatal(my_errors.ErrorsStorageLogsNotExists + err.Error())
	}
}

func init() {
	// 1. 初始化 项目根路径，参见 variable 常量包，相关路径：app\global\variable\variable.go

	//2.检查配置文件以及日志目录等非编译性的必要条件
	checkRequiredFolders()

	//3.初始化表单参数验证器，注册在容器（Web、Api共用容器）
	register_validator.WebRegisterValidator()

	// config>gorm_v2.yml 启动文件变化监听事件
	env.AppEnv = os.Getenv("APP_ENV")
	dockerSuffix := ""
	if env.IsDocker() {
		dockerSuffix = "_docker"
	}
	variable.ConfigYml = yml_config.CreateYamlFactory("/app", "config_"+env.GetEnv())
	variable.ConfigGormv2Yml = yml_config.CreateYamlFactory("/db", "gorm_v2_"+env.GetEnv()+dockerSuffix)

	// 4.启动针对配置文件(config.yml、gorm_v2.yml)变化的监听， 配置文件操作指针，初始化为全局变量
	variable.ConfigYml.ConfigFileChangeListen()
	variable.ConfigGormv2Yml.ConfigFileChangeListen()

	// 5.初始化全局日志句柄，并载入日志钩子处理函数
	variable.AppName = variable.ConfigYml.GetString("AppName")
	variable.ZapLog = zap_factory.CreateZapFactory(sys_log_hook.ZapLogHandler)
	variable.ZapSugarLog = variable.ZapLog.Sugar().With(zap.String("app_name", variable.AppName)).
		With(zap.String("log_type", "server"))

	// 6.根据配置初始化 gorm mysql 全局 *gorm.Db
	if variable.ConfigGormv2Yml.GetInt("Gormv2.Mysql.IsInitGolobalGormMysql") == 1 {
		if dbMysql, err := gorm_v2.GetOneMysqlClient(); err != nil {
			log.Fatal(my_errors.ErrorsGormInitFail + err.Error())
		} else {
			variable.GormDbMysql = dbMysql
		}
	}
	// 根据配置初始化 gorm sqlserver 全局 *gorm.Db
	if variable.ConfigGormv2Yml.GetInt("Gormv2.Sqlserver.IsInitGolobalGormSqlserver") == 1 {
		if dbSqlserver, err := gorm_v2.GetOneSqlserverClient(); err != nil {
			log.Fatal(my_errors.ErrorsGormInitFail + err.Error())
		} else {
			variable.GormDbSqlserver = dbSqlserver
		}
	}
	// 根据配置初始化 gorm postgresql 全局 *gorm.Db
	if variable.ConfigGormv2Yml.GetInt("Gormv2.PostgreSql.IsInitGolobalGormPostgreSql") == 1 {
		if dbPostgre, err := gorm_v2.GetOnePostgreSqlClient(); err != nil {
			log.Fatal(my_errors.ErrorsGormInitFail + err.Error())
		} else {
			variable.GormDbPostgreSql = dbPostgre
		}
	}

	// 7.雪花算法全局变量
	variable.SnowFlake = snow_flake.CreateSnowflakeFactory()

	// 8.websocket Hub中心启动
	if variable.ConfigYml.GetInt("Websocket.Start") == 1 {
		// websocket 管理中心hub全局初始化一份
		variable.WebsocketHub = core.CreateHubFactory()
		if Wh, ok := variable.WebsocketHub.(*core.Hub); ok {
			go Wh.Run()
		}
	}

	// 9.casbin 依据配置文件设置参数(IsInit=1)初始化
	if variable.ConfigYml.GetInt("Casbin.IsInit") == 1 {
		var err error
		if variable.Enforcer, err = casbin_v2.InitCasbinEnforcer(); err != nil {
			log.Fatal(err.Error())
		}
	}
}
