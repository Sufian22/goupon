package handlers

import (
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

func TestGetCouponHandler(t *testing.T) {
	initDB()

	var (
		seed            = rand.NewSource(time.Now().UnixNano())
		random          = rand.New(seed)
		testName        = strconv.Itoa(random.Intn(100000))
		testBrand       = "testbrand"
		invalidCouponID = "99999"
	)
	var testValue float32 = 12.12

	couponID, err := createCoupon(&testName, &testBrand, &testValue, 200)
	if err != nil {
		t.Fatal(err)
	}

	var tdTests = []struct {
		id                 string
		expectedStatusCode int
	}{
		{"", http.StatusInternalServerError},
		{invalidCouponID, http.StatusInternalServerError},
		{strconv.Itoa(int(*couponID)), http.StatusOK},
	}

	for _, tt := range tdTests {
		req, err := http.NewRequest("GET", "/api/coupons/", nil)
		if err != nil {
			t.Fatal(err)
		}
		req = mux.SetURLVars(req, map[string]string{"id": tt.id})

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(GetCouponHandler)

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

				if coupon.Name != testName || coupon.Brand != testBrand || coupon.Value != testValue {
					t.Error("unexpected coupon values")
				}
			}
		}
	}
}
