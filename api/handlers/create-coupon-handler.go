package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/sufian22/goupon/api/utils"
	"github.com/sufian22/goupon/db"
	"github.com/sufian22/goupon/models"
)

// DB is the database access.
var DB db.ORM

type createCouponRequest struct {
	Name  *string  `json:"name"`
	Brand *string  `json:"brand"`
	Value *float32 `json:"value"`
}

// validating the request body (e.g. nullness, regexp..)
func (ccr *createCouponRequest) validate() error {
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

var CreateCouponHandler = func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	createCuponRequest := createCouponRequest{}
	if err := json.NewDecoder(r.Body).Decode(&createCuponRequest); err != nil {
		log.Print(err)
		utils.WriteResponse(w, http.StatusInternalServerError, "Something failed when decoding the request body", nil)
		return
	}

	if err := createCuponRequest.validate(); err != nil {
		log.Print(err)
		utils.WriteResponse(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	coupon := models.Coupon{
		Name:  *createCuponRequest.Name,
		Brand: *createCuponRequest.Brand,
		Value: *createCuponRequest.Value,
	}

	if err := DB.CreateCupon(&coupon); err != nil {
		log.Print(err)
		utils.WriteResponse(w, http.StatusInternalServerError, "Something failed when storing the coupon", nil)
		return
	}

	utils.WriteResponse(w, http.StatusOK, "Coupon updated successfully", coupon)
}
