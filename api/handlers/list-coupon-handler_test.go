package handlers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/sufian22/goupon/api/utils"
	"github.com/sufian22/goupon/models"
)

func TestListCouponsHandler(t *testing.T) {
	initDB()

	var (
		seed      = rand.NewSource(time.Now().UnixNano())
		random    = rand.New(seed)
		testName  = strconv.Itoa(random.Intn(100000))
		testBrand = "testbrand"

		validOrderBy  = "name"
		validOrder    = "ASC"
		validQuantity = "1"
	)
	var testValue float32 = 12.12

	if _, err := createCoupon(&testName, &testBrand, &testValue, 200); err != nil {
		t.Fatal(err)
	}

	var tdTests = []struct {
		filterValue        string
		orderBy            string
		order              string
		quantity           string
		expectedStatusCode int
	}{
		{"", validOrderBy, validOrder, validQuantity, http.StatusOK},
		{"!<invalid", validOrderBy, validOrder, validQuantity, http.StatusBadRequest},
		{"", "invalid", validOrder, validQuantity, http.StatusBadRequest},
		{"", validOrderBy, "invalid", validQuantity, http.StatusBadRequest},
		{"", validOrderBy, validOrder, "invalid", http.StatusBadRequest},
	}

	for _, tt := range tdTests {
		url := fmt.Sprintf("/api/coupons?q=%s&orderBy=%s&order=%s&quantity=%v",
			tt.filterValue, tt.orderBy, tt.order, tt.quantity)

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(ListCouponsHandler)

		handler.ServeHTTP(rr, req)

		// check response status code
		if status := rr.Code; status != tt.expectedStatusCode {
			t.Errorf("unexpected status code: expected %v, got %v", http.StatusOK, status)
		}

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

				coupons := []models.Coupon{}
				if err := json.Unmarshal(resultBytes, &coupons); err != nil {
					t.Error(err)
				}

				quantity, _ := strconv.Atoi(tt.quantity)
				if len(coupons) != quantity {
					t.Errorf("unexpected number of coupons: expected %v got %v", quantity, len(coupons))
				}
			}
		}
	}
}
