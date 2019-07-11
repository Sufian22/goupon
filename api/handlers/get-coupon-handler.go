package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sufian22/goupon/api/utils"
)

var GetCouponHandler = func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	idVar, ok := vars["id"]
	if !ok {
		utils.WriteResponse(w, http.StatusBadRequest, "Coupon identifier could not be empty", nil)
		return
	}

	couponID, err := strconv.Atoi(idVar)
	if err != nil {
		log.Print(err)
		utils.WriteResponse(w, http.StatusInternalServerError, "Coupon identifier should be an integer", nil)
		return
	}

	coupon, err := DB.GetCupon(couponID)
	if err != nil {
		log.Print(err)
		utils.WriteResponse(w, http.StatusInternalServerError, "Something went wrong when retrieving the coupon", nil)
		return
	}

	utils.WriteResponse(w, http.StatusOK, "", coupon)
}
