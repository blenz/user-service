package service

import (
	"strconv"
)

// TODO: consider making a struct
func error(code int, message string) map[string]map[string]string {
	return map[string]map[string]string{
		"error": map[string]string{
			"code": strconv.Itoa(code),
			"message": message,
		},
	}
}
