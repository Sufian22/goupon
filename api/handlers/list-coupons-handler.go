package handlers

import (
	"log"
	"net/http"

	"github.com/sufian22/goupon/api/utils"
	"github.com/sufian22/goupon/models"
)

var ListCouponsHandler = func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	listCouponsQueryValues, err := getAndValidateQueryValues(r)
	if err != nil {
		utils.WriteResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	coupons, err := DB.ListCupons(listCouponsQueryValues)
	if err != nil {
		log.Print(err)
		utils.WriteResponse(w, http.StatusInternalServerError, "Something failed when listing the coupons", nil)
		return
	}

	utils.WriteResponse(w, http.StatusOK, "", coupons)
}

func getAndValidateQueryValues(r *http.Request) (*models.ListCouponsQueryValues, error) {
	queryValues := r.URL.Query()
	filterValue := queryValues.Get("q")
	orderBy := queryValues.Get("orderBy")
	order := queryValues.Get("order")
	quantity := queryValues.Get("quantity")

	if quantity == "" {
		quantity = "10"
	}

	if orderBy == "" {
		orderBy = "name"
	}

	if order == "" {
		order = "ASC"
	}

	listCouponsQueryValues := &models.ListCouponsQueryValues{
		FilterValue: filterValue,
		OrderBy:     orderBy,
		Order:       order,
		Quantity:    quantity,
	}

	if err := listCouponsQueryValues.Validate(); err != nil {
		return nil, err
	}

	return listCouponsQueryValues, nil
}
