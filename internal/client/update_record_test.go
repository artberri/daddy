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
	"github.com/artberri/daddy/internal/types"
)

func TestUpdateDNSRecordThrowsErrorIfNoDomainSet(t *testing.T) {
	testDomain := ""
	testType := "CNAME"
	testName := "www"
	testValue := "@"
	testPriority := 0
	testTTL := 3600
	expectedPath := "/v1/domains/" + testDomain + "/records"
	expectedMethod := "GET"
	server := testutil.CreateSimpleTestServer(t, expectedMethod, expectedPath)

	defer server.Close()

	c := createUpdateDNSRecordClient(server)

	err := c.UpdateDNSRecord(testDomain, testType, testName, testValue, testTTL, testPriority)

	if err == nil {
		t.Fatal("Error was expected")
	}
}

func TestUpdateDNSRecordThrowsErrorIfNoTypeSet(t *testing.T) {
	testDomain := "example.com"
	testType := ""
	testName := "www"
	testValue := "@"
	testPriority := 0
	testTTL := 3600
	expectedPath := "/v1/domains/" + testDomain + "/records"
	expectedMethod := "GET"
	server := testutil.CreateSimpleTestServer(t, expectedMethod, expectedPath)

	defer server.Close()

	c := createUpdateDNSRecordClient(server)

	err := c.UpdateDNSRecord(testDomain, testType, testName, testValue, testTTL, testPriority)

	if err == nil {
		t.Fatal("Error was expected")
	}
}

func TestUpdateDNSRecordThrowsErrorIfNoNameSet(t *testing.T) {
	testDomain := "example.com"
	testType := "CNAME"
	testName := ""
	testValue := "@"
	testPriority := 0
	testTTL := 3600
	expectedPath := "/v1/domains/" + testDomain + "/records"
	expectedMethod := "GET"
	server := testutil.CreateSimpleTestServer(t, expectedMethod, expectedPath)

	defer server.Close()

	c := createUpdateDNSRecordClient(server)

	err := c.UpdateDNSRecord(testDomain, testType, testName, testValue, testTTL, testPriority)

	if err == nil {
		t.Fatal("Error was expected")
	}
}

func TestUpdateDNSRecordThrowsErrorIfNoValueSet(t *testing.T) {
	testDomain := "example.com"
	testType := "CNAME"
	testName := "www"
	testValue := ""
	testPriority := 0
	testTTL := 3600
	expectedPath := "/v1/domains/" + testDomain + "/records"
	expectedMethod := "GET"
	server := testutil.CreateSimpleTestServer(t, expectedMethod, expectedPath)

	defer server.Close()

	c := createUpdateDNSRecordClient(server)

	err := c.UpdateDNSRecord(testDomain, testType, testName, testValue, testTTL, testPriority)

	if err == nil {
		t.Fatal("Error was expected")
	}
}

func TestUpdateDNSRecordCallsProperURLWithProperBody(t *testing.T) {
	testDomain := "example.com"
	testType := "CNAME"
	testName := "www"
	testValue := "@"
	testPriority := 0
	testTTL := 3600
	expectedPath := "/v1/domains/" + testDomain + "/records/" + testType + "/" + testName
	expectedMethod := "PUT"
	server := testutil.CreateTestServerWithBody(t, expectedMethod, expectedPath, []types.Record{{
		Type:     testType,
		Name:     testName,
		Data:     testValue,
		TTL:      testTTL,
		Priority: testPriority,
		Port:     1,
		Service:  "",
		Protocol: "",
		Weight:   0,
	}})

	defer server.Close()

	c := createUpdateDNSRecordClient(server)

	err := c.UpdateDNSRecord(testDomain, testType, testName, testValue, testTTL, testPriority)

	if err != nil {
		t.Fatalf("Not error expected, '%s' error found", err)
	}
}

func createUpdateDNSRecordClient(server *httptest.Server) *Client {
	c, _ := CreateClient(server.URL, testutil.TestAPIKey, testutil.TestAPISecret, server.Client())
	return c
}
