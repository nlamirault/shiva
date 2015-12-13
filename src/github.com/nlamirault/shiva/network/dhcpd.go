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

package network

import (
	"log"
	"net"
	"time"

	dhcp "github.com/krolaw/dhcp4"

	"github.com/nlamirault/shiva/storage"
)

type lease struct {
	nic    string    // Client's CHAddr
	expiry time.Time // When the lease expires
}

type DHCPHandler struct {
	ip            net.IP        // Server IP to use
	options       dhcp.Options  // Options to send to DHCP Clients
	start         net.IP        // Start of IP range to distribute
	leaseRange    int           // Number of IPs to distribute (starting from start)
	leaseDuration time.Duration // Lease period
	leases        storage.Storage
}

func NewDHCPHandler(serverIP net.IP, startIP net.IP, store storage.Storage) *DHCPHandler {
	return &DHCPHandler{
		ip:            serverIP,
		leaseDuration: 2 * time.Hour,
		start:         startIP,
		leaseRange:    50,
		leases:        store,
		options: dhcp.Options{
			dhcp.OptionSubnetMask:       []byte{255, 255, 240, 0},
			dhcp.OptionRouter:           []byte(serverIP), // Presuming Server is also your router
			dhcp.OptionDomainNameServer: []byte(serverIP), // Presuming Server is also your DNS server
		},
	}
}

// ServeDHCP is called by dhcp.ListenAndServe when the service is started
func (d *DHCPHandler) ServeDHCP(packet dhcp.Packet, msgType dhcp.MessageType, reqOptions dhcp.Options) (response dhcp.Packet) {
	switch msgType {
	case dhcp.Discover:
		// RFC 2131 4.3.1
		mac := packet.CHAddr()
		log.Printf("[INFO] [shiva] DHCP Discover from %s\n", mac.String())

	case dhcp.Request:
		// RFC 2131 4.3.2
		mac := packet.CHAddr()
		log.Printf("[INFO] [shiva] DHCP Request from %s\n", mac.String())

	case dhcp.Decline:
		// RFC 2131 4.3.3
		mac := packet.CHAddr()
		log.Printf("[INFO] [shiva] DHCP Decline from %s\n", mac.String())

	case dhcp.Release:
		// RFC 2131 4.3.4
		mac := packet.CHAddr()
		log.Printf("[INFO] [shiva] DHCP Release from %s\n", mac.String())

	case dhcp.Inform:
		// RFC 2131 4.3.5
		mac := packet.CHAddr()
		log.Printf("[INFO] [shiva] DHCP Inform from %s\n", mac.String())
	}

	return nil
}
