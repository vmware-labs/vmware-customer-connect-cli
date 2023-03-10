// Copyright 2022 VMware, Inc.
// SPDX-License-Identifier: Apache 2.0

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vmware-labs/vmware-customer-connect-cli/api"
)

var subProduct string

// versionsCmd represents the versions command
var versionsCmd = &cobra.Command{
	Use:     "versions",
	Aliases: []string{"v"},
	Short:   "List available versions",
	Long:    "List available versions of a sub-product",
	Example: getVersions,
	Run: func(cmd *cobra.Command, args []string) {
		versionString, err := api.ListVersions(slug, subProduct)
		if err != nil {
			handleErrors(err)
		}
		fmt.Println(versionString)
	},
}

func init() {
	getCmd.AddCommand(versionsCmd)
	versionsCmd.Flags().StringVarP(&slug, "product", "p", "", "Product code")
	versionsCmd.Flags().StringVarP(&subProduct, "subproduct", "s", "", "Sub Product code")
	versionsCmd.MarkFlagRequired("product")
	versionsCmd.MarkFlagRequired("sub-product")
}
