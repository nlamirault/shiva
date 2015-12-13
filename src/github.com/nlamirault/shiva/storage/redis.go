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
	"fmt"
	"log"
	"time"

	"github.com/garyburd/redigo/redis"
)

const keyprefix = "shiva"

// Redis represents a storage using the Redis database
type Redis struct {
	//Conn      redis.Conn
	Keyprefix string
	Pool      *redis.Pool
}

// NewRedis instantiates a new Redis database client
func NewRedis(address string) (*Redis, error) {
	log.Printf("[DEBUG] [shiva] New Redis client : %s", address)
	// conn, err := redis.Dial("tcp", fmt.Sprintf(":%s", address))
	// if err != nil {
	// 	return nil, err
	// }
	// log.Printf("[DEBUG] [shiva] Redis connection ready")
	// return &Redis{Conn: conn, Keyprefix: keyprefix}, nil
	pool := &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", fmt.Sprintf(":%s", address))
		},
	}
	log.Printf("[DEBUG] [shiva] Redis client ready")
	return &Redis{Pool: pool, Keyprefix: keyprefix}, nil
}

// Get a value given its key
func (db *Redis) Get(key []byte) ([]byte, error) {
	log.Printf("[DEBUG] [shiva] Get : %v", string(key))
	// exists, err := redis.Bool(c.Do("EXISTS", "foo"))
	// if err != nil {
	// 	return nil, err
	// }
	val, err := db.Pool.Get().Do("HGET", db.Keyprefix, string(key))
	if err != nil {
		return nil, err
	}
	data, err := redis.String(val, nil)
	if err != nil {
		return nil, err
	}
	log.Printf("[INFO] [shiva] Find : %s", data)
	return []byte(data), err
}

// Exists check if a value given its key exists
func (db *Redis) Exists(key []byte) (bool, error) {
	return true, nil
}

// Put a value at the specified key
func (db *Redis) Put(key []byte, value []byte) error {
	log.Printf("[DEBUG] [shiva] Put : %v %v", string(key), string(value))
	_, err := db.Pool.Get().Do("HSET", db.Keyprefix, string(key), value)
	return err
}

// Delete the value at the specified key
func (db *Redis) Delete(key []byte) error {
	log.Printf("[DEBUG] [shiva] Delete : %v", string(key))
	_, err := db.Pool.Get().Do("HDEL", string(key))
	return err
}

// Close the store connection
func (db *Redis) Close() {
	log.Printf("[DEBUG] [shiva] Close")
	//db.Conn.Close()
	if db.Pool != nil {
		db.Pool.Close()
	}
}

// Print backend informations
func (db *Redis) Print() {
	log.Printf("[DEBUG] [shiva] Print")
}
