package errs

func CreateWebErrorMsg(errorMsg string, otherInfo ...interface{}) map[string]interface{} {
	return map[string]interface{}{
		"error": map[string]interface{}{
			"message":   errorMsg,
			"otherInfo": otherInfo,
		},
	}
}

func CreateSuccessMsg(msg string) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
	}
}
