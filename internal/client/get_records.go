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
	"errors"

	"github.com/artberri/daddy/internal/types"
)

// GetDNSRecords retrieves the indicated domain's DNS entries info
func (c *Client) GetDNSRecords(domain string, dnsType string, dnsName string) ([]types.Record, error) {
	if len(domain) == 0 {
		return nil, errors.New("Empty domain, this parameter is required")
	}
	if len(dnsName) > 0 && len(dnsType) == 0 {
		return nil, errors.New("You need to filter also by type to filter by name")
	}

	path := "/v1/domains/" + domain + "/records"
	if len(dnsType) > 0 {
		path += "/" + dnsType
	}
	if len(dnsName) > 0 {
		path += "/" + dnsName
	}
	req, err := c.newRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	var records []types.Record
	_, err = c.do(req, &records)
	return records, err
}
