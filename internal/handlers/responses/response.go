package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Message struct {
	Status     string `json:"status"`
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Data       any    `json:"data"`
}

func SuccessResponse(c *gin.Context, message string, data any) {
	c.JSON(http.StatusOK, Message{
		Status:     "success",
		StatusCode: http.StatusOK,
		Message:    message,
		Data:       data,
	})
}

func FailedResponse(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, Message{
		Status:     "error",
		StatusCode: statusCode,
		Message:    message,
	})
}
