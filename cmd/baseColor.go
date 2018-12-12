// Copyright © 2018 Adrian Simmons <adrian@perlucida.com>
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
	"strings"

	"github.com/lucasb-eyer/go-colorful"
	"github.com/spf13/cobra"
)

// baseColorCmd represents the baseColor command
var baseColorCmd = &cobra.Command{
	Use:   "baseColor [color]",
	Short: "Specify the base color.",
	Long: `A base color from which the entire palette is calculated.
	
  Should be specified in FORMAT`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here

		c := colorful.HappyColor()

		fmt.Println("Random happy color:" + c)
		fmt.Println("baseColor called with:" + strings.Join(args, " "))
	},
}

func init() {
	RootCmd.AddCommand(baseColorCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// baseColorCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// baseColorCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	baseColorCmd.Flags().StringVarP(&baseColor, "basecolor", "c", "", "base color e.g. #ffcc00 (required)")
	baseColorCmd.MarkFlagRequired("basecolor")
}
