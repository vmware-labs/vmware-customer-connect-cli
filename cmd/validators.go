// Copyright 2022 VMware, Inc.
// SPDX-License-Identifier: Apache 2.0

package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/spf13/cobra"
)

func validateOutputDir() {
	if outputDir == "" {
		home := homeDir()
		homeDownloadDir := filepath.Join(home, "vcc-downloads")
		if _, dir_err := os.Stat(homeDownloadDir); os.IsNotExist(dir_err) {
			fmt.Printf("%s does not exist. Creating...\n", homeDownloadDir)
			mkdir_err := os.MkdirAll(homeDownloadDir, os.ModePerm)
			if mkdir_err != nil {
				fmt.Fprintf(os.Stderr, "ERROR: Unable to create Downloads directory under [%s].\n", home)
				os.Exit(1)
			}
		}
		fmt.Printf("No output directory set. Downloading to %s\n", homeDownloadDir)
		outputDir = homeDownloadDir
	} else {
		if _, dir_err := os.Stat(outputDir); os.IsNotExist(dir_err) {
			fmt.Fprintf(os.Stderr, "ERROR: Output directory [%s] does not exist.\n", outputDir)
			os.Exit(1)
		}
	}
}

// homeDir returns the OS-specific home path as specified in the environment.
func homeDir() string {
	if runtime.GOOS == "windows" {
		return filepath.Join(os.Getenv("HOMEDRIVE"), os.Getenv("HOMEPATH"))
	}
	return os.Getenv("HOME")
}

// ensure credentials are passed in and assign env vars if used
func validateCredentials(cmd *cobra.Command) {
	user_ok := validateVarEnv(&username, "VCC_USER")
	pass_ok := validateVarEnv(&password, "VCC_PASS")
	if !user_ok || !pass_ok {
		fmt.Fprintln(os.Stderr, "Credentials not provided!")
		fmt.Fprintln(os.Stderr, "You must either provide the username and password as arguements")
		fmt.Fprintf(os.Stderr, "or you must export them as VCC_USER and VCC_PASS environment variables.\n\n")
		cmd.Usage()
		os.Exit(1)
	}
}

// Check if param is set and if not retrieve env var if set
func validateVarEnv(param *string, key string) bool {
	if *param == "" {
		if value, ok := os.LookupEnv(key); ok {
			*param = value
		} else {
			return false
		}
	}
	return true
}

func validateDownloadFlags(cmd *cobra.Command) (manifestWorkflow bool) {
	if manifestFile != "" {
		manifestWorkflow = true
		return
	} else if (slug != "") && (subProduct != "") &&
		(version != "") && (fileName != "") {
		return
	}
	fmt.Fprintln(os.Stderr, "Incorrect usage!")
	fmt.Fprintln(os.Stderr, "Either --manifest should be passed")
	fmt.Fprintf(os.Stderr, "or --product, --subproduct, --version and --filename should be passed\n\n")
	cmd.Usage()
	os.Exit(2)

	return
}
