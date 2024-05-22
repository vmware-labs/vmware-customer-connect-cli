#!/bin/bash
# Copyright 2022 VMware, Inc.
# SPDX-License-Identifier: Apache 2.0

set -eu

[ -z "${1:-}" ] && echo 'The version must be passed in as a arguement. E.g. build.sh 0.1.0' && exit 1

readonly verison=$1
readonly script_dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )/.."

env GOOS=linux GOARCH=amd64 go build \
  -ldflags="-X 'github.com/vmware-labs/vmware-customer-connect-cli/cmd.cliVersion=${verison}'" \
  -o ${script_dir}/builds/vcc-linux-amd64-v${verison}

env GOOS=darwin GOARCH=amd64 go build \
  -ldflags="-X github.com/vmware-labs/vmware-customer-connect-cli/cmd.cliVersion=${verison}" \
  -o ${script_dir}/builds/vcc-darwin-amd64-v${verison}

env GOOS=darwin GOARCH=arm64 go build \
  -ldflags="-X github.com/vmware-labs/vmware-customer-connect-cli/cmd.cliVersion=${verison}" \
  -o ${script_dir}/builds/vcc-darwin-arm64-v${verison}

env GOOS=windows GOARCH=amd64 go build \
  -ldflags="-X github.com/vmware-labs/vmware-customer-connect-cli/cmd.cliVersion=${verison}" \
  -o ${script_dir}/builds/vcc-windows-amd64-v${verison}.exe
