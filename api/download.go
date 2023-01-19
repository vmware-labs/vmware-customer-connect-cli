// Copyright 2022 VMware, Inc.
// SPDX-License-Identifier: Apache 2.0

package api

import (
	"github.com/vmware-labs/vmware-customer-connect-sdk/sdk"
)

func FetchDownloadPayload(slug, subProduct, version, fileName, username, password string, acceptEula bool) (data []sdk.DownloadPayload, err error) {
	if err = EnsureLogin(username, password); err != nil {
		return
	}

	data, err = authenticatedClient.GenerateDownloadPayload(slug, subProduct, version, fileName, acceptEula)
	if err != nil {
		return
	}
	return
}

func FetchDownloadLink(downloadPayload sdk.DownloadPayload, username, password string) (data sdk.AuthorizedDownload, err error) {
	err = EnsureLogin(username, password)
	if err != nil {
		return
	}
	data, err = authenticatedClient.FetchDownloadLink(downloadPayload)
	return
}
