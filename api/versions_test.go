// Copyright 2022 VMware, Inc.
// SPDX-License-Identifier: Apache 2.0

package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/vmware-labs/vmware-customer-connect-sdk/sdk"
)

func TestGetVersions(t *testing.T) {
	versions, err := ListVersions("vmware_tools", "vmtools")
	require.Nil(t, err)
	assert.Greater(t, len(versions), 10, "Expected response to contain at least 10 items")
}

func TestGetVersionsInvalidSlug(t *testing.T) {
	versions, err := ListVersions("tools", "vmtools")
	assert.ErrorIs(t, err, sdk.ErrorInvalidSlug)
	assert.Empty(t, versions, "Expected response to be empty")
}

func TestGetVersionsInvalidSubProduct(t *testing.T) {
	versions, err := ListVersions("vmware_tools", "tools")
	assert.ErrorIs(t, err, sdk.ErrorInvalidSubProduct)
	assert.Empty(t, versions, "Expected response to be empty")
}
