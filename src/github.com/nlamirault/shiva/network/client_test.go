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
	"testing"

	dhcp "github.com/krolaw/dhcp4"
)

var oneOptionSlice = []dhcp.Option{
	dhcp.Option{
		Code:  dhcp.OptionSubnetMask,
		Value: []byte{255, 255, 255, 0},
	},
}

func Test_DHCPClient(t *testing.T) {
	var tests = []struct {
		description string
		mt          dhcp.MessageType
		chAddr      net.HardwareAddr
		cIAddr      net.IP
		xId         []byte
		broadcast   bool
		options     []dhcp.Option
	}{
		{
			description: "discover request",
			mt:          dhcp.Discover,
			chAddr:      net.HardwareAddr([]byte("4480fa")),
			cIAddr:      net.IP([]byte{192, 168, 1, 1}),
			xId:         []byte{0, 1, 2, 3},
			broadcast:   true,
			options:     nil,
		},
		{
			description: "request request",
			mt:          dhcp.Request,
			chAddr:      net.HardwareAddr([]byte("deadbe")),
			xId:         []byte{4, 5, 6, 7},
			broadcast:   false,
			options:     oneOptionSlice,
		},
	}

	l, _ := net.ListenPacket("udp4", ":68")

	for _, tt := range tests {

		c := dhcp.RequestPacket(tt.mt, tt.chAddr, tt.cIAddr, tt.xId, tt.broadcast, tt.options)
		log.Printf("%v\n", c)

		addr := &net.UDPAddr{IP: net.IPv4bcast, Port: 67}
		if _, e := l.WriteTo(c, addr); e != nil {
			log.Printf("%v\n", e)
		}
	}
	buffer := make([]byte, 1500)

	for {
		n, addr, err := l.ReadFrom(buffer)
		if err != nil {
			log.Println(err)
		}
		log.Printf("[addr]: %v\n", addr.String())
		log.Printf("[n]: %v\n", n)
	}

}
