// Copyright 2022 VMware, Inc.
// SPDX-License-Identifier: Apache 2.0

package downloader

//  Credit https://golangcode.com/download-a-file-with-progress/

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/dustin/go-humanize"
	"github.com/vmware-labs/vmware-customer-connect-sdk/sdk"
)

var ErrorGeneric = errors.New("download: non 200 response")

type WriteCounter struct {
	Total uint64
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += uint64(n)
	wc.PrintProgress()
	return n, nil
}

func (wc WriteCounter) PrintProgress() {
	// Clear the line by using a character return to go back to the start and remove
	// the remaining characters by filling it with spaces
	fmt.Printf("\r%s", strings.Repeat(" ", 35))

	// Return again and print current status of download
	// We use the humanize to print the bytes in a meaningful way (e.g. 10 MB)
	fmt.Printf("\rDownloading... %s complete", humanize.Bytes(wc.Total))
}

func TriggerDownload(authorizedDownload sdk.AuthorizedDownload) (err error) {
	fmt.Printf("\nDownload started to %s\n", authorizedDownload.FileName)

	err = DownloadFile(authorizedDownload.DownloadURL, authorizedDownload.FileName)
	if err != nil {
		return
	}

	fmt.Printf("Download finished\n")
	return
}

func DownloadFile(url string, fileName string) error {

	// Create the file, but give it a tmp file extension, this means we won't overwrite a
	// file until it's downloaded, but we'll remove the tmp extension once downloaded.
	out, err := os.Create(fileName + ".tmp")
	if err != nil {
		return err
	}

	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	if err != nil {
		out.Close()
		return err
	}
	defer resp.Body.Close()

	// Create our progress reporter and pass it to be used alongside our writer
	counter := &WriteCounter{}
	if _, err = io.Copy(out, io.TeeReader(resp.Body, counter)); err != nil {
		out.Close()
		return err
	}

	// The progress use the same line so print a new line once it's finished downloading
	fmt.Print("\n")

	// Close the file without defer so it can happen before Rename()
	out.Close()

	if resp.StatusCode != 200 {
		os.Remove(fileName + ".tmp")
		return ErrorGeneric
	}

	if err = os.Rename(fileName+".tmp", fileName); err != nil {
		return err
	}
	return nil
}
