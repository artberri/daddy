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
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestAPIKey is the api key of the GoDaddy server during the tests
var TestAPIKey = "12345"

// TestAPISecret is the api secret of the GoDaddy server during the tests
var TestAPISecret = "54321"

// CreateSimpleTestServer creates a test server with expectedMethod and expectedPah
func CreateSimpleTestServer(t *testing.T, expectedMethod string, expectedPath string) *httptest.Server {
	return CreateTestServer(t, expectedMethod, expectedPath, "")
}

// CreateTestServerWithQueryString creates a test server with expectedMethod, expectedPah and expected query string
func CreateTestServerWithQueryString(t *testing.T, expectedMethod string, expectedPath string, expectedQuery string) *httptest.Server {
	return CreateTestServer(t, expectedMethod, expectedPath, expectedQuery)
}

// CreateTestServer creates a test server with indicated expectations
func CreateTestServer(t *testing.T, expectedMethod string, expectedPath string, expectedQuery string) *httptest.Server {
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
			t.Fatalf("QUery string should be '%s' but '%s' given", expectedQuery, req.URL.RawQuery)
		}
		rw.WriteHeader(200)
		rw.Write([]byte(`[]`))
	}))
}
