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

// RemoveDNSRecords removes the indicated domain's DNS entries by type and name
func (c *Client) RemoveDNSRecords(domain string, dnsType string, dnsName string) error {
	if len(domain) == 0 {
		return errors.New("Empty domain, this parameter is required")
	}
	if len(dnsType) == 0 {
		return errors.New("Empty dns type, this parameter is required")
	}
	if len(dnsName) == 0 {
		return errors.New("Empty dns name, this parameter is required")
	}

	newRecords := []types.Record{}
	records, err := c.GetDNSRecords(domain, dnsType, "")
	for _, r := range records {
		if r.Type != dnsType || r.Name != dnsName {
			if r.Port == 0 {
				r.Port = 1
			}
			newRecords = append(newRecords, r)
		}
	}

	path := "/v1/domains/" + domain + "/records/" + dnsType
	req, err := c.newRequest("PUT", path, newRecords)
	if err != nil {
		return err
	}

	_, err = c.do(req, nil)
	return err
}
