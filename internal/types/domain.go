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

// Package types contains the types to represent the GoDaddy API objects
package types

import "time"

// Domain represents a domain.
type Domain struct {
	Domain    string
	CreatedAt time.Time
	Expires   time.Time
	RenewAuto bool
	Locked    bool
	Privacy   bool
	Status    string
}
