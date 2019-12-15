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
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var statusGroups []string
var status string

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List owned domains",
	Long: `List owned domains. You can filter them by status (by
default VISIBLE domains are shown). For example:

daddy list

daddy list --statusGroups VISIBLE,INACTIVE

daddy list --status RESERVED

Check https://developer.godaddy.com/doc/endpoint/domains#/v1/list
for more information about valid status and status groups.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		domains, err := GodaddyClient.GetDomains(statusGroups, status)
		if err != nil {
			return err
		}

		var data [][]string = make([][]string, len(domains))
		for i, d := range domains {
			data[i] = []string{
				d.Domain,
				d.CreatedAt.String(),
				d.Expires.String(),
				strconv.FormatBool(d.RenewAuto),
				strconv.FormatBool(d.Locked),
				strconv.FormatBool(d.Privacy),
				d.Status,
			}
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Domain", "Created", "Expires", "Renew", "Locked", "Private", "Status"})
		table.SetBorder(false)
		table.AppendBulk(data)
		table.Render()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().StringSliceVarP(&statusGroups, "statusGroups", "g", []string{"VISIBLE"}, "Domain status groups")
	listCmd.Flags().StringVarP(&status, "status", "t", "", "Domain status groups")
}
