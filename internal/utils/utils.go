package utils

import (
	"log"
	"net/http"
)

func Err(w http.ResponseWriter, err error, status int) {
	log.Println("error:", err)
	w.WriteHeader(status)
	w.Write([]byte(http.StatusText(status)))
}
