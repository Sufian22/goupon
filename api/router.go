package api

import (
	"github.com/gorilla/mux"
	"github.com/sufian22/goupon/handlers"
)

func newRouter() *mux.Router {
	router := mux.NewRouter()
	router.Use(loggerMiddleware)

	router.HandleFunc("/api/coupons", handlers.ListCouponsHandler).Methods("GET")
	router.HandleFunc("/api/coupons", handlers.CreateCouponHandler).Methods("POST")
	router.HandleFunc("/api/coupons/{id}", handlers.GetCouponHandler).Methods("GET")
	router.HandleFunc("/api/coupons/{id}", handlers.UpdateCouponHandler).Methods("PUT")

	return router
}
