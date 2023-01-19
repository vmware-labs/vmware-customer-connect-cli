// Copyright 2022 VMware, Inc.
// SPDX-License-Identifier: Apache 2.0

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vmware-labs/vmware-customer-connect-cli/api"
	"github.com/vmware-labs/vmware-customer-connect-cli/presenters"
)

var slug string

// subproductsCmd represents the subproducts command
var subproductsCmd = &cobra.Command{
	Use:     "subproducts",
	Aliases: []string{"s"},
	Short:   "List sub-products",
	Long:    "List sub-products for a specified product",
	Example: getSubProductsUsage,
	Run: func(cmd *cobra.Command, args []string) {
		products, err := api.ListSubProducts(slug)
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
}
