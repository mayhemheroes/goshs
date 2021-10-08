package mylog

import (
	"log"
	"net/http"
)

// LogRequest will log the request in a uniform way
func LogRequest(remoteAddr, method, url, proto string, status int) {
	if status == 999 {
		log.Printf("WEBDAV:  %s - - \"%s %s %s\"", remoteAddr, method, url, proto)
	}
	if status == http.StatusInternalServerError || status == http.StatusNotFound {
		log.Printf("ERROR: %s - - \"%s %s %s\" - %+v", remoteAddr, method, url, proto, status)
		return
	}
	log.Printf("INFO:  %s - - \"%s %s %s\" - %+v", remoteAddr, method, url, proto, status)
}
