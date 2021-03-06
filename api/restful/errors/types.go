package errors

import "fmt"

type RestError struct {
	Code int `json:"code"`
	Message string `json:"message"`
}

func (e RestError) Error() string {
	return fmt.Sprintf("rest error, code: %d, msg: %s", e.Code, e.Message)
}

// common
func GenUnknownError() RestError {
	return RestError{500, "系统异常，未知错误"}
}

func GenSystemError(msg string) RestError {
	return RestError{501, msg}
}

func GenValidationError() RestError {
	return RestError{1001, "表单验证错误，请检查输入"}
}

func GenInvalidParam() RestError {
	return RestError{1002, "参数异常"}
}

// user-related

func GenPwdError() RestError {
	return RestError{2001, "密码错误"}
}

func GenUserNotExistedError() RestError {
	return RestError{2002, "用户不存在"}
}

func GenRegisteredError(column string) RestError {
	return RestError{2003, fmt.Sprintf("该%s已被注册", column)}
}

func GenNeedLoginError() RestError {
	return RestError{2004, "需要登录"}
}

func GenPermissionDeniedError() RestError {
	return RestError{2005, "权限不足"}
}

// function-service related

func GenFunctionServiceExistedError() RestError {
	return RestError{3001, "该服务名已被使用"}
}

func GenFunctionNotFoundError() RestError {
	return RestError{3002, "指定服务不存在"}
}

// operate-log relatd

func GenOperateLogNotFoundError() RestError {
	return RestError{4001, "操作日志不存在"}
}
