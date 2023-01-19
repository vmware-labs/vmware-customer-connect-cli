// Copyright 2022 VMware, Inc.
// SPDX-License-Identifier: Apache 2.0

package api

import (
	// "fmt"
	// "os"

	"github.com/vmware-labs/vmware-customer-connect-sdk/sdk"
)

func ListSubProducts(slug string) (data [][]string, err error) {
	var subProducts []sdk.SubProductDetails
	subProducts, err = basicClient.GetSubProductsSlice(slug)
	if err != nil {
		return
	}
	for _, v := range subProducts {
		line := []string{v.ProductCode, v.ProductName}
		data = append(data, line)
	}

	return
}
