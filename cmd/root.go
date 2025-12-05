/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"strings"

	"github.com/meyuviofficial/kilikili/constants"
	"github.com/meyuviofficial/kilikili/utils"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kilikili",
	Short: "kilikili is a CLI tool for generating project names",
	Long: `kilikili is a CLI tool for generating project names. It can generate project names based on the category you want. E.g. Country, Animal, Color, etc.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		name1 := utils.RandomNameGenerator(constants.Colours)
		name2 := utils.RandomNameGenerator(constants.Animals)
		strings.Replace(name1, " ", "",-1)
		strings.Replace(name2, " ", "", -1)

		directoryName := name1 + "-" + name2
		directoryName = strings.ToLower(directoryName)
		utils.MakeDirectory(directoryName)
	 },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.kilikili.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


