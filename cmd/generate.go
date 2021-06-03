// Copyright © 2021 DataHen Canada Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"log"

	"github.com/DataHenHQ/useragent"
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate {desktop|mobile|googlebot2}",
	Short: "Generates user agent strings",
	Long: `Generates user agent strings.
Usage: useragent generate <user agent type>
Example: useragent generate browser

`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		uaType := args[0]

		uaConfig, err := cmd.Flags().GetString("ua-config")
		if err != nil {
			log.Fatal(err)
		}
		if uaConfig != "" {
			useragent.LoadUAConfig(uaConfig)
		}

		nums, err := cmd.Flags().GetInt("numbers")
		if err != nil {
			log.Fatalf("Gotten error: %v\n", err.Error())
			return
		}

		// loop through how many strings to generat
		for i := 0; i < nums; i++ {
			var ua string
			switch uaType {
			case "desktop":
				if s, err := useragent.Desktop(); err == nil {
					ua = s
				}
			case "mobile":
				if s, err := useragent.Mobile(); err == nil {
					ua = s
				}
			case "googlebot2":
				if s, err := useragent.GoogleBot2(); err == nil {
					ua = s
				}
			}
			if ua == "" {
				continue
			}

			fmt.Println(ua)
		}

	},
}

var schemas string

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().IntP("numbers", "n", 1, "How many user agent strings to generate")

	generateCmd.Flags().String("ua-config", "", "Path to the user agent config path, if not specified, the default built-in user agent combinations will be used")
}
