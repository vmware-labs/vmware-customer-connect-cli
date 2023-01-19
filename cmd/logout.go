// Copyright 2022 VMware, Inc.
// SPDX-License-Identifier: Apache 2.0

package cmd

import (
	// "fmt"

	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Remove all session cookies",
	Long:  "Remove all session cookies by deleting .vcc.cookies",
	Run: func(cmd *cobra.Command, args []string) {

		cookieFile := filepath.Join(homeDir(), ".vcc.cookies")
		if _, fileErr := os.Stat(cookieFile); os.IsNotExist(fileErr) {
			fmt.Println("No sessions cookies to delete")
			os.Exit(0)
		}
		err := os.Remove(cookieFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: Unable to delete cookie file [%s].\n", cookieFile)
			fmt.Fprintf(os.Stderr, "%e", err)
			os.Exit(1)
		}
		fmt.Println("Deleted all session cookies")
	},
	DisableFlagsInUseLine: true,
}

func init() {
	rootCmd.AddCommand(logoutCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
