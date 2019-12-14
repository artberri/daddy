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
package cmd

import (
	"errors"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var domain string
var dnsType string
var dnsName string

var showCmd = &cobra.Command{
	Use:     "show",
	Aliases: []string{"sh"},
	Short:   "Show domain records",
	Long: `Show DNS records of a single domain. You can filter them by type
and/or by name. For example:

daddy show

daddy show --domain mydomain.com

daddy show --domain mydomain.com -t A

daddy show --domain mydomain.com -t A -n www

Check https://developer.godaddy.com/doc/endpoint/domains#/operations/v1/recordGet
for more information.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		records, err := GodaddyClient.GetDNSRecords(domain, dnsType, dnsName)
		if err != nil {
			return err
		}

		if len(records) == 0 {
			return errors.New("No results found")
		}

		var data [][]string = make([][]string, len(records))
		for i, r := range records {
			data[i] = []string{
				r.Type,
				r.Name,
				r.Data,
				strconv.Itoa(r.TTL),
				strconv.Itoa(r.Priority),
			}
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Type", "Name", "Data", "TTL", "Priority"})
		table.SetBorder(false)
		table.AppendBulk(data)
		table.Render()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
	showCmd.Flags().StringVarP(&domain, "domain", "d", "", "Domain you want to query records from (Required)")
	showCmd.Flags().StringVarP(&dnsType, "type", "t", "", "Filter by DNS Record type")
	showCmd.Flags().StringVarP(&dnsName, "name", "n", "", "Filter by DNS Record name")
}
