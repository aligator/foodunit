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

type User struct {
	DB *sqlx.DB
}

func NewUser(db *sqlx.DB) *User {
	user := User{
		DB: db,
	}
	return &user
}

func (u *User) Create() error {
	query := `
CREATE TABLE users (
	id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	mail_addr VARCHAR(254) NOT NULL,
	name VARCHAR(50) NOT NULL,
	is_admin BOOLEAN NOT NULL,
	paypal_mail_addr VARCHAR(254) NOT NULL,
	score INTEGER NOT NULL,
	password_hash CHAR(60) NOT NULL,
	created DATETIME NOT NULL
)`

	_, err := u.DB.Exec(query)
	return err
}

func (u *User) Drop() error {
	query := `DROP TABLE IF EXISTS users`
	_, err := u.DB.Exec(query)

	return err
}

func (u *User) Find(id uint64) (model.User, error) {
	panic("implement me")
}

func (u *User) Exists(id uint64) (bool, error) {
	panic("implement me")
}
