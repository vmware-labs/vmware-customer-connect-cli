// Copyright 2022 VMware, Inc.
// SPDX-License-Identifier: Apache 2.0

package api

import (
	"net/http"
	"strings"

	"github.com/vmware-labs/vmware-customer-connect-sdk/sdk"
)

var basicClient = sdk.Client{HttpClient: &http.Client{}}

func ListProducts() (data [][]string, err error) {
	products, err := basicClient.GetProductsSlice()
	if err != nil {
		return
	}

	for _, v := range products {
		slug := strings.Split(v.MajorProductEntities[0].Target, "/")[4]
		line := []string{slug, v.Name}
		data = append(data, line)
	}
	return
}
