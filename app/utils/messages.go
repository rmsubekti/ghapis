package utils

// Message template
func Message(status bool, msg string) map[string]interface{} {
	return map[string]interface{}{
		"status":  status,
		"message": msg,
	}
}
