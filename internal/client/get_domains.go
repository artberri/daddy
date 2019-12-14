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
	"strings"

	"github.com/artberri/daddy/internal/types"
)

// GetDomains retrieves the domains from the GoDaddy API
func (c *Client) GetDomains(statusGroups []string, status string) ([]types.Domain, error) {
	req, err := c.newRequest("GET", "/v1/domains", nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("statusGroups", strings.Join(statusGroups, ","))
	q.Add("status", status)
	req.URL.RawQuery = q.Encode()

	var domains []types.Domain
	_, err = c.do(req, &domains)
	return domains, err
}
