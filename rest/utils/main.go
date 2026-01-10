package utils

import (
	"log"
	"net/http"
)

type routerFunc func(w http.ResponseWriter, r *http.Request) error

func UseHTTPRouterFunc(routerFunc routerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := routerFunc(w, r); err != nil {

			// send error response here
			log.Println("ERROR OCCURED IN ROUTER FUNCTION")
		}
	}
}
