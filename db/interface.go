package db

import "github.com/sufian22/goupon/models"

// ORM is an interface that satisfies object-relational mapping
type ORM interface {
	CreateCupon(coupon *models.Coupon) error
	GetCupon(couponID int) (*models.Coupon, error)
	ListCupons(queryValues *models.ListCouponsQueryValues) ([]*models.Coupon, error)
	UpdateCupon(coupon *models.Coupon) error
}
