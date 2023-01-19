// Copyright 2022 VMware, Inc.
// SPDX-License-Identifier: Apache 2.0

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vmware-labs/vmware-customer-connect-cli/api"
	"github.com/vmware-labs/vmware-customer-connect-cli/presenters"
)

// productsCmd represents the products command
var productsCmd = &cobra.Command{
	Use:     "products",
	Aliases: []string{"p"},
	Short:   "List of available products",
	Long:    "List of available products",
	Example: getProductsUsage,
	Run: func(cmd *cobra.Command, args []string) {

		products, err := api.ListProducts()
		if err != nil {
			handleErrors(err)
		}
		headings := []string{"Product code", "Product description"}
		presenters.RenderTable(headings, products)
	},
}

func init() {
	getCmd.AddCommand(productsCmd)
}
