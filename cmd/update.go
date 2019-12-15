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

// Package cmd defines the commands of the daddy cli
package cmd

import (
	"github.com/spf13/cobra"
)

var updateDomain string
var updateDNSType string
var updateDNSName string
var updateDNSValue string
var updateDNSTTL int
var updateDNSPriority int

var updateCmd = &cobra.Command{
	Use:     "update",
	Aliases: []string{"u"},
	Short:   "Update domain record",
	Long: `Update DNS record of a single domain. BE CAREFUL, this command will always replace the values of
ALL the records of the given type and name. For example:

daddy update --domain mydomain.com -t CNAME -n www -v other.com

daddy update --domain mydomain.com -t A -n @ --ttl 3600 --priority 30 -v 8.8.8.8

Check https://developer.godaddy.com/doc/endpoint/domains#/v1/recordReplaceTypeName
for more information.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return GodaddyClient.UpdateDNSRecord(updateDomain, updateDNSType, updateDNSName, updateDNSValue, updateDNSTTL, updateDNSPriority)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringVarP(&updateDomain, "domain", "d", "", "Domain you want to add record to (Required)")
	updateCmd.Flags().StringVarP(&updateDNSType, "type", "t", "", "DNS Record type (Required)")
	updateCmd.Flags().StringVarP(&updateDNSName, "name", "n", "", "DNS Record name (Required)")
	updateCmd.Flags().StringVarP(&updateDNSValue, "value", "v", "", "DNS Record data (Required)")
	updateCmd.Flags().IntVarP(&updateDNSTTL, "ttl", "l", 3600, "DNS Record TTL")
	updateCmd.Flags().IntVarP(&updateDNSPriority, "priority", "p", 0, "DNS Record TTL")
}
