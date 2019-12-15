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
	"errors"

	"github.com/artberri/daddy/internal/types"
)

// UpdateDNSRecord updates the indicated domain's DNS entriy
func (c *Client) UpdateDNSRecord(domain string, dnsType string, dnsName string, dnsValue string, dnsTTL int, dnsPrority int) error {
	if len(domain) == 0 {
		return errors.New("Empty domain, this parameter is required")
	}
	if len(dnsType) == 0 {
		return errors.New("Empty dns type, this parameter is required")
	}
	if len(dnsName) == 0 {
		return errors.New("Empty dns name, this parameter is required")
	}
	if len(dnsValue) == 0 {
		return errors.New("Empty dns value, this parameter is required")
	}

	path := "/v1/domains/" + domain + "/records/" + dnsType + "/" + dnsName

	req, err := c.newRequest("PUT", path, []types.Record{{
		Type:     dnsType,
		Name:     dnsName,
		Data:     dnsValue,
		TTL:      dnsTTL,
		Priority: dnsPrority,
		Port:     1,
		Service:  "",
		Protocol: "",
		Weight:   0,
	}})
	if err != nil {
		return err
	}

	_, err = c.do(req, nil)
	return err
}
