// Copyright (C) 2015, 2016 Nicolas Lamirault <nicolas.lamirault@gmail.com>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package network

import (
	"encoding/json"
	"log"
	"net"
	"testing"

	dhcp "github.com/krolaw/dhcp4"

	"github.com/nlamirault/shiva/storage"
	"github.com/nlamirault/shiva/utils"
)

var oneOptionSlice = []dhcp.Option{
	dhcp.Option{
		Code:  dhcp.OptionSubnetMask,
		Value: []byte{255, 255, 255, 0},
	},
}

func getStorageBackend() (storage.Storage, error) {
	return storage.New("boltdb", &storage.Config{
		BackendURL: utils.Tempfile(),
	})
}

func createDHCPRequestPacket(chAddr net.HardwareAddr) dhcp.Packet {
	return dhcp.RequestPacket(
		dhcp.Request,
		chAddr,
		nil,
		[]byte{4, 5, 6, 7},
		false,
		oneOptionSlice)
}

func TestDHCPD_WithUnknownMACAddress(t *testing.T) {
	store, err := getStorageBackend()
	if err != nil {
		t.Fatalf("Can't create BoltDB backend.")
	}
	dhcpd := NewDHCPHandler(
		net.IP{192, 168, 0, 1},
		net.IP{192, 168, 0, 2},
		store)
	mac, err := net.ParseMAC("01:23:45:67:89:ab")
	if err != nil {
		t.Fatalf("Can't parse MAC address : %v", err)
	}
	packet := createDHCPRequestPacket(mac)
	resp := dhcpd.ServeDHCP(packet, dhcp.Request, nil)
	if resp != nil {
		t.Fatalf("Invalid DHCP response")
	}
	log.Printf("DHCP response: %v", resp)
}

func TestDHCPD_WithValidMACAddress(t *testing.T) {
	store, err := getStorageBackend()
	if err != nil {
		t.Fatalf("Can't create BoltDB backend.")
	}
	dhcpd := NewDHCPHandler(
		net.IP{192, 168, 0, 1},
		net.IP{192, 168, 0, 2},
		store)
	mac, err := net.ParseMAC("01:23:45:67:89:de")
	if err != nil {
		t.Fatalf("Can't parse MAC address : %v", err)
	}
	data, err := json.Marshal(mac)
	if err != nil {
		t.Fatalf("Can't convert MAC to json : %v", err)
	}
	store.Put([]byte(mac.String()), data)
	packet := createDHCPRequestPacket(mac)
	resp := dhcpd.ServeDHCP(packet, dhcp.Request, nil)
	if resp == nil {
		t.Fatalf("Invalid DHCP response")
	}
	log.Printf("DHCP response: %v", resp)
}
