// Copyright (c) Jargon Author(s) 2020. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

package cmd

import (
	"fmt"
	"os"

	"github.com/imdevin567/jargon/config"

	"github.com/morikuni/aec"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const jargonASCIILogo = `
   oooo
   '888
    888  .oooo.   oooo d8b  .oooooooo  .ooooo.  ooo. .oo.
    888 'P  )88b  '888""8P 888' '88b  d88' '88b '888P"Y88b
    888  .oP"888   888     888   888  888   888  888   888
    888 d8(  888   888     '88bod8P'  888   888  888   888
.o. 88P 'Y888""8o d888b    '8oooooo.  'Y8bod8P' o888o o888o
'Y888P                     d"     YD
                           "Y88888'
`

// Version ...
var Version string
var jargonfile string

var jargonCmd = &cobra.Command{
	Use:   "jargon",
	Short: "",
	Long:  "",
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display the version information.",
	Run:   parseBaseCommand,
}

func init() {
	cobra.OnInitialize(initConfig)
	jargonCmd.PersistentFlags().StringVar(&jargonfile, "jargonfile", "Jargonfile", "Config file")
	jargonCmd.AddCommand(versionCmd)
}

func initConfig() {
	viper.SetConfigType("yaml")
	viper.SetConfigFile(jargonfile)
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("No Jargonfile specified!")
	}

	err := viper.Unmarshal(&config.Config)
	if err != nil {
		fmt.Println(err)
	}

	config.Config.CreateEntries()
}

func getVersion() string {
	if len(Version) != 0 {
		return Version
	}

	return "development"
}

func parseBaseCommand(_ *cobra.Command, _ []string) {
	printLogo()

	fmt.Println("Version:", getVersion())
	os.Exit(0)
}

func printLogo() {
	jargonLogo := aec.WhiteF.Apply(jargonASCIILogo)
	fmt.Println(jargonLogo)
}

// Execute ...
func Execute(version string) error {
	Version = version

	if err := jargonCmd.Execute(); err != nil {
		return err
	}

	return nil
}
