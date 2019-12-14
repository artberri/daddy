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
	"fmt"
	"os"

	"github.com/artberri/daddy/internal/client"
	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

// GodaddyClient is the http client for the API
var GodaddyClient client.Client

var godaddyAPIURL string
var godaddyAPIKey string
var godaddyAPISecret string
var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "daddy",
	Short: "Command line to manage domain records in GoDaddy",
	Long: `Command line to manage domain records in GoDaddy. You need to
setup your API key and secret to obtain/set data. For example:

daddy list --key=1234567689 --secret=1234566

If you do not want to pass the parameters on each command you can create
the file $HOME/.daddy.yaml and put your configuration there. For example:

---
key: 1234567689
secret: 1234567689
`,
	SilenceUsage: false,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		url := viper.GetString("url")
		key := viper.GetString("key")
		if len(key) == 0 {
			return errors.New("Empty API Key, this parameter is required")
		}

		secret := viper.GetString("secret")
		if len(secret) == 0 {
			return errors.New("Empty API Secret, this parameter is required")
		}

		c, err := client.CreateClient(url, key, secret)
		if err != nil {
			return err
		}

		GodaddyClient = *c
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.daddy.yaml)")
	rootCmd.PersistentFlags().StringP("url", "u", "https://api.godaddy.com", "URL base of the GoDaddy API")
	rootCmd.PersistentFlags().StringP("key", "k", "", "API Key for the GoDaddy API (Required)")
	rootCmd.PersistentFlags().StringP("secret", "s", "", "API Secret for the GoDaddy API (Required)")
	viper.BindPFlag("url", rootCmd.PersistentFlags().Lookup("url"))
	viper.BindPFlag("key", rootCmd.PersistentFlags().Lookup("key"))
	viper.BindPFlag("secret", rootCmd.PersistentFlags().Lookup("secret"))
	viper.SetDefault("url", "https://api.godaddy.com")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".daddy" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".daddy")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
