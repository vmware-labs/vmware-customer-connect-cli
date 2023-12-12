// Copyright 2022 VMware, Inc.
// SPDX-License-Identifier: Apache 2.0

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vmware-labs/vmware-customer-connect-cli/api"
)

var majorVersionsSlug string

// versionsCmd represents the versions command
var MajorVersionsCmd = &cobra.Command{
	Use:     "majorversions",
	Aliases: []string{"v"},
	Short:   "List available majors versions (for driver_tools, custom_iso, addons)",
	Long:    "List available major versions to help query drives_tools, custom_isos and addons",
	Example: getMajorVersions,
	Run: func(cmd *cobra.Command, args []string) {
		majorVersions, err := api.GetMajorVersionsString(majorVersionsSlug)
		handleErrors(err)
		fmt.Println(majorVersions)
	},
}

func init() {
	getCmd.AddCommand(MajorVersionsCmd)
	MajorVersionsCmd.Flags().StringVarP(&majorVersionsSlug, "product", "p", "", "Product code")
	MajorVersionsCmd.MarkFlagRequired("product")
}
