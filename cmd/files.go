// Copyright 2022 VMware, Inc.
// SPDX-License-Identifier: Apache 2.0

package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/vmware-labs/vmware-customer-connect-cli/api"
	"github.com/vmware-labs/vmware-customer-connect-cli/presenters"
	"github.com/vmware-labs/vmware-customer-connect-sdk/sdk"
)

type JsonOutput struct {
	EulaAccepted bool `json:"eula_accepted"`
	EligibleToDownload bool `json:"eligible_to_download"`
	Version string `json:"version"`
	Files []sdk.DownloadDetails `json:"files"`
}

var version, outputFormat string

// filesCmd represents the files command
var filesCmd = &cobra.Command{
	Use:     "files",
	Aliases: []string{"f"},
	Short:   "List available files",
	Long: `List available files of a version of a sub-product

Either VCC_USER and VCC_PASS environment variable must be set
or the --user and --pass flags should be added`,
	Example: getFiles,
	Run: func(cmd *cobra.Command, args []string) {
		dlgType = validateDlgType(dlgType)
		if !(outputFormat == "text" || outputFormat == "json") {
			fmt.Fprintf(os.Stderr, "Format type %s is not supported\n", outputFormat)
			os.Exit(128)
		}

		validateCredentials(cmd)
		if outputFormat == "text" {
			files, availability, apiVersions, err := api.ListFilesArray(slug, subProduct, version, username, password, dlgType)
			handleErrors(err)
			printText(apiVersions, availability, files)
		} else if outputFormat == "json" {
			dlgDetails, apiVersions, err := api.ListFiles(slug, subProduct, version, username, password, dlgType)
			handleErrors(err)
			printJson(dlgDetails, apiVersions)
		}
	},
}

func printText(apiVersions sdk.APIVersions, availability api.Availability, files [][]string) {
	fmt.Printf("\nVersion:               %s\n", apiVersions.MinorVersion)
	fmt.Printf("Eula Accepted:         %t\n", availability.EulaAccepted)
	fmt.Printf("Eligable to Download:  %t\n\n", availability.EligibleToDownload)
	
	headings := []string{"Filename", "Size", "Build number", "Description"}
	presenters.RenderTable(headings, files)
}

func printJson(dlgDetails sdk.DlgDetails, apiVersions sdk.APIVersions) {
	var jsonOutput JsonOutput
	jsonOutput.Version = apiVersions.MinorVersion
	jsonOutput.EligibleToDownload = dlgDetails.EligibilityResponse.EligibleToDownload
	jsonOutput.EulaAccepted = dlgDetails.EulaResponse.EulaAccepted
	jsonOutput.Files = dlgDetails.DownloadDetails

	var p []byte
	//    var err := error
	p, err := json.MarshalIndent(jsonOutput, "", "\t")
	if err != nil {
			fmt.Println(err)
			return
	}
	fmt.Printf("%s \n", p)
}

func init() {
	getCmd.AddCommand(filesCmd)
	filesCmd.Flags().StringVarP(&slug, "product", "p", "", "Product code")
	filesCmd.Flags().StringVarP(&subProduct, "subproduct", "s", "", "Sub Product code")
	filesCmd.Flags().StringVarP(&version, "version", "v", "", "Version string")
	filesCmd.Flags().StringVarP(&outputFormat, "format", "", "text", "Format of command output. Options [text, json]")
	filesCmd.MarkFlagRequired("product")
	filesCmd.MarkFlagRequired("sub-product")
	filesCmd.MarkFlagRequired("version")
	filesCmd.Flags().StringVarP(&dlgType, "type", "t", "product_binary", "(optional) Download type. One of: (product_binary, drivers_tools, custom_iso, addons). Default: product_binary")
}
