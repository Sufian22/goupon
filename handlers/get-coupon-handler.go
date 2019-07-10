package handlers

import (
	"net/http"
)

func GetCouponHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}
