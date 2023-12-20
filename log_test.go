package logs

import (
	"testing"
)

func Test_Log(t *testing.T) {
	tests := []struct {
		name        string
		logRequest  *LogErrorRequest
		logResponse *LogResponse
	}{
		{
			name: "Test logging with default log type for error",
			logRequest: &LogErrorRequest{
				AppName:      "Outlook Calendar",
				ErrorCode:    404,
				ErrorMessage: "Not Found",
				EventType:    "Get Event",
				LogType:      Default, // Using default log type
				Information:  map[string]string{"key": "value"},
			},
			logResponse: &LogResponse{
				LogType: Error,
				msg:     "Outlook Calendar - Unable to Get Event : 404 - Not Found",
			},
		},
		{
			name: "Test logging with default log type for info",
			logRequest: &LogErrorRequest{
				AppName:      "Outlook Calendar",
				ErrorCode:    200,
				ErrorMessage: "Success",
				EventType:    "Get Event",
				LogType:      Default, // Using info log type
				Information:  map[string]string{"key": "value"},
			},
			logResponse: &LogResponse{
				LogType: Info,
				msg:     "Outlook Calendar - Unable to Get Event : 200 - Success",
			},
		},
		{
			name: "Test logging with Info log type",
			logRequest: &LogErrorRequest{
				AppName:      "Outlook Calendar",
				ErrorCode:    200,
				ErrorMessage: "Success",
				EventType:    "Get Event",
				LogType:      Info, // Using info log type
				Information:  map[string]string{"key": "value"},
			},
			logResponse: &LogResponse{
				LogType: Info,
				msg:     "Outlook Calendar - Unable to Get Event : 200 - Success",
			},
		},
		{
			name: "Test logging with 403 as info log type",
			logRequest: &LogErrorRequest{
				AppName:      "Outlook Calendar",
				ErrorCode:    403,
				ErrorMessage: "User Rate Limit Exceeded",
				EventType:    "Get Event",
				LogType:      Info,
				Information:  map[string]string{"key": "value"},
			},
			logResponse: &LogResponse{
				LogType: Info,
				msg:     "Outlook Calendar - Unable to Get Event : 403 - User Rate Limit Exceeded",
			},
		},
		{
			name: "Test logging with 403 as Error log type",
			logRequest: &LogErrorRequest{
				AppName:      "Outlook Calendar",
				ErrorCode:    403,
				ErrorMessage: "Rate limit Exceeded",
				EventType:    "Get Event",
				LogType:      Error, // Using info log type
				Information:  map[string]string{"key": "value"},
			},
			logResponse: &LogResponse{
				LogType: Error,
				msg:     "Outlook Calendar - Unable to Get Event : 403 - Rate limit Exceeded",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actualLog := Log(test.logRequest)

			t.Log(actualLog.msg)

			if actualLog.LogType != test.logResponse.LogType {
				t.Errorf("Expected log type: %v, Got: %v", test.logResponse.LogType, actualLog.LogType)
			}
			if actualLog.msg != test.logResponse.msg {
				t.Errorf("Expected log message: %v, Got: %v", test.logResponse.msg, actualLog.msg)
			}
		})
	}
}
