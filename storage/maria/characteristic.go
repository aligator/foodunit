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

type Characteristic struct {
	DB *sqlx.DB
}

func NewCharacteristic(db *sqlx.DB) *Characteristic {
	characteristic := Characteristic{
		DB: db,
	}
	return &characteristic
}

func (c *Characteristic) Create() error {
	query := `
CREATE TABLE characteristics (
	id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	name VARCHAR(50) NOT NULL,
	multiple BOOLEAN NOT NULL
)`

	_, err := c.DB.Exec(query)
	return err
}

func (c *Characteristic) Drop() error {
	query := `DROP TABLE IF EXISTS characteristics`
	_, err := c.DB.Exec(query)

	return err
}

func (c *Characteristic) FindByDish(dishID uint64) ([]model.Characteristic, error) {
	query := `
SELECT c.id, c.name, c.multiple FROM dishes_characteristics dc
inner join characteristics c
on c.id = dc.characteristic_id
where dc.dish_id = ?`

	rows, err := c.DB.Queryx(query, dishID)
	if err != nil {
		return nil, err
	}

	characteristics := make([]model.Characteristic, 0)

	for rows.Next() {
		var characteristic model.Characteristic

		if err := rows.StructScan(&characteristic); err != nil {
			return nil, err
		}
		characteristics = append(characteristics, characteristic)
	}

	return characteristics, nil
}
