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
	//"fmt"
	"testing"
)

func Test_GetNotSupportedBackend(t *testing.T) {
	backend, err := InitStorage(
		"google", &Config{Data: "/tmp/google.db"})
	if backend != nil || err == nil {
		t.Fatalf("Error retrieve an invalid backend")
	}
}

func Test_GetMemDBBackend(t *testing.T) {
	backend, err := InitStorage("memdb", &Config{Data: "/tmp"})
	if err != nil || backend == nil {
		t.Fatalf("Error retrieve MemDB backend %v.", err)
	}
}

func Test_GetBoltDBBackend(t *testing.T) {
	backend, err := InitStorage(
		"boltdb", &Config{Data: "/tmp/foo.db"})
	if err != nil || backend == nil {
		t.Fatalf("Error retrieve BoltDB backend %v.", err)
	}
}

func Test_GetLevelDBBackend(t *testing.T) {
	backend, err := InitStorage(
		"leveldb", &Config{Data: "/tmp"})
	if err != nil || backend == nil {
		t.Fatalf("Error retrieve LevelDB backend %v.", err)
	}
}

func Test_GetRedisBackend(t *testing.T) {
	backend, err := InitStorage(
		"redis", &Config{BackendURL: "127.0.0.1:6379"})
	if err != nil || backend == nil {
		t.Fatalf("Error retrieve Redis backend %v.", err)
	}
}

func Test_GetMongoDBBackend(t *testing.T) {
	backend, err := InitStorage(
		"mongodb", &Config{BackendURL: "127.0.0.1:27017"})
	if err != nil || backend == nil {
		t.Fatalf("Error retrieve MongoDB backend %v.", err)
	}
}
