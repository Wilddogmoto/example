package api

import "github.com/gin-gonic/gin"

type customError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var responses = map[int]customError{
	1: {200, "success"},
	2: {400, "second error"},
}

func sendResponse(id int, c *gin.Context) {
	r := responses[id]
	c.JSON(r.Code, r)
}
