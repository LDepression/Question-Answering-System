package errcode

var (
	StatusOK           = NewErr(0, "成功")
	ErrParamsNotValid  = NewErr(1001, "参数有误")
	ErrNotFound        = NewErr(1002, "未找到资源")
	ErrServer          = NewErr(1003, "系统错误")
	ErrTooManyRequests = NewErr(1004, "请求过多")
	ErrTimeOut         = NewErr(1005, "请求超时，请稍后再试")
	ErrLogin           = NewErr(1006, "账号或密码错误")
	ErrUsername        = NewErr(1007, "账号长度应该大于6小于12位")
	ErrPassword        = NewErr(1008, "密码长度应该大于6小于12位")
	ErrUsenameExist    = NewErr(1009, "注册失败,用户名已经存在了")
)

var (
	ErrUnauthorizedAuthNotExist  = NewErr(2001, "鉴权失败,无法解析")
	ErrUnauthorizedTokenTimeout  = NewErr(2002, "鉴权失败,Token超时")
	ErrUnauthorizedTokenGenerate = NewErr(2003, "鉴权失败,Token 生成失败")
	ErrInsufficientPermissions   = NewErr(2004, "鉴权失败,权限不足")
	ErrOutTimeRefreshToken       = NewErr(2005, "refreshToken过期")
	ErrGenerateToken             = NewErr(2006, "生成token失败")
	ErrUnCorrentAuthor           = NewErr(2007, "不是正确的作者")
)
