package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ok(c *gin.Context, responseDataMap map[string]interface{}) {
	jsonMap := map[string]interface{}{
		"code": 0,
		"msg":  "",
	}
	jsonMap["data"] = responseDataMap
	c.JSON(http.StatusOK, jsonMap)
	return
}

func Err(c *gin.Context, responseDataMap map[string]interface{}) {
	jsonMap := map[string]interface{}{
		"code": 0,
		"msg":  "",
	}
	jsonMap["data"] = responseDataMap
	c.JSON(http.StatusOK, jsonMap)
	return
}

func Custom(c *gin.Context, responseData interface{}) {
	c.JSON(http.StatusOK, responseData)
	return
}
