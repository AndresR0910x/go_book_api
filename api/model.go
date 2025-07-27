package api

import "github.com/gin-gonic/gin"

//Instace book struct
type Book struct {
	ID     uint   `json:"id" gorm:"primarykey"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   string `json:"year"`
}

// Create a func for json response with status, message and Data using strcut tags ""
type JsonResponse struct {
	Status  int `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

// Create a func for ResponseJSON using gin context, status, message and data
func ResponseJSON (c *gin.Context, status int, message string, data any) {
	response := JsonResponse {
		Status: status,
		Message: message,
		Data: data,
	}

	c.JSON(status, response)
}

