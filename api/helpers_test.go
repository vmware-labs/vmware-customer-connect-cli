// Copyright 2022 VMware, Inc.
// SPDX-License-Identifier: Apache 2.0

package api

import (
	"fmt"
	"os"
)

var err error

func mustEnv(k string) string {
	if v, ok := os.LookupEnv(k); ok {
		return v
	}
	fmt.Println("Environment variables not set correctly")
	os.Exit(1)
	return ""
}
