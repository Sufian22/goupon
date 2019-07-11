package models

import (
	"fmt"
	"regexp"
	"strconv"
)

type ListCouponsQueryValues struct {
	FilterValue string
	OrderBy     string
	Order       string
	Quantity    string
}

const (
	// set the regular expression that the filter value should satisfy
	filterValueRegexp = `^[a-z0-9]*$`
)

func (queryValues *ListCouponsQueryValues) Validate() error {
	if _, err := strconv.Atoi(queryValues.Quantity); err != nil {
		return fmt.Errorf("Quantity should be an integer")
	}

	orderBy := queryValues.OrderBy
	if orderBy != "name" && orderBy != "brand" {
		return fmt.Errorf("Coupon list only can be ordered by name or brand")
	}

	order := queryValues.Order
	if order != "ASC" && order != "DESC" {
		return fmt.Errorf("Order must be ASC or DESC")
	}

	matched, err := regexp.MatchString(filterValueRegexp, queryValues.FilterValue)
	if err != nil {
		return err
	}

	if !matched {
		return fmt.Errorf("Filter value should match %s pattern", filterValueRegexp)
	}

	return nil
}
