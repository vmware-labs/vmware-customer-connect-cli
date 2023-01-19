// Copyright 2022 VMware, Inc.
// SPDX-License-Identifier: Apache 2.0

package cmd

import (
	"fmt"
	"os"

	"github.com/vmware-labs/vmware-customer-connect-sdk/sdk"
)

func handleErrors(err error) {
	if err != nil {
		if err == sdk.ErrorInvalidSlug {
			fmt.Fprintln(os.Stderr, "Invalid slug provided")
		} else if err == sdk.ErrorInvalidSubProduct {
			fmt.Fprintln(os.Stderr, "Invalid sub-product provided")
		} else if err == sdk.ErrorInvalidVersion {
			fmt.Fprintln(os.Stderr, "Invalid version provided")
		} else if err == sdk.ErrorNoMatchingVersions {
			fmt.Fprintln(os.Stderr, "Version glob did not match any files")
		} else if err == sdk.ErrorNoMatchingFiles {
			fmt.Fprintln(os.Stderr, "No matching files for provided glob")
		} else if err == sdk.ErrorMultipleMatchingFiles {
			fmt.Fprintln(os.Stderr, "Glob matches multiple files, must be restricted to match a single file")
		} else if err == sdk.ErrorEulaUnaccepted {
			fmt.Fprintln(os.Stderr, "Eula has not been accepted for this sub-product")
		} else if err == sdk.ErrorNotEntitled {
			fmt.Fprintln(os.Stderr, "You are not entitled to download this sub-product")
		} else if err == sdk.ErrorNoVersinGlob {
			fmt.Fprintln(os.Stderr, "No version glob provided")
		} else if err == sdk.ErrorMultipleVersionGlob {
			fmt.Fprintln(os.Stderr, "Multiple version globs not supported")
		} else if err == sdk.ErrorAuthenticationFailure {
			fmt.Fprintln(os.Stderr, "Authentication failure. Check your username and/or password!")
		} else if err == sdk.ErrorConnectionFailure {
			fmt.Fprintln(os.Stderr, "Unable to connect to customerconnect.vmware.com. Check your proxy settings.")
		} else {
			fmt.Fprintf(os.Stderr, "%e\n", err)
		}
		os.Exit(1)
	}
}
