package handlers

import (
	"net/http"
)

func ListCouponsHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}
