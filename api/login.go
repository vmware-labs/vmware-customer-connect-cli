// Copyright 2022 VMware, Inc.
// SPDX-License-Identifier: Apache 2.0

package api

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/vmware-labs/vmware-customer-connect-sdk/sdk"
	"github.com/orirawlings/persistent-cookiejar"
)

var authenticatedClient *sdk.Client
var jar *cookiejar.Jar

func EnsureLogin(username, password string) (err error) {
	if authenticatedClient == nil {
		// Store cookies under the user profile
		jar, err = cookiejar.New(&cookiejar.Options{
			Filename:              filepath.Join(homeDir(), ".vcc.cookies"),
			PersistSessionCookies: true,
		})
		if err != nil {
			return
		}
		fmt.Println("Logging in...")
		authenticatedClient, err = sdk.Login(username, password, jar)
		if err == nil {
			err = jar.Save()
		}
	} else {
		// If tokens are still valid leave existing authenticatedClient in place
		err = authenticatedClient.CheckLoggedIn()
		if err == nil {
			return
		}

		authenticatedClient, err = sdk.Login(username, password, jar)
		if err == nil {
			err = jar.Save()
		}
	}
	return
}

// homeDir returns the OS-specific home path as specified in the environment.
func homeDir() string {
	if runtime.GOOS == "windows" {
		return filepath.Join(os.Getenv("HOMEDRIVE"), os.Getenv("HOMEPATH"))
	}
	return os.Getenv("HOME")
}
