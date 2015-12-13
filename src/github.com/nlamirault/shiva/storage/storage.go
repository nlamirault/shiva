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
	//"log"
	"time"
)

const (
	// BOLTDB backend
	BOLTDB string = "boltdb"

	// REDIS backend
	REDIS string = "redis"
)

var (
	// ErrNotSupported is thrown when the backend store is not supported
	ErrNotSupported = errors.New("Backend storage not supported.")

	// ErrNotImplemented is thrown when a method is not implemented by the current backend
	ErrNotImplemented = errors.New("Call not implemented in current backend")

	// ErrEntityNotSaved is thrown when an entity can't be save into the backend
	ErrEntityNotSaved = errors.New("Can't save data")

	// ErrEntityNotStore is thrown when an entity isn't store into the backend
	ErrEntityNotStore = errors.New("Not store data")
)

// Config represents storage configuration
type Config struct {
	BackendURL string
}

// Storage represents the Abraracourcix backend storage
// Each storage should support every call listed
// here.
type Storage interface {

	// Put a value at the specified key
	Put(key []byte, value []byte) error

	// Get a value given its key
	Get(key []byte) ([]byte, error)

	// Delete the value at the specified key
	Delete(key []byte) error

	// Verify if a Key exists in the store
	Exists(key []byte) (bool, error)

	// Close the store connection
	Close()

	// Print backend informations
	Print()
}

// New creates an instance of storage
func New(backend string, config *Config) (Storage, error) {
	switch backend {
	case BOLTDB:
		return NewBoltDB(config.BackendURL)
	case REDIS:
		return NewRedis(config.BackendURL)
	default:
		return nil, fmt.Errorf("%s %s", ErrNotSupported.Error(), "")
	}

}

// URL represents an URL into storage backend
type URL struct {
	// Key is the short URL that expands to the long URL you provided
	Key string `json:"key"`
	// LongURL is the long URL to which it expands.
	LongURL string `json:"longUrl"`
	// CreationDate is the time at which this short URL was created
	CreationDate time.Time `json:"creation_date"`
}
