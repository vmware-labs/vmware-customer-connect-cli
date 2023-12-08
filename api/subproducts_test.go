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
	products, err := ListSubProducts("vmware_tools", "PRODUCT_BINARY", "")
	require.Nil(t, err)
	assert.NotEmpty(t, products)
}

func TestGetSubProductsDriversMajorVersion(t *testing.T) {
	var products [][]string
	products, err := ListSubProducts("vmware_vsphere", "DRIVERS_TOOLS", "8_0")
	require.Nil(t, err)
	assert.NotEmpty(t, products)
	assert.LessOrEqual(t, len(products), 400)
}

func TestGetSubProductsInvalidSlug(t *testing.T) {
	products, err := ListSubProducts("tools", "PRODUCT_BINARY", "")
	assert.ErrorIs(t, err, sdk.ErrorInvalidSlug)
	assert.Empty(t, products, "Expected response to be empty")
}
