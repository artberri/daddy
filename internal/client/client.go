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
package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const defaultUserAgent = "daddy-cli"

// Client is the HTTPClient for the GoDaddy API
type Client struct {
	baseURL   *url.URL
	apiKey    string
	apiSecret string
	userAgent string

	httpClient *http.Client
}

// CreateClient creates an HTTPClient for the GoDaddy API
func CreateClient(baseURL string, apiKey string, apiSecret string) (*Client, error) {
	url, err := parseBaseURL(baseURL)
	if err != nil {
		return nil, err
	}

	return &Client{
		baseURL:    url,
		apiKey:     apiKey,
		apiSecret:  apiSecret,
		userAgent:  defaultUserAgent,
		httpClient: &http.Client{},
	}, nil
}

func (c *Client) newRequest(method, path string, body interface{}) (*http.Request, error) {
	rel := &url.URL{
		Path: path,
	}
	u := c.baseURL.ResolveReference(rel)
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.userAgent)
	req.Header.Set("Authorization", "sso-key "+c.apiKey+":"+c.apiSecret)
	return req, nil
}

func (c *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if !(resp.StatusCode >= 200 && resp.StatusCode <= 299) {
		return nil, fmt.Errorf("Invalid response status `%s`", resp.Status)
	}

	err = json.NewDecoder(resp.Body).Decode(v)
	return resp, err
}

func parseBaseURL(host string) (*url.URL, error) {
	protoAddrParts := strings.SplitN(host, "://", 2)
	if len(protoAddrParts) == 1 {
		return nil, fmt.Errorf("Unable to parse godaddy host `%s`", host)
	}

	proto, addr := protoAddrParts[0], protoAddrParts[1]

	return &url.URL{
		Scheme: proto,
		Host:   addr,
	}, nil
}
