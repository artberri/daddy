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

// Package client contains the http client for the GoDaddy API
package client

import (
	"net/http/httptest"
	"testing"

	"github.com/artberri/daddy/internal/testutil"
)

func TestGetDomainsCallsProperURLWithoutParameters(t *testing.T) {
	testStatusGroups := []string{}
	testStatus := ""
	expectedMethod := "GET"
	expectedPath := "/v1/domains"
	server := testutil.CreateSimpleTestServer(t, expectedMethod, expectedPath)

	defer server.Close()

	c := createDomainsClient(server)

	records, err := c.GetDomains(testStatusGroups, testStatus)

	if err != nil {
		t.Fatalf("Not error expected, '%s' error found", err)
	}

	if len(records) > 0 {
		t.Fatalf("Result not expected")
	}
}

func TestGetDomainsCallsProperURLWithStatusParameter(t *testing.T) {
	testStatusGroups := []string{}
	testStatus := "ACTIVE"
	expectedMethod := "GET"
	expectedPath := "/v1/domains"
	expectedQueryString := "status=ACTIVE"
	server := testutil.CreateTestServerWithQueryString(t, expectedMethod, expectedPath, expectedQueryString)

	defer server.Close()

	c := createDomainsClient(server)

	records, err := c.GetDomains(testStatusGroups, testStatus)

	if err != nil {
		t.Fatalf("Not error expected, '%s' error found", err)
	}

	if len(records) > 0 {
		t.Fatalf("Result not expected")
	}
}

func TestGetDomainsCallsProperURLWithStatusGroupsParameter(t *testing.T) {
	testStatusGroups := []string{"VISIBLE", "INACTIVE"}
	testStatus := ""
	expectedMethod := "GET"
	expectedPath := "/v1/domains"
	expectedQueryString := "statusGroups=VISIBLE%2CINACTIVE"
	server := testutil.CreateTestServerWithQueryString(t, expectedMethod, expectedPath, expectedQueryString)

	defer server.Close()

	c := createDomainsClient(server)

	records, err := c.GetDomains(testStatusGroups, testStatus)

	if err != nil {
		t.Fatalf("Not error expected, '%s' error found", err)
	}

	if len(records) > 0 {
		t.Fatalf("Result not expected")
	}
}

func TestGetDomainsCallsProperURLWithStatusGroupsAndStatusParameter(t *testing.T) {
	testStatusGroups := []string{"VISIBLE", "INACTIVE"}
	testStatus := "ACTIVE"
	expectedMethod := "GET"
	expectedPath := "/v1/domains"
	expectedQueryString := "status=ACTIVE&statusGroups=VISIBLE%2CINACTIVE"
	server := testutil.CreateTestServerWithQueryString(t, expectedMethod, expectedPath, expectedQueryString)

	defer server.Close()

	c := createDomainsClient(server)

	records, err := c.GetDomains(testStatusGroups, testStatus)

	if err != nil {
		t.Fatalf("Not error expected, '%s' error found", err)
	}

	if len(records) > 0 {
		t.Fatalf("Result not expected")
	}
}

func createDomainsClient(server *httptest.Server) *Client {
	c, _ := CreateClient(server.URL, testutil.TestAPIKey, testutil.TestAPISecret, server.Client())
	return c
}
