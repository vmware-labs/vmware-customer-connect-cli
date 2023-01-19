// Copyright 2022 VMware, Inc.
// SPDX-License-Identifier: Apache 2.0

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var username string
var password string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Version: "0.2.0",
	Use:     "vcc",
	Short:   "Download binaries from customerconnect.vmware.com",
	Long:    "vcc downloads binaries from customerconnect.vmware.com",
	Example: fmt.Sprintf("%s\n\n%s\n\n%s\n\n%s\n\n%s\n\n%s", downloadUsage, getProductsUsage, getSubProductsUsage, getVersions, getFiles, getManifestExample),
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.PersistentFlags().StringVar(&username, "user", "", "Username used to authenticate [$VCC_USER]")
	rootCmd.PersistentFlags().StringVar(&password, "pass", "", "Password used to authenticate [$VCC_PASS]")
}
