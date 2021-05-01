package config

import (
	"github.com/gin-gonic/gin"
	"time"
)

func ResponseSuccess(data interface{}) gin.H {
	//var res map[string]interface{}
	t, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	result := make(map[string]interface{})
	result["status"] = 200
	result["data"] = data
	result["time"] = t
	result["message"] = nil
	return result
}

func ResponseUnauthorized() gin.H {
	t, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	result := make(map[string]interface{})
	result["status"] = 401
	result["data"] = nil
	result["message"] = "unauthorized request"
	result["time"] = t
	return result
}

func ResponseBadRequest(message string) gin.H {
	t, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	result := make(map[string]interface{})
	result["status"] = 400
	result["message"] = "Bad Request"
	if message != "" {
		result["message"] = message
	}
	result["time"] = t
	result["data"] = nil
	return result
}

func ResponseNotFound(message string) gin.H {
	t, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	result := make(map[string]interface{})
	result["status"] = 400
	result["message"] = "Data not found"
	if message != "" {
		result["message"] = message
	}
	result["time"] = t
	result["data"] = nil
	return result
}

func ResponseServerError(message string) gin.H {
	t, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	result := make(map[string]interface{})
	result["status"] = 500
	result["message"] = "Server Internal Error"
	if message != "" {
		result["message"] = message
	}
	result["time"] = t
	result["data"] = nil
	return result
}
