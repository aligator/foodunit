// Package mariadb provides repository implementations as SQL gateways.
package mariadb

import (
	"github.com/dominikbraun/foodunit/dl"
	"time"
)

type (
	// OfferRepository holds methods of dl.OfferRepository.
	OfferRepository struct{}
	// OrderRepository holds methods of dl.OrderRepository.
	OrderRepository struct{}
	// PositionRepository holds methods of dl.PositionRepository.
	PositionRepository struct{}
)

// Migrate implements dl.OfferRepository.Migrate.
func (o OfferRepository) Migrate() error {
	schema := buildCreate("offers", fieldMap{
		"id":          "BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY",
		"user_id":     "BIGINT UNSIGNED",
		"supplier_id": "BIGINT UNSIGNED",
		"valid_from":  "DATETIME",
		"valid_to":    "DATETIME",
		"is_placed":   "BIT(1)",
		"pickup_info": "VARCHAR(100)",
	})

	db, err := GetDB()
	if err != nil {
		return err
	}

	_ = db.MustExec(schema)
	return nil
}

// Create implements dl.OfferRepository.Create.
func (o OfferRepository) Create(offer *dl.Offer) (uint64, error) {
	query := buildInsert("offers", fieldMap{
		"user_id":     ":user_id",
		"supplier_id": ":supplier_id",
		"valid_from":  ":valid_from",
		"valid_to":    ":valid_to",
		"is_placed":   ":is_placed",
		"pickup_info": ":pickup_info",
	})

	db, err := GetDB()
	if err != nil {
		return 0, err
	}

	r, err := db.NamedExec(query, offer)
	if err != nil {
		return 0, err
	}

	return lastInsertID(r)
}

// Find implements dl.OfferRepository.Find.
func (o OfferRepository) Find(id uint64) (*dl.Offer, error) {
	var offer dl.Offer
	query := buildSelect("offers", []string{"*"}, nil, conditionMap{
		"id": "= ?",
	})

	db, err := GetDB()
	if err != nil {
		return nil, err
	}

	_ = db.QueryRowx(query, id).StructScan(&offer)
	return &offer, nil
}

// FindAllActive implements dl.OfferRepository.FindAllActive.
func (o OfferRepository) FindAllActive() ([]*dl.Offer, error) {
	var offers []*dl.Offer
	query := buildSelect("offers", []string{"*"}, nil, conditionMap{
		"valid_from":   "<= ?",
		"AND valid_to": "> ?",
		"AND NOT":      "is_placed",
	})

	db, err := GetDB()
	if err != nil {
		return nil, err
	}

	now := time.Now().Format(time.RFC3339)

	rows, err := db.Queryx(query, now, now)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var offer dl.Offer
		// ToDo: The error value of Scan has to be handled.
		_ = rows.Scan(&offer)
		offers = append(offers, &offer)
	}

	return offers, nil
}

// Update implements dl.OfferRepository.Update.
func (o OfferRepository) Update(offer *dl.Offer) error {
	query := buildUpdate("offers", fieldMap{
		"user_id":     ":user_id",
		"supplier_id": ":supplier_id",
		"valid_from":  ":valid_from",
		"valid_to":    ":valid_to",
		"is_replaced": ":is_replaced",
		"pickup_info": ":pickup_info",
	}, conditionMap{"id": "= :id"})

	db, err := GetDB()
	if err != nil {
		return err
	}

	if _, err = db.NamedExec(query, &offer); err != nil {
		return err
	}
	return nil
}

// Delete implements dl.OfferRepository.Delete.
func (o OfferRepository) Delete(offer *dl.Offer) error {
	query := buildDelete("offers", conditionMap{
		"id": "= ?",
	})

	db, err := GetDB()
	if err != nil {
		return err
	}

	if _, err = db.Exec(query, offer.ID); err != nil {
		return err
	}
	return nil
}

// Migrate implements dl.OrderRepository.Migrate.
func (o OrderRepository) Migrate() error {
	schema := buildCreate("orders", fieldMap{
		"id":       "BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY",
		"user_id":  "BIGINT UNSIGNED",
		"offer_id": "BIGINT UNSIGNED",
	})

	db, err := GetDB()
	if err != nil {
		return err
	}

	_ = db.MustExec(schema)
	return nil
}

// Create implements dl.OrderRepository.Create.
func (o OrderRepository) Create(order *dl.Order) (uint64, error) {
	query := buildInsert("orders", fieldMap{
		"user_id":  ":user_id",
		"offer_id": ":offer_id",
	})

	db, err := GetDB()
	if err != nil {
		return 0, err
	}

	r, err := db.NamedExec(query, order)
	if err != nil {
		return 0, err
	}

	return lastInsertID(r)
}

// Find implements dl.OrderRepository.Find.
func (o OrderRepository) Find(id uint64) (*dl.Order, error) {
	var order dl.Order
	query := buildSelect("orders", []string{"*"}, nil, conditionMap{
		"id": "= ?",
	})

	db, err := GetDB()
	if err != nil {
		return nil, err
	}

	_ = db.QueryRowx(query, id).StructScan(&order)
	return &order, nil
}

// FindByOfferID implements dl.OrderRepository.FindByOfferID.
func (o OrderRepository) FindByOfferID(offerID uint64) ([]*dl.Order, error) {
	var orders []*dl.Order
	query := buildSelect("orders", []string{"*"}, nil, conditionMap{
		"offer_id": "= ?",
	})

	db, err := GetDB()
	if err != nil {
		return nil, err
	}

	now := time.Now().Format(time.RFC3339)

	rows, err := db.Queryx(query, now, now)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var order dl.Order
		// ToDo: The error value of Scan has to be handled.
		_ = rows.Scan(&order)
		orders = append(orders, &order)
	}

	return orders, nil
}

// Update implements dl.OrderRepository.Update.
func (o OrderRepository) Update(order *dl.Order) error {
	panic("implement me")
}

// Delete implements dl.OrderRepository.Delete.
func (o OrderRepository) Delete(order *dl.Order) error {
	panic("implement me")
}

// Migrate implements dl.PositionRepository.Migrate.
func (p PositionRepository) Migrate() error {
	schema := buildCreate("positions", fieldMap{
		"id":       "BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY",
		"order_id": "BIGINT UNSIGNED",
		"dish_id":  "BIGINT UNSIGNED",
		"note":     "VARCHAR(100)",
	})

	db, err := GetDB()
	if err != nil {
		return err
	}

	_ = db.MustExec(schema)
	return nil
}

// Create implements dl.PositionRepository.Create.
func (p PositionRepository) Create(position *dl.Position) (uint64, error) {
	panic("implement me")
}

// Find implements dl.PositionRepository.Find.
func (p PositionRepository) Find(id uint64) (*dl.Position, error) {
	panic("implement me")
}

// FindByOrderID implements dl.PositionRepository.FindByOrderID.
func (p PositionRepository) FindByOrderID(orderID uint64) ([]*dl.Position, error) {
	panic("implement me")
}

// Update implements dl.PositionRepository.Update.
func (p PositionRepository) Update(position *dl.Position) error {
	panic("implement me")
}

// Delete implements dl.PositionRepository.Delete.
func (p PositionRepository) Delete(position *dl.Position) error {
	panic("implement me")
}