// Copyright 2022 VMware, Inc.
// SPDX-License-Identifier: Apache 2.0

package api

import (
	"fmt"
	"strings"
	// "github.com/vmware-labs/vmware-customer-connect-sdk/sdk"
)

func ListVersions(slug, subProduct string) (data string, err error) {
	versionArray, err := basicClient.GetVersionSlice(slug, subProduct)
	if err != nil {
		return
	}

	data = strings.Join(versionArray[:], "' '")
	data = fmt.Sprintf("'%s'", data)

	// fmt.Println(versionString)
	return
}
