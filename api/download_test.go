// Copyright 2022 VMware, Inc.
// SPDX-License-Identifier: Apache 2.0

package api

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/vmware-labs/vmware-customer-connect-sdk/sdk"
)

func TestFetchDownloadLinkVersionGlob(t *testing.T) {
	var downloadPayload []sdk.DownloadPayload
	downloadPayload, err = FetchDownloadPayload("vmware_tools", "vmtools", "11.*", "VMware-Tools-darwin-*.tar.gz", testing_user, testing_pass, true)
	require.Nil(t, err)
	require.NotEmpty(t, downloadPayload)
	assert.NotEmpty(t, downloadPayload[0].DlgType, "Expected response not to be empty")

	var authorizedDownload sdk.AuthorizedDownload
	authorizedDownload, _ = authenticatedClient.FetchDownloadLink(downloadPayload[0])
	assert.Nil(t, err)
	assert.NotEmpty(t, authorizedDownload.DownloadURL, "Expected response not to be empty")

	t.Logf(fmt.Sprintf("download_details: %+v\n", authorizedDownload))
}

func TestFetchDownloadPayloadVersionGlobMultiple(t *testing.T) {
	var downloadPayload []sdk.DownloadPayload
	downloadPayload, err = FetchDownloadPayload("vmware_tools", "vmtools", "11.*", "VMware-Tools-*", testing_user, testing_pass, true)
	require.Nil(t, err)
	require.NotEmpty(t, downloadPayload)
	assert.NotEmpty(t, downloadPayload[0].DlgType, "Expected response not to be empty")
	assert.Greater(t, len(downloadPayload), 3)
}

func TestFetchDownloadLinkInvalidVersion(t *testing.T) {
	var downloadPayload []sdk.DownloadPayload
	downloadPayload, err = FetchDownloadPayload("vmware_tools", "vmtools", "666", "VMware-Tools-darwin-*.tar.gz", testing_user, testing_pass, true)
	assert.ErrorIs(t, err, sdk.ErrorInvalidVersion)
	assert.Empty(t, downloadPayload, "Expected response to be empty")
}

func TestFetchDownloadLinkNeedEula(t *testing.T) {
	var downloadPayload []sdk.DownloadPayload
	downloadPayload, err = FetchDownloadPayload("vmware_tools", "vmtools", "11.1.0", "VMware-Tools-darwin-*.tar.gz", testing_user, testing_pass, false)
	assert.ErrorIs(t, err, sdk.ErrorEulaUnaccepted)
	assert.Empty(t, downloadPayload, "Expected response to be empty")
}

func TestFetchDownloadLinkNotEntitled(t *testing.T) {
	var downloadPayload []sdk.DownloadPayload
	downloadPayload, err = FetchDownloadPayload("vmware_nsx_t_data_center", "nsx-t", "3.2.3.1", "nsx-unified-appliance-secondary-*.qcow2", testing_user, testing_pass, true)
	assert.ErrorIs(t, err, sdk.ErrorNotEntitled)
	assert.Empty(t, downloadPayload, "Expected response to be empty")
}

func TestGenerateDownloadInvalidVersionGlob(t *testing.T) {
	var downloadPayload []sdk.DownloadPayload
	downloadPayload, err = FetchDownloadPayload("vmware_tools", "vmtools", "666.*", "VMware-Tools-darwin-*.tar.gz", testing_user, testing_pass, true)
	assert.ErrorIs(t, err, sdk.ErrorNoMatchingVersions)
	assert.Empty(t, downloadPayload, "Expected response to be empty")
}

func TestGenerateDownloadDoubleVersion(t *testing.T) {
	var downloadPayload []sdk.DownloadPayload
	downloadPayload, err = FetchDownloadPayload("vmware_tools", "vmtools", "*.*", "VMware-Tools-darwin-*.tar.gz", testing_user, testing_pass, true)
	assert.ErrorIs(t, err, sdk.ErrorMultipleVersionGlob)
	assert.Empty(t, downloadPayload, "Expected response to be empty")
}
