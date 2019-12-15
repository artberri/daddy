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

func TestRemoveDNSRecordsThrowsErrorIfNoDomainSet(t *testing.T) {
	testDomain := ""
	testType := "CNAME"
	testName := "www"
	expectedPath := "/v1/domains/" + testDomain + "/records/" + testType
	expectedMethod := "PUT"
	server := testutil.CreateSimpleTestServer(t, expectedMethod, expectedPath)

	defer server.Close()

	c := createRemoveDNSRecordsClient(server)

	err := c.RemoveDNSRecords(testDomain, testType, testName)

	if err == nil {
		t.Fatal("Error was expected")
	}
}

func TestRemoveDNSRecordsThrowsErrorIfNoTypeSet(t *testing.T) {
	testDomain := "example.com"
	testType := ""
	testName := "www"
	expectedPath := "/v1/domains/" + testDomain + "/records/" + testType
	expectedMethod := "PUT"
	server := testutil.CreateSimpleTestServer(t, expectedMethod, expectedPath)

	defer server.Close()

	c := createRemoveDNSRecordsClient(server)

	err := c.RemoveDNSRecords(testDomain, testType, testName)

	if err == nil {
		t.Fatal("Error was expected")
	}
}

func TestRemoveDNSRecordsThrowsErrorIfNoNameSet(t *testing.T) {
	testDomain := "example.com"
	testType := "CNAME"
	testName := ""
	expectedPath := "/v1/domains/" + testDomain + "/records/" + testType
	expectedMethod := "PUT"
	server := testutil.CreateSimpleTestServer(t, expectedMethod, expectedPath)

	defer server.Close()

	c := createRemoveDNSRecordsClient(server)

	err := c.RemoveDNSRecords(testDomain, testType, testName)

	if err == nil {
		t.Fatal("Error was expected")
	}
}

func TestRemoveDNSRecordsCallsProperURLWithProperBody(t *testing.T) {
	testDomain := "example.com"
	testType := "CNAME"
	testName := "www"
	expectedPath := "/v1/domains/" + testDomain + "/records/" + testType
	expectedMethod := "PUT"

	record1 := types.Record{
		Type:     testType,
		Name:     testName,
		Data:     "data1",
		TTL:      3600,
		Priority: 0,
		Port:     0,
		Service:  "",
		Protocol: "",
		Weight:   0,
	}
	record2 := types.Record{
		Type:     testType,
		Name:     testName,
		Data:     "data2",
		TTL:      3600,
		Priority: 0,
		Port:     0,
		Service:  "",
		Protocol: "",
		Weight:   0,
	}
	record3 := types.Record{
		Type:     testType,
		Name:     "otherName",
		Data:     "data3",
		TTL:      3600,
		Priority: 0,
		Port:     0,
		Service:  "",
		Protocol: "",
		Weight:   0,
	}

	expectedRequestRecord := record3
	expectedRequestRecord.Port = 1

	server := testutil.CreateTestServerWithPreviousGetResponse(
		t,
		[]types.Record{record1, record2, record3},
		expectedMethod,
		expectedPath,
		[]types.Record{expectedRequestRecord},
	)

	defer server.Close()

	c := createRemoveDNSRecordsClient(server)

	err := c.RemoveDNSRecords(testDomain, testType, testName)

	if err != nil {
		t.Fatalf("Not error expected, '%s' error found", err)
	}
}

func createRemoveDNSRecordsClient(server *httptest.Server) *Client {
	c, _ := CreateClient(server.URL, testutil.TestAPIKey, testutil.TestAPISecret, server.Client())
	return c
}
