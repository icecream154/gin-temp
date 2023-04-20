package consts

// 这里定义的常量，一般是具有错误代码+错误说明组成，一般用于接口返回
const (
	RequestKey string = "Request"

	// 进程被结束
	ProcessKilled string = "收到信号，进程被结束"
	// 表单验证器前缀
	ValidatorPrefix              string = "Form_Validator_"
	ValidatorParamsCheckFailCode int    = -400300
	ValidatorParamsCheckFailMsg  string = "参数校验失败"

	//服务器代码发生错误
	ServerOccurredErrorCode int    = -500100
	ServerOccurredErrorMsg  string = "内部服务错误, "

	// token相关
	JwtTokenOK            int    = 200100  //token有效
	JwtTokenInvalid       int    = -400100 //无效的token
	JwtTokenInvalidMsg    string = "token 无效"
	JwtTokenExpired       int    = -400101          //过期的token
	JwtTokenFormatErrCode int    = -400102          //提交的 token 格式错误
	JwtTokenFormatErrMsg  string = "提交的 token 格式错误" //提交的 token 格式错误

	JwtTokenInvalidTypeCode int    = -400100 //无效的token类型
	JwtTokenInvalidTypeMsg  string = "非法访问"

	//SnowFlake 雪花算法
	StartTimeStamp = int64(1483228800000) //开始时间截 (2017-01-01)
	MachineIdBits  = uint(10)             //机器id所占的位数
	SequenceBits   = uint(12)             //序列所占的位数
	//MachineIdMax   = int64(-1 ^ (-1 << MachineIdBits)) //支持的最大机器id数量
	SequenceMask   = int64(-1 ^ (-1 << SequenceBits)) //
	MachineIdShift = SequenceBits                     //机器id左移位数
	TimestampShift = SequenceBits + MachineIdBits     //时间戳左移位数

	// CURD 常用业务状态码
	CurdStatusOkCode         int    = 200
	CurdStatusOkMsg          string = "Success"
	CurdCreateFailCode       int    = -400200
	CurdCreateFailMsg        string = "新增失败"
	CurdUpdateFailCode       int    = -400201
	CurdUpdateFailMsg        string = "更新失败"
	CurdDeleteFailCode       int    = -400202
	CurdDeleteFailMsg        string = "删除失败"
	CurdSelectFailCode       int    = -400203
	CurdSelectFailMsg        string = "查询失败"
	CurdRegisterFailCode     int    = -400204
	CurdRegisterFailMsg      string = "注册失败"
	CurdLoginFailCode        int    = -400205
	CurdLoginFailMsg         string = "登录失败"
	CurdRefreshTokenFailCode int    = -400206
	CurdRefreshTokenFailMsg  string = "刷新Token失败"

	//文件上传
	FilesUploadFailCode            int    = -400250
	FilesUploadFailMsg             string = "文件上传失败, 获取上传文件发生错误!"
	FilesUploadMoreThanMaxSizeCode int    = -400251
	FilesUploadMoreThanMaxSizeMsg  string = "长传文件超过系统设定的最大值,系统允许的最大值（M）："
	FilesUploadMimeTypeFailCode    int    = -400252
	FilesUploadMimeTypeFailMsg     string = "文件mime类型不允许"

	//websocket
	WsServerNotStartCode int    = -400300
	WsServerNotStartMsg  string = "websocket 服务没有开启，请在配置文件开启，相关路径：config/config_dev.yml"
	WsOpenFailCode       int    = -400301
	WsOpenFailMsg        string = "websocket open阶段初始化基本参数失败"

	PhoneCodeInvalidCode     int    = -400400
	PhoneCodeInvalidMsg      string = "手机验证码无效"
	PhoneCodeExpiredCode     int    = -400401
	PhoneCodeExpiredMsg      string = "手机验证码过期"
	PhoneCodeCheckFailedCode int    = -400402
	PhoneCodeCheckFailedMsg  string = "手机验证码校验失败"
)
