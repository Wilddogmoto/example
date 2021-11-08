package api

import (
	"github.com/gin-gonic/gin"
)

type customError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var responses = map[int]customError{

	1:  {200, "success"},
	2:  {400, "bad request"},
	3:  {400, "wrong password"},
	4:  {400, "name is taken"},
	5:  {200, "account created"},
	6:  {400, "invalid username"},
	7:  {400, "wrong password"},
	8:  {401, "empty authorized header"},
	9:  {401, "invalid authorized header"},
	10: {500, "Internal error"},
	11: {401, "signature is invalid"},
}

func sendResponse(id int, c *gin.Context) {
	r := responses[id]
	c.JSON(r.Code, r)
}
