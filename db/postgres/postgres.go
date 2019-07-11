package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"github.com/sufian22/goupon/models"
)

const errDuplicateEntry = "23505"

// DB is the complete database.
// impl is the implementation of the DB interface for postgresql
type DB struct {
	impl *sql.DB
}

// NewDB returns a new DB instance from its configuration
func NewDB(url string) (*DB, error) {
	return NewDBFromConfig(url)
}

// NewDBFromConfig configures the proper sql driver and tries the db connection
func NewDBFromConfig(url string) (*DB, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &DB{db}, err
}

func (db *DB) CreateCupon(coupon *models.Coupon) error {
	coupon.CreatedAt = time.Now().UTC()
	coupon.Expiry = time.Now().Add(time.Hour * 24 * 30).UTC()

	tx, err := db.impl.Begin()
	if err != nil {
		return err
	}

	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}()

	err = tx.QueryRow(`INSERT INTO coupons (
		name,
		brand,
		value,
		createdAt,
		expiry
	) VALUES (
		$1,
		$2,
		$3,
		$4,
		$5
	) RETURNING id`, coupon.Name, coupon.Brand, coupon.Value, coupon.CreatedAt, coupon.Expiry).Scan(&coupon.ID)
	if err != nil {
		dbError, ok := err.(*pq.Error)
		if !ok {
			return err
		}
		if dbError.Code == errDuplicateEntry {
			errorMessage := fmt.Sprintf("Duplicate entry error when inserting %s coupon", coupon.Name)
			return errors.New(errorMessage)
		}
		return err
	}

	return nil
}

func (db *DB) GetCupon(couponID int) (*models.Coupon, error) {
	tx, err := db.impl.Begin()
	if err != nil {
		return nil, err
	}

	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}()

	coupon := &models.Coupon{}
	err = tx.QueryRow(`SELECT
		id,
		name,
		brand,
		value,
		createdAt,
		expiry
		FROM coupons WHERE id = $1
	`, couponID).Scan(&coupon.ID, &coupon.Name, &coupon.Brand, &coupon.Value, &coupon.CreatedAt, &coupon.Expiry)
	if err != nil {
		return nil, err
	}

	return coupon, nil
}

func (db *DB) ListCupons(queryValues *models.ListCouponsQueryValues) ([]*models.Coupon, error) {
	coupons := []*models.Coupon{}

	tx, err := db.impl.Begin()
	if err != nil {
		return coupons, err
	}

	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}()

	filter := fmt.Sprint("%", queryValues.FilterValue, "%")
	query := fmt.Sprintf(
		`SELECT
			id,
			name,
			brand,
			value,
			createdAt,
			expiry
		FROM coupons
		WHERE 
			name like $1 OR
			brand like $1
		ORDER BY %s %s
		LIMIT %s`, queryValues.OrderBy, queryValues.Order, queryValues.Quantity,
	)

	rows, err := tx.Query(query, filter)
	if err != nil {
		return coupons, err
	}
	defer rows.Close()

	for rows.Next() {
		coupon := &models.Coupon{}
		if err := rows.Scan(
			&coupon.ID,
			&coupon.Name,
			&coupon.Brand,
			&coupon.Value,
			&coupon.CreatedAt,
			&coupon.Expiry,
		); err != nil {
			return coupons, err
		}

		coupons = append(coupons, coupon)
	}

	return coupons, nil
}

func (db *DB) UpdateCupon(coupon *models.Coupon) error {
	tx, err := db.impl.Begin()
	if err != nil {
		return err
	}

	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}()

	r, err := tx.Exec(`UPDATE coupons
	SET
		name = $1,
		brand = $2,
		value = $3
	WHERE
		id = $4;`, coupon.Name, coupon.Brand, coupon.Value, coupon.ID)
	if err != nil {
		return err
	}

	if _, err := r.RowsAffected(); err != nil {
		return err
	}

	return nil
}
