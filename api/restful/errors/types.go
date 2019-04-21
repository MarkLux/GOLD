package errors

import "fmt"

type RestError struct {
	Code int
	Message string
}

func (e RestError) Error() string {
	return fmt.Sprintf("rest error, code: %d, msg: %s", e.Code, e.Message)
}

// 通用错误

func GenUnknownError() RestError {
	return RestError{500, "系统异常，未知错误"}
}

func GenValidationError(column string) RestError {
	return RestError{1001, fmt.Sprintf("表单验证错误，请检查您输入的%s", column)}
}

// 用户相关

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

// 系统相关
