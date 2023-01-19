// Copyright 2022 VMware, Inc.
// SPDX-License-Identifier: Apache 2.0

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vmware-labs/vmware-customer-connect-cli/api"
)

// filesCmd represents the files command
var eulaCmd = &cobra.Command{
	Use:     "eula",
	Aliases: []string{"e"},
	Short:   "Display the Eula of a product",
	Long: `Display the eula of a version of a sub-product

Either VCC_USER and VCC_PASS environment variable must be set
or the --user and --pass flags should be added`,
	Example: getFiles,
	Run: func(cmd *cobra.Command, args []string) {
		validateCredentials(cmd)
		eula, err := api.GetEula(slug, subProduct, version, username, password)
		handleErrors(err)
		fmt.Printf("Open the URL in your browser: %s\n", eula)
	},
}

func init() {
	getCmd.AddCommand(eulaCmd)
	eulaCmd.Flags().StringVarP(&slug, "product", "p", "", "Product code")
	eulaCmd.Flags().StringVarP(&subProduct, "subproduct", "s", "", "Sub Product code")
	eulaCmd.Flags().StringVarP(&version, "version", "v", "", "Version string")
	eulaCmd.MarkFlagRequired("product")
	eulaCmd.MarkFlagRequired("sub-product")
	eulaCmd.MarkFlagRequired("version")
}
