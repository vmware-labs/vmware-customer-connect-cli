// Copyright 2022 VMware, Inc.
// SPDX-License-Identifier: Apache 2.0

package api

import (
	"fmt"
	"os"

	"github.com/vmware-labs/vmware-customer-connect-sdk/sdk"
)

type Availability struct {
	EulaAccepted       bool
	EligibleToDownload bool
}

func ListFiles(slug, subProduct, version, username, password string) (dlgDetails sdk.DlgDetails, apiVersions sdk.APIVersions, err error) {
	if err = EnsureLogin(username, password); err != nil {
		return
	}

	var productID string
	productID, apiVersions, err = authenticatedClient.GetDlgProduct(slug, subProduct, version)
	if err != nil {
		return
	}

	fmt.Fprintf(os.Stderr, "Getting DLG Details\n")
	dlgDetails, err = authenticatedClient.GetDlgDetails(apiVersions.Code, productID)
	if err != nil {
		return
	}

	return
}

func ListFilesArray(slug, subProduct, version, username, password string) (data [][]string, availability Availability, apiVersions sdk.APIVersions, err error) {
	dlgDetails, apiVersions, err := ListFiles(slug, subProduct, version, username, password)
	if err != nil {
		return
	}
	for _, v := range dlgDetails.DownloadDetails {
		if v.FileName != "" {
			line := []string{v.FileName, v.FileSize, v.Build, v.Title}
			data = append(data, line)
		}
	}
	availability = Availability{
		EulaAccepted:       dlgDetails.EulaResponse.EulaAccepted,
		EligibleToDownload: dlgDetails.EligibilityResponse.EligibleToDownload,
	}
	return
}