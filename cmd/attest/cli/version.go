//
// Copyright 2020 IBM Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	goVersion "go.hein.dev/go-version"
)

var (
	shortened = false
	version   = "dev"
	commit    = "none"
	date      = "unknown"
	output    = "json"
)

// func init() {
// 	RootCmd.AddCommand(versionCmd)
// }

//VersionCmd :
func VersionCmd() *cobra.Command {
	// versionCmd represents the serve command
	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "tkn-intoto-formatter version",
		Long:  `Prints the tkn-intoto-formatter tool version`,
		Run: func(cmd *cobra.Command, args []string) {
			resp := goVersion.FuncWithOutput(shortened, version, commit, date, output)
			fmt.Print(resp)
		},
	}
	return versionCmd
}
