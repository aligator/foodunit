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
	"database/sql"
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
	if err != nil {
		return err
	}
	query = `
CREATE TABLE confirmation_tokens (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT UNSIGNED NOT NULL,
    token VARCHAR(255) NOT NULL,
    is_confirmed BOOLEAN NOT NULL
)`
	_, err = u.DB.Exec(query)

	return err
}

func (u *User) Drop() error {
	query := `DROP TABLE IF EXISTS users, confirmation_tokens`
	_, err := u.DB.Exec(query)

	return err
}

func (u *User) Store(user *model.User) (uint64, error) {
	query := `
INSERT INTO users (
    mail_addr, name, is_admin, paypal_mail_addr, score, password_hash, created
) VALUES (?, ?, ?, ?, ?, ?, ?)`

	created := user.Created.Format("2006-01-02 15:04:05")

	result, err := u.DB.Exec(query, user.MailAddr, user.Name, user.IsAdmin, user.PaypalMailAddr, user.Score, user.PasswordHash, created)
	if err != nil {
		return uint64(0), err
	}

	// ToDo: When does LastInsertId return an error?
	id, _ := result.LastInsertId()

	return uint64(id), nil
}

func (u *User) Find(id uint64) (model.User, error) {
	query := `SELECT * FROM users WHERE id = ?`

	var user model.User
	err := u.DB.QueryRowx(query, id).StructScan(&user)

	return user, err
}

func (u *User) FindByMailAddr(mailAddr string) (model.User, error) {
	query := `SELECT * FROM users where mail_addr = ?`

	var user model.User
	err := u.DB.QueryRowx(query, mailAddr).StructScan(&user)

	return user, err
}

func (u *User) MailExists(mailAddr string) (bool, error) {
	query := `SELECT * FROM users where mail_addr = ?`

	var user model.User
	err := u.DB.QueryRowx(query, mailAddr).StructScan(&user)

	if err == sql.ErrNoRows {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil
}

func (u *User) StoreConfirmationToken(userID uint64, token string) error {
	query := `INSERT INTO confirmation_tokens (user_id, token, is_confirmed) VALUES (?, ?, ?)`
	_, err := u.DB.Exec(query, userID, token, false)

	return err
}

func (u *User) ConfirmUser(token string) error {
	query := `UPDATE confirmation_tokens SET is_confirmed = 1 WHERE token = ?`
	_, err := u.DB.Exec(query, token)

	return err
}

func (u *User) SetPaypalMailAddr(id uint64, paypalMailAddr string) error {
	query := `UPDATE users SET paypal_mail_addr = ? WHERE id = ?`
	_, err := u.DB.Exec(query, paypalMailAddr, id)

	return err
}
