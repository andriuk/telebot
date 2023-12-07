/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var appVersion = "Version"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Get actual version of kBot application",
	Long:  `Returns version of Telegram kBot application`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(appVersion)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
