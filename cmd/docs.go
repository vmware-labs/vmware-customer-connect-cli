// Copyright 2022 VMware, Inc.
// SPDX-License-Identifier: Apache 2.0

package cmd

// Usage Section
const (
	getProductsUsage = `  # List of available products
  vcc get products`

	getSubProductsUsage = `  # List of available sub-products of product vmware_tools
  vcc get subproducts -p vmware_tools`

	getVersions = `  # List of available versions of sub-products vmtools of vmware_tools
  vcc get versions -p vmware_tools -s vmtools`

  getMajorVersions = `  # List of available major versions of product vmware_tools
  vcc get versions -p vmware_tools`

	getFiles = `  # List of available files of version 11.3.0 of vmware_tools
  vcc get files -p vmware_tools -s vmtools -v 11.3.0`

	getManifestExample = `  # Display example manifest file
  vcc get manifestexample`

	downloadUsage = `  # Download the latest version of release 11 with a file matching the pattern
  # If using a * in the filename value, make sure to wrap the text in single quotes on linux/macos
  vcc download -p vmware_tools -s vmtools -v 11.* -f 'VMware-Tools-darwin-*.zip' --accepteula

  # Download files using a manifest file
  # Show an example manifest using 'vcc get manifestexample'
  vcc download -m manifest.yml --accepteula`
)

const exampleManifest = `---
# This section will download the latest version of vmware_tools
# Each glob pattern will download a single file each
product: vmware_tools
subproduct: vmtools
version: "*"
filename_globs:
  - "VMware-Tools-darwin-*.tar.gz"
  - "VMware-Tools-darwin-*.zip"
---
# This section will download the latest minor release from major version 10
# The single glob pattern will download 2 files
product: vmware_tools
subproduct: vmtools
version: "10.*"
filename_globs:
  - "VMware-Tools-other-*"
---`
