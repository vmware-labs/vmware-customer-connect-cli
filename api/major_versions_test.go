// Copyright 2022 VMware, Inc.
// SPDX-License-Identifier: Apache 2.0

package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	// "github.com/vmware-labs/vmware-customer-connect-sdk/sdk"
)

func TestGetGetMajorVersionsString(t *testing.T) {
	var majorVersions string
	majorVersions, err := GetMajorVersionsString("vmware_vsphere")
	require.Nil(t, err)
	assert.Contains(t, majorVersions, "8_0")
}
