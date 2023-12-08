// Copyright 2022 VMware, Inc.
// SPDX-License-Identifier: Apache 2.0

package cmd

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/vmware-labs/vmware-customer-connect-cli/api"
	"github.com/vmware-labs/vmware-customer-connect-cli/downloader"
	"github.com/vmware-labs/vmware-customer-connect-cli/manifest"
	"github.com/vmware-labs/vmware-customer-connect-sdk/sdk"
)

var (
	acceptEula    bool
	fileName      string
	forceDownload bool
	outputDir     string
	manifestFile  string
)

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:     "download",
	Aliases: []string{"d"},
	Short:   "Download file from VMware",
	Long: `Download one or more files

Either VCC_USER and VCC_PASS environment variable must be set
or the --user and --pass flags should be added`,
	Example: downloadUsage,
	Run: func(cmd *cobra.Command, args []string) {
		dlgType = validateDlgType(dlgType)
		validateCredentials(cmd)
		validateOutputDir()
		manifestWorkflow := validateDownloadFlags(cmd)
		err := api.EnsureLogin(username, password)
		handleErrors(err)
		if manifestWorkflow {
			downloadFromManifest()
		} else {
			fmt.Println("Collecting download payload")
			downloadPayloads, err := api.FetchDownloadPayload(slug, subProduct, version, fileName, username, password, dlgType, acceptEula)
			handleErrors(err)
			downloadFiles(downloadPayloads)
		}
	},
}

func downloadFromManifest() {
	fmt.Printf("Opening manifest file: %s\n", manifestFile)
	manifestArray, err := manifest.ProcessFile(manifestFile)
	if err == manifest.ErrorFileDoesNotExist {
		fmt.Fprintf(os.Stderr, "File %s does not exist\n", manifestFile)
		os.Exit(1)
	} else if err == manifest.ErrorInvalidSpec {
		os.Exit(1)
	} else if err != nil {
		fmt.Fprintf(os.Stderr, "Parsing file failed with error: %e\n", err)		
		os.Exit(1)
	}

	var allPayloads [][]sdk.DownloadPayload
	for _, manifestSpec := range manifestArray {
		for _, glob := range manifestSpec.FilenameGlobs {
			fmt.Printf("Collecting download payload for [%s] [%s] [%s] [%s]\n", manifestSpec.Slug, manifestSpec.SubProduct,
				manifestSpec.Version, glob)
			downloadPayloads, err := api.FetchDownloadPayload(manifestSpec.Slug, manifestSpec.SubProduct, manifestSpec.Version,
				glob, username, password, dlgType, acceptEula)
			handleErrors(err)
			allPayloads = append(allPayloads, downloadPayloads)
		}
	}

	for _, downloadPayloads := range allPayloads {
		downloadFiles(downloadPayloads)
	}

}

func downloadFiles(downloadPayloads []sdk.DownloadPayload) {
	for _, downloadPayload := range downloadPayloads {
		authorizedDownload, err := api.FetchDownloadLink(downloadPayload, username, password)
		handleErrors(err)
		authorizedDownload.FileName = filepath.Join(outputDir, authorizedDownload.FileName)
		if forceDownload || checkToDownload(authorizedDownload.FileName, downloadPayload.Md5checksum) {
			err = downloader.TriggerDownload(authorizedDownload)
			handleErrors(err)
		}
	}
}

func checkToDownload(fileName string, expectedMD5 string) bool {
	if fileExists(fileName) {
		fmt.Printf("Found file %s, calculating MD5 checksum to validate\n", fileName)
		file, err := os.Open(fileName)
		handleErrors(err)
		defer file.Close()

		// Create a hash instance and pass the file through it
		hash := md5.New()
		_, err = io.Copy(hash, file)
		handleErrors(err)
		// Usage for Sprintf needed as a standard string conversation broke some strings
		calculatedMD5 := fmt.Sprintf("%x", hash.Sum(nil))

		if expectedMD5 != calculatedMD5 {
			fmt.Printf("Expected checksum of [%s], but found [%s].\nAttempting to re-download.\n", expectedMD5, calculatedMD5)
			return true
		} else {
			fmt.Println("Checksum validate completed successfully. No need to re-download.")
			return false
		}
	}
	return true
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func init() {
	rootCmd.AddCommand(downloadCmd)
	downloadCmd.Flags().StringVarP(&slug, "product", "p", "", "Product code")
	downloadCmd.Flags().StringVarP(&subProduct, "subproduct", "s", "", "Sub Product code")
	downloadCmd.Flags().StringVarP(&version, "version", "v", "", "Version string. Can contain a glob.")
	downloadCmd.Flags().StringVarP(&fileName, "filename", "f", "", "Filename string. Can contain one or more globs. When using * wrap the text in single quotes.")
	downloadCmd.Flags().StringVarP(&manifestFile, "manifest", "m", "", "Filename of the manifest containing details of what to download")
	downloadCmd.Flags().StringVarP(&outputDir, "output", "o", "", "Directory to download files to")
	downloadCmd.Flags().BoolVarP(&acceptEula, "accepteula", "a", false, "Filename string")
	downloadCmd.Flags().BoolVarP(&forceDownload, "forcedownload", "d", false, "(optional) Force a file to be re-downloaded even if it already exists")
	downloadCmd.Flags().StringVarP(&dlgType, "type", "t", "product_binary", "(optional) Download type. One of: (product_binary, drivers_tools, custom_iso, addons). Default: product_binary")
}
