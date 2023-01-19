// Copyright 2022 VMware, Inc.
// SPDX-License-Identifier: Apache 2.0

package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/vmware-labs/vmware-customer-connect-sdk/sdk"
)

func TestGetSubProducts(t *testing.T) {
	var products [][]string
	products, err := ListSubProducts("vmware_tools")
	require.Nil(t, err)
	assert.NotEmpty(t, products)
}

func TestGetSubProductsInvalidSlug(t *testing.T) {
	versions, err := ListVersions("tools", "vmtools")
	assert.ErrorIs(t, err, sdk.ErrorInvalidSlug)
	assert.Empty(t, versions, "Expected response to be empty")
}
