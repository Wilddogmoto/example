package api

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authHeader = "Authorization"

	signaturekey = "full" // ключ зашифровки \ расшифровки
)

func userIdentification(c *gin.Context) {

	header := c.GetHeader(authHeader) // получение значения из хедера Authorized
	if header == "" {
		sendResponse(8, c)
		return
	}

	headerSplit := strings.Split(header, " ") // убераем "Breare"
	if len(headerSplit) != 2 {
		sendResponse(9, c)
		return
	}

	userId, err := parseToken(headerSplit[1])
	if err != nil {

		sendResponse(9, c)
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"userId":        userId,
		"Authorization": "success",
	})
}

func parseToken(incomingToken string) (int, error) {

	token, err := jwt.ParseWithClaims(incomingToken, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok { // возвращается ключ подпись

			return nil, errors.New("invalid signature method")
		}

		return []byte(signaturekey), nil // ключ зашифровки \ расшифровки
	})

	if err != nil {
		return 0, err
	}

	cl, ok := token.Claims.(*TokenClaims)
	if !ok {
		fmt.Println("error token claims")
		return 0, errors.New("token claims if not a type TokenClaims")
	}

	return cl.UserId, nil
}
