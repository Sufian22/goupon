package api

import (
	"net/http"
	"os"

	gorillaHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/sufian22/goupon/api/handlers"
)

func newRouter() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/api/coupons", handlers.ListCouponsHandler).Methods("GET")
	router.HandleFunc("/api/coupons", handlers.CreateCouponHandler).Methods("POST")
	router.HandleFunc("/api/coupons/{id}", handlers.GetCouponHandler).Methods("GET")
	router.HandleFunc("/api/coupons/{id}", handlers.UpdateCouponHandler).Methods("PUT")

	return gorillaHandlers.LoggingHandler(os.Stdout, router)
}
