// Copyright 2022 VMware, Inc.
// SPDX-License-Identifier: Apache 2.0

package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/vmware-labs/vmware-customer-connect-cli/api"
	"github.com/vmware-labs/vmware-customer-connect-cli/presenters"
)

var dlgType, majorVersion, slug string

// subproductsCmd represents the subproducts command
var subproductsCmd = &cobra.Command{
	Use:     "subproducts",
	Aliases: []string{"s"},
	Short:   "List sub-products",
	Long:    "List sub-products for a specified product",
	Example: getSubProductsUsage,
	Run: func(cmd *cobra.Command, args []string) {
		products, err := api.ListSubProducts(slug, dlgType, majorVersion)
		handleErrors(err)
		headings := []string{"Sub-Product Code", "Description"}
		presenters.RenderTable(headings, products)

		if dlgType == "drivers_tools" && len(products) > 200 {
			majorVersions, err := api.GetMajorVersionsString(slug)
			handleErrors(err)
			fmt.Println("\nDue to the high number of results it's recommended to provide the major version to cut down results.")
			fmt.Fprintf(os.Stdout, "Major Versions: %s\n", majorVersions)
			fmt.Println("\nE.g. vcc get subproducts -p vmware_vsphere -t drivers_tools -m '7_0'")
		}
	},
}

func init() {
	getCmd.AddCommand(subproductsCmd)
	subproductsCmd.Flags().StringVarP(&slug, "product", "p", "", "Product code")
	subproductsCmd.MarkFlagRequired("product")
	subproductsCmd.Flags().StringVarP(&dlgType, "type", "t", "product_binary", "(optional) Download type. One of: (product_binary, drivers_tools, custom_iso, addons). Default: product_binary")
	dlgType = strings.ToUpper(dlgType)
	subproductsCmd.Flags().StringVarP(&majorVersion, "majorversion", "m", "", "(optional) Reduce the number of results by passing in the major verion.\nGet versions with `vcc get majorversions -p vmware_vsphere`")
}
