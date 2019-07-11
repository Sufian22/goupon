package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/sufian22/goupon/api/utils"
	"github.com/sufian22/goupon/config"
	"github.com/sufian22/goupon/models"
)

const (
	configTestFilePath = "../../config/config.json"
)

func initDB() {
	if DB == nil {
		var jsonConfig config.Config
		if err := config.ReadConfigFile(configTestFilePath, &jsonConfig); err != nil {
			log.Fatal(err)
		}

		var err error
		DB, err = config.ConfigureDatabase(jsonConfig.DBConfig)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func TestCreateCouponHandler(t *testing.T) {
	initDB()

	var (
		seed      = rand.NewSource(time.Now().UnixNano())
		random    = rand.New(seed)
		testName  = strconv.Itoa(random.Intn(100000))
		testBrand = "testbrand"
	)
	var testValue float32 = 12.12

	var tdTests = []struct {
		name               *string
		brand              *string
		value              *float32
		expectedStatusCode int
	}{
		{nil, &testBrand, &testValue, http.StatusBadRequest},
		{&testName, nil, &testValue, http.StatusBadRequest},
		{&testName, &testBrand, nil, http.StatusBadRequest},
		{&testName, &testBrand, &testValue, http.StatusOK},
		{&testName, &testBrand, &testValue, http.StatusInternalServerError},
	}

	for _, tt := range tdTests {
		if _, err := createCoupon(tt.name, tt.brand, tt.value, tt.expectedStatusCode); err != nil {
			t.Error(err)
		}
	}
}

func createCoupon(name, brand *string, value *float32, expectedStatusCode int) (*int64, error) {
	createCuponRequest := &createCouponRequest{
		Name:  name,
		Brand: brand,
		Value: value,
	}

	body, err := json.Marshal(createCuponRequest)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", "/api/coupons", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateCouponHandler)

	handler.ServeHTTP(rr, req)

	// check response status code
	if status := rr.Code; status != expectedStatusCode {
		return nil, fmt.Errorf("unexpected status code: expected %v, got %v", expectedStatusCode, status)
	}

	// if OK, check response body
	if expectedStatusCode == http.StatusOK {
		resp := utils.CuponServiceResponse{}
		if err := json.Unmarshal(rr.Body.Bytes(), &resp); err != nil {
			return nil, err
		}

		if resp.Result == nil {
			return nil, fmt.Errorf("unexpected response: result field should not be empty")
		} else {
			resultBytes, err := json.Marshal(resp.Result)
			if err != nil {
				return nil, err
			}

			coupon := models.Coupon{}
			if err := json.Unmarshal(resultBytes, &coupon); err != nil {
				return nil, err
			}

			if coupon.Name != *name || coupon.Brand != *brand || coupon.Value != *value {
				return nil, fmt.Errorf("unexpected coupon values")
			}

			return &coupon.ID, nil
		}
	}

	return nil, nil
}
