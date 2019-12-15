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

func TestGetDNSRecordsThrowsErrorIfNoDomainSet(t *testing.T) {
	testDomain := ""
	expectedPath := "/v1/domainss/" + testDomain + "/records"
	expectedMethod := "GET"
	server := testutil.CreateSimpleTestServer(t, expectedMethod, expectedPath)

	defer server.Close()

	c := createRecordsClient(server)

	_, err := c.GetDNSRecords(testDomain, "", "")

	if err == nil {
		t.Fatal("Error was expected")
	}
}

func TestGetDNSRecordsThrowsErrorIfNameSetWithoutType(t *testing.T) {
	testDomain := "example.com"
	testDNSType := ""
	testDNSName := "other.org"
	expectedPath := "/v1/domains/" + testDomain + "/records/" + testDNSType + "/" + testDNSName
	expectedMethod := "GET"
	server := testutil.CreateSimpleTestServer(t, expectedMethod, expectedPath)

	defer server.Close()

	c := createRecordsClient(server)

	_, err := c.GetDNSRecords(testDomain, testDNSType, testDNSName)

	if err == nil {
		t.Fatal("Error was expected")
	}
}

func TestGetDNSRecordsCallsProperURLWithOnlyDomainParameter(t *testing.T) {
	testDomain := "example.com"
	expectedPath := "/v1/domains/" + testDomain + "/records"
	expectedMethod := "GET"
	server := testutil.CreateSimpleTestServer(t, expectedMethod, expectedPath)

	defer server.Close()

	c := createRecordsClient(server)

	records, err := c.GetDNSRecords(testDomain, "", "")

	if err != nil {
		t.Fatalf("Not error expected, '%s' error found", err)
	}

	if len(records) > 0 {
		t.Fatalf("Result not expected")
	}
}

func TestGetDNSRecordsCallsProperURLWithDomainAndTypeParameter(t *testing.T) {
	testDomain := "example.com"
	testDNSType := "A"
	expectedPath := "/v1/domains/" + testDomain + "/records/" + testDNSType
	expectedMethod := "GET"
	server := testutil.CreateSimpleTestServer(t, expectedMethod, expectedPath)

	defer server.Close()

	c := createRecordsClient(server)

	records, err := c.GetDNSRecords(testDomain, testDNSType, "")

	if err != nil {
		t.Fatalf("Not error expected, '%s' error found", err)
	}

	if len(records) > 0 {
		t.Fatalf("Result not expected")
	}
}

func TestGetDNSRecordsCallsProperURLWithDomainAndTypeAndNameParameter(t *testing.T) {
	testDomain := "example.com"
	testDNSType := "A"
	testDNSName := "other.org"
	expectedPath := "/v1/domains/" + testDomain + "/records/" + testDNSType + "/" + testDNSName
	expectedMethod := "GET"
	server := testutil.CreateSimpleTestServer(t, expectedMethod, expectedPath)

	defer server.Close()

	c := createRecordsClient(server)

	records, err := c.GetDNSRecords(testDomain, testDNSType, testDNSName)

	if err != nil {
		t.Fatalf("Not error expected, '%s' error found", err)
	}

	if len(records) > 0 {
		t.Fatalf("Result not expected")
	}
}

func createRecordsClient(server *httptest.Server) *Client {
	c, _ := CreateClient(server.URL, testutil.TestAPIKey, testutil.TestAPISecret, server.Client())
	return c
}
