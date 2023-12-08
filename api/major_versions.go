// Copyright 2022 VMware, Inc.
// SPDX-License-Identifier: Apache 2.0

package api

import (
	"fmt"
	"strings"
)

func GetMajorVersionsString(slug string) (majorVersionString string, err error) {
	var majorVersionSlice []string
	majorVersionSlice, err = basicClient.GetMajorVersionsSlice(slug)
	if err != nil {
		return
	}
	majorVersionString = fmt.Sprintf("'%s'", strings.Join(majorVersionSlice,"', '"))
	return
}