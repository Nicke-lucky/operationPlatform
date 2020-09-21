package types

const (
	StatusContinue           = 100 // RFC 7231, 6.2.1
	StatusSwitchingProtocols = 101 // RFC 7231, 6.2.2
	StatusProcessing         = 102 // RFC 2518, 10.1
	StatusEarlyHints         = 103 // RFC 8297

	StatusSuccessfully = 200 //成功

	StatusGETRedisError = 600 //getredis 错误
	StatusSETRedisError = 601 //setredis 错误
	StatusDELRedisError = 602 //setredis 错误

	StatusRepeatedRegistration = 4001 //注册
	StatusPleaseRegister       = 4002 //请先注册
	StatusPasswordError        = 4003 //密码错误,请重新输入
	StatusNoVerificationcode   = 4019 //密码错误,请重新输入

	StatusExportExcelError = 6000
)

var statusText = map[int]string{
	StatusContinue:             "Continue",
	StatusSwitchingProtocols:   "Switching Protocols",
	StatusProcessing:           "Processing",
	StatusEarlyHints:           "Early Hints",
	StatusRepeatedRegistration: "Repeated Registration",
	StatusPleaseRegister:       "Please Register",
	StatusPasswordError:        "Password Error",
	StatusSuccessfully:         "QuerySuccess",
	StatusExportExcelError:     "Export Excel Error",
	StatusGETRedisError:        "GETRedis Error",
	StatusSETRedisError:        "SETRedis Error",
	StatusDELRedisError:        "DELRedis Error",
	StatusNoVerificationcode:   "No Verificationcode",
}

// StatusText returns a text for the HTTP status code. It returns the empty
// string if the code is unknown.
func StatusText(code int) string {
	return statusText[code]
}
