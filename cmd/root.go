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
	"os"

	colorful "github.com/lucasb-eyer/go-colorful"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var version = "0.1"

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "palegen h s l",
	Short: "A CSS color palette generator",
	Long: `A CSS colour palette generator:

Specify a base color in hexadecimal  eg. ffcc00
  palegen ffcc00`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {

		ic := "#" + args[0]

		// hc = color input as colorful.color
		hc, _ := colorful.Hex(ic)

		// bc = colorful.color converted to hcl space (hue, chroma, lightness)
		h, c, l := hc.Hcl()

		fmt.Println("Input color:", ic)
		fmt.Printf("bc type is: %T \n", hc)
		fmt.Println("coloful.color:", hc)
		fmt.Println("Converted to HCL space:", h, c, l)
	},
}

func generateHues() {
	// rotate the hue value around the full 360 degress with 12 steps
	// r := 360/12
	// for i = 0 to 12, i++
	// newHue = h +(i * r)
	// need these in some sort of array (slice or map?)
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.palegen.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".palegen") // name of config file (without extension)
	viper.AddConfigPath("$HOME")    // adding home directory as first search path
	viper.AutomaticEnv()            // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
