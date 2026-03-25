package response

import "github.com/gin-gonic/gin"

func Success(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"success": true,
		"data":    data,
	})
}

func Created(c *gin.Context, msg string, data interface{}) {
	c.JSON(201, gin.H{
		"success": true,
		"message": msg,
		"data":    data,
	})
}
func Error(c *gin.Context, status int, msg string) {
	c.JSON(status, gin.H{
		"success": false,
		"message":   msg,
	})
}

func Paginated(c *gin.Context, data interface{}, total int) {
	c.JSON(200, gin.H{
		"success": true,
		"data":    data,
		"total":   total,
	})
}
