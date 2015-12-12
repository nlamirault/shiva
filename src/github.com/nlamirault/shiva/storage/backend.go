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

package storage

import (
	"errors"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/docker/libkv"
	"github.com/docker/libkv/store"
	"github.com/docker/libkv/store/boltdb"
	"github.com/docker/libkv/store/consul"
	"github.com/docker/libkv/store/etcd"
	"github.com/docker/libkv/store/zookeeper"
)

const (
	// CONSUL backend
	CONSUL string = "consul"

	// BOLTDB backend
	BOLTDB string = "boltdb"

	// ETCD backend
	ETCD string = "etcd"

	// ZOOKEEPER backend
	ZOOKEEPER string = "zk"

	DefaultBucket string = "shiva"
)

var (
	// ErrNotSupported is thrown when the backend store is not supported
	ErrNotSupported = errors.New("Backend storage not supported.")
)

type MACEntry struct {
	MAC      net.HardwareAddr
	IP       net.IP
	Duration time.Duration
	Attr     map[string]string
}

type IPEntry struct {
	MAC net.HardwareAddr
}

type Storage struct {
	store.Store
}

// type Backend interface {
// 	InitDHCP()
// 	GetIP(net.IP) (IPEntry, error)
// 	HasIP(net.IP) bool
// 	GetMAC(mac net.HardwareAddr, cascade bool) (entry *MACEntry, found bool, err error)
// 	CreateLease(lease *MACEntry) error
// 	WriteLease(lease *MACEntry) error
// }

func New(backend string, client string) (*Storage, error) {
	st, err := newStore(backend, client)
	if err != nil {
		return nil, err
	}
	log.Printf("[INFO] Storage backend : %s", backend)
	return &Storage{Store: st}, nil
}

func newStore(backend string, client string) (store.Store, error) {
	conf := &store.Config{
		Bucket:            DefaultBucket,
		ConnectionTimeout: 10 * time.Second,
	}
	switch backend {
	case BOLTDB:
		boltdb.Register()
		return libkv.NewStore(
			store.BOLTDB,
			[]string{client},
			conf,
		)
	case CONSUL:
		consul.Register()
		return libkv.NewStore(
			store.CONSUL,
			[]string{client},
			conf,
		)
	case ETCD:
		etcd.Register()
		return libkv.NewStore(
			store.ETCD,
			[]string{client},
			conf,
		)
	case ZOOKEEPER:
		zookeeper.Register()
		return libkv.NewStore(
			store.ZK,
			[]string{client},
			conf,
		)
	default:
		return nil, fmt.Errorf("%s %s", ErrNotSupported.Error(), "")
	}
}
