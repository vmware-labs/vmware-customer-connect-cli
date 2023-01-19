// Copyright 2022 VMware, Inc.
// SPDX-License-Identifier: Apache 2.0

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vmware-labs/vmware-customer-connect-cli/api"
	"github.com/vmware-labs/vmware-customer-connect-cli/presenters"
)

var version string

// filesCmd represents the files command
var filesCmd = &cobra.Command{
	Use:     "files",
	Aliases: []string{"f"},
	Short:   "List available files",
	Long: `List available files of a version of a sub-product

Either VCC_USER and VCC_PASS environment variable must be set
or the --user and --pass flags should be added`,
	Example: getFiles,
	Run: func(cmd *cobra.Command, args []string) {
		validateCredentials(cmd)
		files, availability, err := api.ListFiles(slug, subProduct, version, username, password)
		if err != nil {
			handleErrors(err)
		}
		headings := []string{"Filename", "Size", "Build number", "Description"}

		fmt.Printf("\nEula Accepted:         %t\n", availability.EulaAccepted)
		fmt.Printf("Eligable to Download:  %t\n\n", availability.EligibleToDownload)
		presenters.RenderTable(headings, files)
	},
}

func init() {
	getCmd.AddCommand(filesCmd)
	filesCmd.Flags().StringVarP(&slug, "product", "p", "", "Product code")
	filesCmd.Flags().StringVarP(&subProduct, "subproduct", "s", "", "Sub Product code")
	filesCmd.Flags().StringVarP(&version, "version", "v", "", "Version string")
	filesCmd.MarkFlagRequired("product")
	filesCmd.MarkFlagRequired("sub-product")
	filesCmd.MarkFlagRequired("version")
}
