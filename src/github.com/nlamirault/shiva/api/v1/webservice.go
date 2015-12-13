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
	//"encoding/json"
	"errors"
	//"fmt"
	"log"
	//"net/http"

	"github.com/nlamirault/shiva/storage"
)

var (
	// ErrAnalyticsNotEncoded is thrown when an analytics can't be encoded
	ErrAnalyticsNotEncoded = errors.New("Can't encode analytics")

	// ErrAnalyticsNotDecoded is thrown when an analytics can't be decoded
	ErrAnalyticsNotDecoded = errors.New("Can't decode analytics")

	// ErrURLNotEncoded is thrown when an analytics can't be encoded
	ErrURLNotEncoded = errors.New("Can't encode analytics")

	// ErrURLNotDecoded is thrown when an URL can't be decoded
	ErrURLNotDecoded = errors.New("Can't decode url")
)

// WebService represents the Restful API
type WebService struct {
	Store storage.Storage
}

// APIVersion represents version of the REST API
type APIVersion struct {
	Version string `json:"version"`
}

// APIErrorResponse reprensents an error in JSON
type APIErrorResponse struct {
	Error string `json:"error"`
}

// NewWebService creates a new WebService instance
func NewWebService(store storage.Storage) *WebService {
	log.Printf("[DEBUG] [shiva] Creates webservice with backend : %v",
		store)
	return &WebService{Store: store}
}
