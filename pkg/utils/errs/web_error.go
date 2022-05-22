package errs

import "github.com/XC-Zero/yinwan/pkg/utils/logger"

func CreateWebErrorMsg(errorMsg string, otherInfo ...interface{}) map[string]interface{} {
	logger.Waring(nil, errorMsg)
	return map[string]interface{}{
		"status":    "error",
		"message":   errorMsg,
		"otherInfo": otherInfo,
	}
}

func CreateSuccessMsg(msg string, otherInfo ...interface{}) map[string]interface{} {
	logger.Info(msg)

	return map[string]interface{}{
		"status":    "success",
		"message":   msg,
		"otherInfo": otherInfo,
	}
}
