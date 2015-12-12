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

package v1

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/docker/libkv/store"
	"github.com/labstack/echo"
)

type MACAddress struct {
	MAC net.HardwareAddr `json:"mac"`
}

// ListMacAddress display all MAC addresses
func (ws *WebService) ListMacAddress(c *echo.Context) error {
	return nil
}

// AddMacAddress add into storage a MAC address
func (ws *WebService) AddMacAddress(c *echo.Context) error {
	var mac MACAddress
	c.Bind(&mac)
	log.Printf("[INFO] [shiva] MAC address to store: %v", mac)
	exists, err := ws.Store.Exists(mac.MAC.String())
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			&APIErrorResponse{Error: err.Error()})
	}
	if exists {
		return c.JSON(http.StatusBadRequest,
			&APIErrorResponse{
				Error: fmt.Sprintf("MAC already exists %s", mac),
			})
	}
	ws.Store.Put(mac.MAC.String(), mac.MAC, &store.WriteOptions{IsDir: true})
	log.Printf("[INFO] [shiva] MAC stored : %v", mac)
	return c.JSON(http.StatusOK, mac)
}

// GetMacAddress search into storage a MAC address
func (ws *WebService) GetMacAddress(c *echo.Context) error {
	mac := c.Param("mac")
	log.Printf("[INFO] [shiva] Retrieve MAC using: %s\n", mac)
	exists, err := ws.Store.Exists(mac)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			&APIErrorResponse{Error: err.Error()})
	}
	if !exists {
		return c.JSON(http.StatusNotFound,
			&APIErrorResponse{
				Error: fmt.Sprintf("Unknown MAC %s", mac),
			})
	}
	kv, err := ws.Store.Get(mac)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			&APIErrorResponse{Error: err.Error()})
	}
	log.Printf("[INFO] [shiva] Find MAC : %v", kv)
	return c.JSON(http.StatusOK, kv)

}
