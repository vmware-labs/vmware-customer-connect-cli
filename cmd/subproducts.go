// Copyright 2022 VMware, Inc.
// SPDX-License-Identifier: Apache 2.0

package cmd

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/vmware-labs/vmware-customer-connect-cli/api"
	"github.com/vmware-labs/vmware-customer-connect-cli/presenters"
)

var slug, dlgType string

// subproductsCmd represents the subproducts command
var subproductsCmd = &cobra.Command{
	Use:     "subproducts",
	Aliases: []string{"s"},
	Short:   "List sub-products",
	Long:    "List sub-products for a specified product",
	Example: getSubProductsUsage,
	Run: func(cmd *cobra.Command, args []string) {
		products, err := api.ListSubProducts(slug, dlgType)
		if err != nil {
			handleErrors(err)
		}
		headings := []string{"Sub-Product Code", "Description"}
		presenters.RenderTable(headings, products)
	},
}

func init() {
	getCmd.AddCommand(subproductsCmd)
	subproductsCmd.Flags().StringVarP(&slug, "product", "p", "", "Product code")
	subproductsCmd.MarkFlagRequired("product")
	subproductsCmd.Flags().StringVarP(&dlgType, "type", "t", "product_binary", "(optional) Download type. One of: (product_binary, drivers_tools, custom_iso, addons). Default: product_binary")
	dlgType = strings.ToUpper(dlgType)
}
