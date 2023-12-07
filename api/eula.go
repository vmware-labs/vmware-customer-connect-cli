// Copyright 2022 VMware, Inc.
// SPDX-License-Identifier: Apache 2.0

package api

import (
	"github.com/vmware-labs/vmware-customer-connect-sdk/sdk"
)


func GetEula(slug, subProduct, version, username, password, dlgType string) (data string, err error) {
	var productID string
	var apiVersions sdk.APIVersions
	productID, apiVersions, err = basicClient.GetDlgProduct(slug, subProduct, version, dlgType)
	if err != nil {
		return
	}
	if err = EnsureLogin(username, password); err != nil {
		return
	}

	data, err = authenticatedClient.FetchEulaUrl(apiVersions.Code, productID)
	return
}
