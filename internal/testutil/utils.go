/*
Copyright © 2019 Alberto Varela <alberto@berriart.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/

// Package testutil contains utilities to run tests
package testutil

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/artberri/daddy/internal/types"
)

// TestAPIKey is the api key of the GoDaddy server during the tests
var TestAPIKey = "12345"

// TestAPISecret is the api secret of the GoDaddy server during the tests
var TestAPISecret = "54321"

// CreateSimpleTestServer creates a test server with expectedMethod and expectedPah
func CreateSimpleTestServer(t *testing.T, expectedMethod string, expectedPath string) *httptest.Server {
	return CreateTestServer(t, expectedMethod, expectedPath, "", nil)
}

// CreateTestServerWithQueryString creates a test server with expectedMethod, expectedPah and expected query string
func CreateTestServerWithQueryString(t *testing.T, expectedMethod string, expectedPath string, expectedQuery string) *httptest.Server {
	return CreateTestServer(t, expectedMethod, expectedPath, expectedQuery, nil)
}

// CreateTestServerWithBody creates a test server with expectedMethod, expectedPah and expected body
func CreateTestServerWithBody(t *testing.T, expectedMethod string, expectedPath string, expectedBody interface{}) *httptest.Server {
	return CreateTestServer(t, expectedMethod, expectedPath, "", expectedBody)
}

// CreateTestServer creates a test server with indicated expectations
func CreateTestServer(t *testing.T, expectedMethod string, expectedPath string, expectedQuery string, expectedBody interface{}) *httptest.Server {
	expectedAuthorizationHeader := "sso-key " + TestAPIKey + ":" + TestAPISecret

	return httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.Header.Get("Authorization") != expectedAuthorizationHeader {
			t.Fatalf("Authorization header should be '%s' but '%s' given", expectedAuthorizationHeader, req.Header.Get("Authorization"))
		}
		if req.Method != expectedMethod {
			t.Fatalf("Method should be '%s', '%s' given", expectedMethod, req.Method)
		}
		if req.URL.Path != expectedPath {
			t.Fatalf("Path should be '%s' but '%s' given", expectedPath, req.URL.Path)
		}
		if req.URL.RawQuery != expectedQuery {
			t.Fatalf("Query string should be '%s' but '%s' given", expectedQuery, req.URL.RawQuery)
		}

		var buf io.ReadWriter
		if expectedBody != nil {
			buf = new(bytes.Buffer)
			err := json.NewEncoder(buf).Encode(expectedBody)
			if err != nil {
				t.Fatal("Can't encode test body")
			}
			var records []types.Record
			err = json.NewDecoder(req.Body).Decode(&records)
			if err != nil {
				t.Fatal("Can't decode request body")
			}
			if !reflect.DeepEqual(expectedBody, records) {
				t.Fatal("Not expected body")
			}
		}

		rw.WriteHeader(200)
		rw.Write([]byte(`[]`))
	}))
}

// CreateTestServerWithPreviousGetResponse creates a test server ready for two response
func CreateTestServerWithPreviousGetResponse(t *testing.T, firstResponse interface{}, expectedMethod string, expectedPath string, expectedBody interface{}) *httptest.Server {
	expectedAuthorizationHeader := "sso-key " + TestAPIKey + ":" + TestAPISecret

	return httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.Header.Get("Authorization") != expectedAuthorizationHeader {
			t.Fatalf("Authorization header should be '%s' but '%s' given", expectedAuthorizationHeader, req.Header.Get("Authorization"))
		}
		if req.Method != expectedMethod && req.Method != "GET" {
			t.Fatalf("Method should be '%s', '%s' given", expectedMethod, req.Method)
		}
		if req.URL.Path != expectedPath {
			t.Fatalf("Path should be '%s' but '%s' given", expectedPath, req.URL.Path)
		}

		var buf io.ReadWriter
		if req.Method != "GET" {
			buf = new(bytes.Buffer)
			err := json.NewEncoder(buf).Encode(expectedBody)
			if err != nil {
				t.Fatal("Can't encode test body")
			}
			var records []types.Record
			err = json.NewDecoder(req.Body).Decode(&records)
			if err != nil {
				t.Fatal("Can't decode request body")
			}
			if !reflect.DeepEqual(expectedBody, records) {
				t.Fatal("Not expected body")
			}
		}

		fr := new(bytes.Buffer)
		err := json.NewEncoder(fr).Encode(firstResponse)
		if err != nil {
			t.Fatal("Can't encode first response")
		}

		rw.WriteHeader(200)
		rw.Write(fr.Bytes())
	}))
}
