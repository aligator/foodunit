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

// Package restaurant provides services and types for Restaurant-related data.
package restaurant

type Info struct {
	Name       string `json:"name"`
	Street     string `json:"street"`
	PostalCode string `json:"postal_code"`
	City       string `json:"city"`
	Phone      string `json:"phone"`
	Open       string `json:"open"`
	Website    string `json:"website"`
}

type Menu struct {
	Categories []MenuCategory `json:"categories"`
}

type MenuCategory struct {
	Name   string     `json:"name"`
	Dishes []MenuDish `json:"dishes"`
}

type MenuDish struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	Price        uint   `json:"price"`
	IsUncertain  bool   `json:"is_uncertain"`
	IsHealthy    bool   `json:"is_healthy"`
	IsVegetarian bool   `json:"is_vegetarian"`
}
