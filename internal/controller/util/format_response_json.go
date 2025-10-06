package util

import "github.com/gin-gonic/gin"

func FormatJsonResponseHttp(result any, err error) map[string]any {
	if err != nil {
		return gin.H{
			"result": result,
			"error":  err.Error(),
		}
	} else {
		return gin.H{
			"result": result,
			"error":  nil,
		}
	}

}
