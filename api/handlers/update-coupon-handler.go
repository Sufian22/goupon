package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sufian22/goupon/api/utils"
	"github.com/sufian22/goupon/models"
)

type updateCouponRequest struct {
	Name  *string  `json:"name"`
	Brand *string  `json:"brand"`
	Value *float32 `json:"value"`
}

// validating the request body (e.g. nullness, regexp..)
func (ccr *updateCouponRequest) validate() error {
	if ccr.Name == nil {
		return fmt.Errorf("Coupon name could not be empty")
	}

	if ccr.Brand == nil {
		return fmt.Errorf("Coupon brand could not be empty")
	}

	if ccr.Value == nil {
		return fmt.Errorf("Coupon value could not be empty")
	}

	return nil
}

var UpdateCouponHandler = func(w http.ResponseWriter, r *http.Request) {
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

	updateCouponRequest := updateCouponRequest{}
	if err := json.NewDecoder(r.Body).Decode(&updateCouponRequest); err != nil {
		log.Print(err)
		utils.WriteResponse(w, http.StatusInternalServerError, "Something failed when decoding the request body", nil)
		return
	}

	if err := updateCouponRequest.validate(); err != nil {
		log.Print(err)
		utils.WriteResponse(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	coupon := models.Coupon{
		ID:    int64(couponID),
		Name:  *updateCouponRequest.Name,
		Brand: *updateCouponRequest.Brand,
		Value: *updateCouponRequest.Value,
	}

	if err := DB.UpdateCupon(&coupon); err != nil {
		log.Print(err)
		utils.WriteResponse(w, http.StatusInternalServerError, "Something failed when updating the coupon", nil)
		return
	}

	utils.WriteResponse(w, http.StatusOK, "Cupon created successfully", updateCouponRequest)
}
