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
	"errors"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var removeDomain string
var removeType string
var removeName string
var removeForce bool

var removeCmd = &cobra.Command{
	Use:     "remove",
	Aliases: []string{"rm", "r"},
	Short:   "Remove DNS records",
	Long: `Remove DNS records of a domain by type and by name. For example:

daddy remove --domain mydomain.com -t CNAME -n www

Check https://developer.godaddy.com/doc/endpoint/domains#/v1/recordReplaceType
for more information.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if !removeForce {
			prompt := promptui.Prompt{
				Label:     "Are you sure you want to remove the indicated records?",
				Default:   "n",
				IsConfirm: true,
			}

			_, err := prompt.Run()
			if err != nil {
				return errors.New("Action not confirmed")
			}
		}

		return GodaddyClient.RemoveDNSRecords(removeDomain, removeType, removeName)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
	removeCmd.Flags().StringVarP(&removeDomain, "domain", "d", "", "Domain you want to query records from (Required)")
	removeCmd.Flags().StringVarP(&removeType, "type", "t", "", "DNS Record type (Required)")
	removeCmd.Flags().StringVarP(&removeName, "name", "n", "", "DNS Record name (Required)")
	removeCmd.Flags().BoolVarP(&removeForce, "force", "f", false, "Do not ask for confirmation")
}
