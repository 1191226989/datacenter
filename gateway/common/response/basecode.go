package response

// 返回码和错误描述

// 成功返回
const (
	SUCCESS     int    = 200
	SUCCESS_MSG string = "ok"
)

// 用户相关
const (
	LOGIN_PARAM_INCORRECT     int    = 2001
	LOGIN_PARAM_INCORRECT_MSG string = "登录用户名或密码不正确"

	LOGIN_JWTTOKEN_ERROR     int    = 2002
	LOGIN_JWTTOKEN_ERROR_MSG string = "登录jwttoken获取错误"

	LOGIN_USERRPC_ERROR     int    = 2003
	LOGIN_USERRPC_ERROR_MSG string = "登录调用USERRPC错误"

	SIGNUP_PARAM_INCORRECT     int    = 2004
	SIGNUP_PARAM_INCORRECT_MSG string = "注册用户名或密码不正确"
)
