// Copyright 2022 VMware, Inc.
// SPDX-License-Identifier: Apache 2.0

package manifest

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

type ManifestSpec struct {
	Slug          string   `yaml:"product"`
	SubProduct    string   `yaml:"subproduct"`
	Version       string   `yaml:"version"`
	FilenameGlobs []string `yaml:"filename_globs"`
}

var (
	ErrorFileDoesNotExist = errors.New("manifest file does not exist")
	ErrorInvalidSpec      = errors.New("manifest file has invalid section")
)

func ProcessFile(manifestFile string) (downloads []ManifestSpec, err error) {
	if !strings.Contains(manifestFile, "/") {
		cwd, _ := filepath.Abs("./")
		fmt.Println(cwd)
		manifestFile = cwd + "/" + manifestFile
	}

	if _, fileErr := os.Stat(manifestFile); os.IsNotExist(fileErr) {
		err = errors.New("manifest file not found")
		os.Exit(2)
	}

	f, err := os.Open(manifestFile)
	if err != nil {
		return
	}
	d := yaml.NewDecoder(f)

	entry := 0
	for {
		// create new spec here
		download := new(ManifestSpec)
		// pass a reference to spec reference
		err = d.Decode(&download)
		// check it was parsed
		if download == nil {
			continue
		}

		// break the loop in case of EOF
		if errors.Is(err, io.EOF) {
			err = nil
			break
		}
		if err != nil {
			return
		}

		err = ensureInitialised(*download, entry)
		if err != nil {
			return
		}
		entry++
		downloads = append(downloads, *download)
	}
	return
}

func ensureInitialised(dl ManifestSpec, entry int) (err error) {
	if (dl.Slug == "") || (dl.SubProduct == "") || (dl.Version == "") || (len(dl.FilenameGlobs) == 0) {
		fmt.Fprintf(os.Stderr, "Manifest entry %d does not have the 4 required keys!\n", entry)
		err = ErrorInvalidSpec
	}
	return
}
