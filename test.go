package main

import (
	log "github.com/sirupsen/logrus"
	"net/http"
)

func init() {

}
func main() {

	log.SetFormatter(&log.JSONFormatter{})

	log.WithFields(log.Fields{
		"method":  "POST",
		"status":  http.StatusOK,
		"source":  "file: api/login.go",
		"func":    "api.loginUser",
		"request": "/auth/login",
	}).Info("login user to service")
}
