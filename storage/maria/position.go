// Copyright 2019 The FoodUnit Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package maria provides storage implementations for MariaDB.
package maria

import (
	"github.com/dominikbraun/foodunit/model"
	"github.com/jmoiron/sqlx"
)

type Position struct {
	DB *sqlx.DB
}

func NewPosition(db *sqlx.DB) *Position {
	position := Position{
		DB: db,
	}
	return &position
}

func (p *Position) Create() error {
	query := `
CREATE TABLE positions (
	id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	dish_id BIGINT UNSIGNED NOT NULL,
	alternative_dish_id BIGINT UNSIGNED NOT NULL,
	note VARCHAR(200) NOT NULL,
	order_id BIGINT UNSIGNED NOT NULL
)`
	_, err := p.DB.Exec(query)
	return err
}

func (p *Position) Drop() error {
	query := `DROP TABLE IF EXISTS positions`
	_, err := p.DB.Exec(query)

	return err
}

func (p *Position) Store(orderID uint64, position *model.Position) (uint64, error) {
	query := `INSERT INTO positions (dish_id, alternative_dish_id, note, order_id) VALUES (?, ?, ?, ?)`

	result, err := p.DB.Exec(query, position.Dish.ID, position.Alternative.ID, position.Note, orderID)
	if err != nil {
		return uint64(0), err
	}

	// ToDo: When does LastInsertId return an error?
	id, _ := result.LastInsertId()

	return uint64(id), nil
}

func (p *Position) FindByOrder(orderID uint64) ([]model.Position, error) {
	query := `
SELECT p.id, d.id as "dish_id.id", d2.id as "alternative_dish_id.id", note
FROM positions p
INNER JOIN dishes d
ON d.id = p.dish_id
INNER JOIN dishes d2
ON d2.id = p.alternative_dish_id
WHERE p.order_id = ?`

	rows, err := p.DB.Queryx(query, orderID)
	if err != nil {
		return nil, err
	}

	positions := make([]model.Position, 0)

	for rows.Next() {
		var position model.Position

		if err := rows.StructScan(&position); err != nil {
			return nil, err
		}
		positions = append(positions, position)
	}

	return positions, nil
}
