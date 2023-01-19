// Copyright 2022 VMware, Inc.
// SPDX-License-Identifier: Apache 2.0

package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetProducts(t *testing.T) {
	var products [][]string
	products, err := ListProducts()
	require.Nil(t, err)
	assert.Greater(t, len(products), 80, "Expected response to contain at least 80 items")
}
