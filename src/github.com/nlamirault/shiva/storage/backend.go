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

	// "github.com/docker/libkv"
	"github.com/docker/libkv/store"
)

const (
	// CONSUL backend
	CONSUL string = "consul"

	// BOLTDB backend
	BOLTDB string = "boltdb"

	// ETCD backend
	ETCD string = "etcd"

	// ZOOKEEPER backend
	ZOOKEEPER string = "zookeeper"
)

var (
	// ErrNotSupported is thrown when the backend store is not supported
	ErrNotSupported = errors.New("Backend storage not supported.")
)

type Storage interface {
	Cat(path string) (*store.KVPair, error)
	Ls(path string) ([]*store.KVPair, error)
	Rm(path string, recursive bool) error
	Dump(path string) ([]byte, error)
	Restore(archive []byte) error
}

// New creates an instance of storage
func New(backend string) (store.Store, error) {
	// switch backend {
	// case BOLTDB:
	// 	return libkv.NewStore(
	// 		store.BOLTDB,
	// 		[]string{client},
	// 		&store.Config{
	// 			ConnectionTimeout: 10 * time.Second,
	// 		},
	// 	)
	// case CONSUL:
	// 	return libkv.NewStore(
	// 		store.CONSUL,
	// 		[]string{client},
	// 		&store.Config{
	// 			ConnectionTimeout: 10 * time.Second,
	// 		},
	// 	)
	// case ETCD:
	// 	return libkv.NewStore(
	// 		store.ETCD,
	// 		[]string{client},
	// 		&store.Config{
	// 			ConnectionTimeout: 10 * time.Second,
	// 		},
	// 	)
	// case ZOOKEEPER:
	// 	return libkv.NewStore(
	// 		store.ZOOKEEPER,
	// 		[]string{client},
	// 		&store.Config{
	// 			ConnectionTimeout: 10 * time.Second,
	// 		},
	// 	)
	// default:
	// 	return nil, fmt.Errorf("%s %s", ErrNotSupported.Error(), "")
	// }
	return nil, nil

}
