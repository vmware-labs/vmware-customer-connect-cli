// Copyright 2022 VMware, Inc.
// SPDX-License-Identifier: Apache 2.0

package api

import (
	"fmt"
	"strings"
)

func ListVersions(slug, subProduct string) (data string, err error) {
	versionArray, err := basicClient.GetVersionSlice(slug, subProduct)
	if err != nil {
		return
	}

	data = strings.Join(versionArray[:], "' '")
	data = fmt.Sprintf("'%s'", data)

	return
}
