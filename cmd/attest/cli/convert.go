/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cli

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/tkn-intoto-formatter/pkg/common"
	"github.com/tkn-intoto-formatter/pkg/convertor"
)

// // convertCmd represents the serve command
// var convertCmd = &cobra.Command{
// 	Use:   "convert -i [file-name] -f [<]output filepath] -o [format]",
// 	Short: "converts yaml tekton spec to specified format",
// 	Long: `A longer description that spans multiple lines and likely contains examples
// and usage of using your command. For example:

// Cobra is a CLI library for Go that empowers applications.
// This application is a tool to generate the needed files
// to quickly create a Cobra application.`,
// 	Run: func(cmd *cobra.Command, args []string) {

// 	},
// }

// func init() {
// 	RootCmd.AddCommand(convertCmd)
// }

//ConvertCmd :
func ConvertCmd() *cobra.Command {
	var cOpts common.ConvertOpt

	var cmd = &cobra.Command{
		Use:   "convert -i [file-name] -f [<]output filepath] -o [format]",
		Short: "converts yaml tekton spec to specified format",
		Long: `A longer description that spans multiple lines and likely contains examples
				and usage of using your command. For example:
			`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := execConvertOp(cOpts); err != nil {
				return errors.Wrapf(err, "error converting the input file: %s", cOpts.InputFilepath)
			}
			return nil
		},
	}

	cmd.PersistentFlags().StringVarP(&cOpts.InputFilepath, "infile", "i", "", "input filepath for for tekton resource yaml")
	cmd.PersistentFlags().StringVarP(&cOpts.OutputFilepath, "outfile", "f", "", "output filepath to store result")
	cmd.PersistentFlags().StringVarP(&cOpts.InputFilepath, "format", "o", "", "output format to use (default: intoto-ite6)")
	cmd.PersistentFlags().BoolVar(&cOpts.StubVal, "stubval", false, "stub variable values (default: false)")

	return cmd
}

func execConvertOp(cOpts common.ConvertOpt) error {
	convertor.ConvertToIntoto(cOpts)
	return nil
}
