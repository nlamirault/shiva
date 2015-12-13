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
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"

	//"github.com/docker/libkv/store"
	"github.com/labstack/echo"
)

type MACAddress struct {
	Address string `json:"mac"`
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
	exists, err := ws.Store.Exists([]byte(mac.Address))
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			&APIErrorResponse{Error: err.Error()})
	}
	if exists {
		return c.JSON(http.StatusBadRequest,
			&APIErrorResponse{
				Error: fmt.Sprintf("MAC already exists %s",
					mac.Address),
			})
	}
	_, err = net.ParseMAC(mac.Address)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			&APIErrorResponse{Error: err.Error()})
	}
	data, err := json.Marshal(mac)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			&APIErrorResponse{Error: err.Error()})
	}
	ws.Store.Put([]byte(mac.Address), data)
	// log.Printf("[INFO] [shiva] MAC stored : %v", mac)
	return c.JSON(http.StatusOK, mac)
}

// GetMacAddress search into storage a MAC address
func (ws *WebService) GetMacAddress(c *echo.Context) error {
	mac := c.Param("mac")
	log.Printf("[INFO] [shiva] Retrieve MAC using: %s\n", mac)
	data := []byte(mac)
	exists, err := ws.Store.Exists(data)
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
	entry, err := ws.Store.Get(data)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			&APIErrorResponse{Error: err.Error()})
	}
	log.Printf("[INFO] [shiva] Make MAC : %s\n", string(entry))
	var addr *MACAddress
	err = json.Unmarshal(entry, &addr)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			&APIErrorResponse{Error: err.Error()})
	}
	log.Printf("[INFO] [shiva] Find MAC : %v", addr)
	return c.JSON(http.StatusOK, addr)

}
