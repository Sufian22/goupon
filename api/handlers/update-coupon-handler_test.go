package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"math/rand"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sufian22/goupon/api/utils"
	"github.com/sufian22/goupon/models"
)

func TestUpdateCouponHandler(t *testing.T) {
	initDB()

	var (
		seed            = rand.NewSource(time.Now().UnixNano())
		random          = rand.New(seed)
		testName        = strconv.Itoa(random.Intn(100000))
		testBrand       = "testbrand"
		invalidCouponID = "as!"

		newName  = testName
		newBrand = "testbrand2"
	)
	var testValue float32 = 12.12
	var newValue float32 = 25

	couponID, err := createCoupon(&testName, &testBrand, &testValue, 200)
	if err != nil {
		t.Fatal(err)
	}

	validCouponID := strconv.Itoa(int(*couponID))
	var tdTests = []struct {
		id                 string
		name               *string
		brand              *string
		value              *float32
		expectedStatusCode int
	}{
		{invalidCouponID, &newName, &newBrand, &newValue, http.StatusBadRequest},
		{validCouponID, nil, &newBrand, &newValue, http.StatusBadRequest},
		{validCouponID, &newName, nil, &newValue, http.StatusBadRequest},
		{validCouponID, &newName, &newBrand, nil, http.StatusBadRequest},
		{validCouponID, &newName, &newBrand, &newValue, http.StatusOK},
	}

	for _, tt := range tdTests {
		updateCuponRequest := &updateCouponRequest{
			Name:  tt.name,
			Brand: tt.brand,
			Value: tt.value,
		}

		body, err := json.Marshal(updateCuponRequest)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest("PUT", "/api/coupons/", bytes.NewBuffer(body))
		if err != nil {
			t.Fatal(err)
		}
		req = mux.SetURLVars(req, map[string]string{"id": tt.id})

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(UpdateCouponHandler)

		handler.ServeHTTP(rr, req)

		// check response status code
		if status := rr.Code; status != tt.expectedStatusCode {
			t.Errorf("unexpected status code: expected %v, got %v", tt.expectedStatusCode, status)
		}

		// if OK, check response body
		if tt.expectedStatusCode == http.StatusOK {
			resp := utils.CuponServiceResponse{}
			if err := json.Unmarshal(rr.Body.Bytes(), &resp); err != nil {
				t.Error(err)
			}

			if resp.Result == nil {
				t.Errorf("unexpected response: result field should not be empty")
			} else {
				resultBytes, err := json.Marshal(resp.Result)
				if err != nil {
					t.Error(err)
				}

				coupon := models.Coupon{}
				if err := json.Unmarshal(resultBytes, &coupon); err != nil {
					t.Error(err)
				}

				if coupon.Name != newName || coupon.Brand != newBrand || coupon.Value != newValue {
					t.Error("unexpected new coupon values")
				}
			}
		}
	}
}
