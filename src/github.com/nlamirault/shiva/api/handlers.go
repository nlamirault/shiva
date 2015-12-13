// Copyright (C) 2015 Nicolas Lamirault <nicolas.lamirault@gmail.com>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package api

import (
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/nlamirault/shiva/api/v1"
	"github.com/nlamirault/shiva/storage"
)

// Authentication represents Basic Authentication entities
type Authentication struct {
	Username string
	Password string
}

// GetWebService return a new gin.Engine
func GetWebService(store storage.Storage, auth *Authentication) *echo.Echo {
	log.Printf("[DEBUG] [shiva] Creating web service using %v",
		store)
	ws := v1.NewWebService(store)
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// Routes
	e.Get("/", ws.Help)
	e.Get("/api/version", ws.DisplayAPIVersion)
	v1 := e.Group("/api/v1")
	if auth != nil {
		log.Printf("[INFO] [shiva] Setup secure mode : %v", auth)
		v1.Use(middleware.BasicAuth(func(user string, pwd string) bool {
			if user == auth.Username && pwd == auth.Password {
				return true
			}
			return false
		}))
	}
	v1.Get("/address/:mac", ws.GetMacAddress)
	v1.Post("/address", ws.AddMacAddress)
	return e
}
