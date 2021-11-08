package api

import (
	"github.com/Wilddogmoto/example_project/logging"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authHeader = "Authorization"
)

func userIdentification(c *gin.Context) {

	var (
		logger = logging.InitLogger()
		err    error
		ok     bool
		token  *jwt.Token
		cl     *TokenClaims
	)

	header := c.GetHeader(authHeader) // получение значения из хедера Authorized
	if header == "" {
		logger.Warn("empty authorized header")
		sendResponse(8, c)
		return
	}

	headerSplit := strings.Split(header, " ") // убераем "Breare"
	if len(headerSplit) != 2 {
		logger.Warn("invalid authorized header")
		sendResponse(9, c)
		return
	}

	token, err = jwt.ParseWithClaims(headerSplit[1], &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok = token.Method.(*jwt.SigningMethodHMAC); !ok { // возвращается ключ подпись
			logger.Errorf("invalid signature method: %v", err)
			return nil, err
		}
		return []byte(signaturekey), nil // ключ зашифровки \ расшифровки
	})
	if err != nil {
		logger.Errorf("error ParseWithClaims: %v", err)
		sendResponse(11, c)
		return
	}

	cl, ok = token.Claims.(*TokenClaims) // проверяем клэйм токина
	if !ok {
		logger.Warn("token claims if not a type TokenClaims")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"userId":        cl.UserId,
		"Authorization": "success",
	})

	logger.Infof("token authorization was successful for user_id: %v", cl.UserId)
}
