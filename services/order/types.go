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

// Package order provides services and types for Order-related data.
package order

type Order struct {
	ID        uint64     `json:"id"`
	User      User       `json:"user_id"`
	Positions []Position `json:"positions"`
	IsPaid    bool       `json:"is_paid"`
	Total     uint       `json:"total"`
}

type User struct {
	Name string `json:"name"`
}

type Position struct {
	ID          uint64 `json:"id"`
	Dish        Dish   `json:"dish_id"`
	Alternative Dish   `json:"alternative_dish_id"`
	Note        string `json:"note"`
}

type Dish struct {
	ID              uint64           `json:"id"`
	Name            string           `json:"name"`
	Description     string           `json:"description"`
	Price           uint             `json:"price"`
	IsUncertain     bool             `json:"is_uncertain"`
	IsHealthy       bool             `json:"is_healthy"`
	IsVegetarian    bool             `json:"is_vegetarian"`
	Characteristics []Characteristic `json:"characteristics"`
}

type Characteristic struct {
	Name     string   `json:"name"`
	Multiple bool     `json:"multiple"`
	Values   []string `json:"values"`
}