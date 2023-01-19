// Copyright 2022 VMware, Inc.
// SPDX-License-Identifier: Apache 2.0

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// productsCmd represents the products command
var manifestExampleCmd = &cobra.Command{
	Use:   "manifestexample",
	Short: "Display an example download manifest",
	Long:  "Display an example download manifest",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(exampleManifest)

	},
	DisableFlagsInUseLine: true,
}

func init() {
	getCmd.AddCommand(manifestExampleCmd)
	manifestExampleCmd.ResetFlags()
}
