// Copyright 2022 VMware, Inc.
// SPDX-License-Identifier: Apache 2.0

package manifest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnsureInitialized(t *testing.T) {
	var testSpec = ManifestSpec{
		Slug:          "test",
		SubProduct:    "test",
		Version:       "test",
		FilenameGlobs: []string{"test"},
	}
	err := ensureInitialised(testSpec, 0)
	assert.Nil(t, err)
}

func TestEnsureInitializedInvalid(t *testing.T) {
	var testSpec = ManifestSpec{
		Slug:       "test",
		SubProduct: "test",
		Version:    "test",
	}
	err := ensureInitialised(testSpec, 0)
	assert.ErrorIs(t, err, ErrorInvalidSpec)
}
