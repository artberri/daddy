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

// CreateTestServer creates a test server with indicated expectations
func CreateTestServer(t *testing.T, expectedMethod string, expectedPath string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.Method != expectedMethod {
			t.Fatalf("Method should be '%s', '%s' given", expectedMethod, req.Method)
		}
		if req.URL.Path != expectedPath {
			t.Fatalf("Path should be '%s' but '%s' given", expectedPath, req.URL.Path)
		}
		rw.WriteHeader(200)
		rw.Write([]byte(`[]`))
	}))
}
