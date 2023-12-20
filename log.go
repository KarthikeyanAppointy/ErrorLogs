package logs

import (
	"go.uber.org/zap"
	"strconv"
)

type LogType int

const (
	Default LogType = iota
	Info
	Warning
	Error
	Debug
)

var (
	InfoCodes  = []int{100, 101, 200, 201, 202, 203, 204, 205, 206}
	ErrorCodes = []int{300, 301, 302, 303, 304, 305, 307, 400, 401, 402, 403, 404, 405, 406, 407, 408, 409,
		410, 411, 412, 413, 414, 415, 416, 417, 418, 421, 422, 423, 424, 426, 428, 429, 431, 451, 500, 501,
		502, 503, 504, 505, 506, 507, 508, 510, 511}
	WarnCodes   = []int{}
	DebugCodes  = []int{}
	statusCodes = map[int]string{
		100: "Continue",
		101: "Switching Protocols",
		200: "OK",
		201: "Created",
		202: "Accepted",
		203: "Non-Authoritative Information",
		204: "No Content",
		205: "Reset Content",
		206: "Partial Content",
		300: "Multiple Choices",
		301: "Moved Permanently",
		302: "Found",
		303: "See Other",
		304: "Not Modified",
		305: "Use Proxy",
		307: "Temporary Redirect",
		400: "Bad Request",
		401: "Unauthorized",
		402: "Payment Required",
		403: "Forbidden",
		404: "Not Found",
		405: "Method Not Allowed",
		406: "Not Acceptable",
		407: "Proxy Authentication Required",
		408: "Request Timeout",
		409: "Conflict",
		410: "Gone",
		411: "Length Required",
		412: "Precondition Failed",
		413: "Payload Too Large",
		414: "URI Too Long",
		415: "Unsupported Media Type",
		416: "Range Not Satisfiable",
		417: "Expectation Failed",
		418: "I'm a teapot",
		421: "Misdirected Request",
		422: "Unprocessable Entity",
		423: "Locked",
		424: "Failed Dependency",
		426: "Upgrade Required",
		428: "Precondition Required",
		429: "Too Many Requests",
		431: "Request Header Fields Too Large",
		451: "Unavailable For Legal Reasons",
		500: "Internal Server Error",
		501: "Not Implemented",
		502: "Bad Gateway",
		503: "Service Unavailable",
		504: "Gateway Timeout",
		505: "HTTP Version Not Supported",
		506: "Variant Also Negotiates",
		507: "Insufficient Storage",
		508: "Loop Detected",
		510: "Not Extended",
		511: "Network Authentication Required",
	}
)

type LogErrorRequest struct {
	AppName      string
	ErrorCode    int
	ErrorMessage string
	EventType    string
	LogType      LogType
	Information  map[string]string
}

// LogResponse - Just for testing Purpose
type LogResponse struct {
	LogType LogType
	msg     string
}

func Log(in *LogErrorRequest) LogResponse {

	if in.LogType == Default {
		in.LogType = getDefaultLogType(in.ErrorCode)
	}

	msg := in.AppName + " - Unable to " + in.EventType + " : " + strconv.Itoa(in.ErrorCode) + " - " + in.ErrorMessage

	logInfo := make([]zap.Field, 0)
	for key, value := range in.Information {
		logInfo = append(logInfo, zap.String(key, value))
	}

	switch in.LogType {
	case Info:
		zap.L().Info(msg, logInfo...)
		return LogResponse{
			LogType: Info,
			msg:     msg,
		}
	case Warning:
		zap.L().Warn(msg, logInfo...)
		return LogResponse{
			LogType: Warning,
			msg:     msg,
		}
	case Error:
		zap.L().Error(msg, logInfo...)
		return LogResponse{
			LogType: Error,
			msg:     msg,
		}
	case Debug:
		zap.L().Debug(msg, logInfo...)
		return LogResponse{
			LogType: Debug,
		}
	default:
		panic("Invalid Log Type")
	}
	return LogResponse{
		LogType: Default,
		msg:     "",
	}
}

func getDefaultLogType(errorCode int) LogType {
	if isErrorCodePresent(InfoCodes, errorCode) {
		return Info
	} else if isErrorCodePresent(ErrorCodes, errorCode) {
		return Error
	} else if isErrorCodePresent(WarnCodes, errorCode) {
		return Warning
	}
	// Debug is avoided in Default.
	return Info
}

func isErrorCodePresent(errorCodes []int, code int) bool {
	for _, errorCode := range errorCodes {
		if errorCode == code {
			return true
		}
	}
	return false
}
