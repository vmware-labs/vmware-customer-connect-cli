// Copyright 2022 VMware, Inc.
// SPDX-License-Identifier: Apache 2.0

package cmd

import (
	// "fmt"

	"fmt"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var getCmd = &cobra.Command{
	Use:                   "get",
	Aliases:               []string{"g"},
	Short:                 "Display responses",
	Long:                  `Display responses`,
	Example:               fmt.Sprintf("%s\n\n%s\n\n%s\n\n%s\n\n%s", getProductsUsage, getSubProductsUsage, getVersions, getFiles, getManifestExample),
	DisableFlagsInUseLine: true,
}

func init() {
	rootCmd.AddCommand(getCmd)
}
