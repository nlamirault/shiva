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

package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	dhcp "github.com/krolaw/dhcp4"

	"github.com/nlamirault/shiva/api"
	"github.com/nlamirault/shiva/logging"
	"github.com/nlamirault/shiva/network"
	"github.com/nlamirault/shiva/storage"
	"github.com/nlamirault/shiva/version"
)

const (
	DefaultBucket string = "shiva"
)

var (
	port     string
	debug    bool
	vrs      bool
	backend  string
	url      string
	username string
	password string
)

func init() {
	// parse flags
	flag.BoolVar(&vrs, "version", false, "print version and exit")
	flag.BoolVar(&vrs, "v", false, "print version and exit (shorthand)")
	flag.BoolVar(&debug, "d", false, "run in debug mode")
	flag.StringVar(&port, "port", "8080", "port to use")
	flag.StringVar(&backend, "backend", "", "Storage backend")
	flag.StringVar(&url, "url", "", "URL backend")
	flag.StringVar(&username, "username", "", "Username authentication")
	flag.StringVar(&password, "password", "", "Password authentication")
	flag.Parse()
}

func main() {
	if debug {
		logging.SetLogging("DEBUG")
	} else {
		logging.SetLogging("INFO")
	}
	if vrs {
		fmt.Printf("Shiva v%s\n", version.Version)
		return
	}

	store, err := storage.New(
		backend,
		&storage.Config{
			BackendURL: url,
		})
	if err != nil {
		log.Printf("[ERROR] [shiva] %s", err.Error())
		return
	}

	var auth *api.Authentication
	if len(username) > 0 && len(password) > 0 {
		auth = &api.Authentication{
			Username: username,
			Password: password,
		}
	}
	e := api.GetWebService(store, auth)
	if debug {
		e.Debug()
	}
	log.Printf("[INFO] [shiva] Launch Shiva on %s using %s backend",
		port, backend)
	go e.Run(fmt.Sprintf(":%s", port))

	dhcpd := network.NewDHCPHandler(
		net.IP{172, 30, 0, 1},
		net.IP{172, 30, 0, 2},
		store)
	err = dhcp.ListenAndServe(dhcpd)
	if err != nil {
		log.Printf("[ERROR] [shiva] %v", err)
		return
	}
}
