package static

import (
	"fmt"
	"github.com/wuzehv/ginframework/util/journal"
)

type Code int

// 状态码
const (
	Success Code = iota
	// 外部状态码
	ParamsError
	SignatureError
	UserDisabled
	TokenNotExists
	TokenParseError
	SessionExpired
	SessionNotExists
	SessionStatusNotLogin
	ClientNotExists
	ClientDisabled
	SystemError
	// 内部状态码
	UsernamePasswdNotMatch
	UsernamePasswdFailNumOut
	UserNotLogin
)

var errors = [...]string{
	Success:                  "success",
	ParamsError:              "params error",
	SignatureError:           "signature error",
	UserDisabled:             "user disabled",
	TokenNotExists:           "token not exists",
	TokenParseError:          "token parse exists",
	SessionExpired:           "session expired",
	SessionNotExists:         "session not exists",
	SessionStatusNotLogin:    "session status not login",
	ClientNotExists:          "client not exists",
	ClientDisabled:           "client disabled",
	SystemError:              "system error",
	UsernamePasswdNotMatch:   "用户名密码错误",
	UsernamePasswdFailNumOut: "失败次数过多",
	UserNotLogin:             "用户未登录",
}

// 通用key
const (
	Domain     = "domain"
	Jump       = "jump"
	CookieFlag = "flag"
	Token      = "token"
	Client     = "client"
	Uid        = "uid"
	Session    = "session"
	User       = "user"
	Timestamp  = "timestamp"
	Sign       = "sign"
)

type SvcRequest struct {
	Token     string `form:"token" binding:"required"`            // 登录token
	Domain    string `form:"domain" binding:"required,unix_addr"` // 客户端域名
	Timestamp string `form:"timestamp" binding:"required"`        // 时间戳
	Sign      string `form:"sign" binding:"required"`             // 签名
}

// 响应结构体
type Response struct {
	Code    Code        `json:"-"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (c Code) Msg(data interface{}) *Response {
	if c != Success {
		if e, ok := data.(error); ok {
			data = e.Error()
		}

		journal.Error("response", data)

		// 系统错误屏蔽前端输出
		if c == SystemError {
			data = nil
		}
	}

	return &Response{
		Code:    c,
		Message: c.Error(),
		Data:    data,
	}
}

func (c Code) Error() string {
	if int(c) >= len(errors) {
		return fmt.Sprintf("code: %d, msg: %v", SystemError, errors[SystemError])
	}

	return fmt.Sprintf("code: %d, msg: %v", c, errors[c])
}
