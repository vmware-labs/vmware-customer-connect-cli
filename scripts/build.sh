#!/bin/bash

set -eu

[ -z "${1:-}" ] && echo 'The version must be passed in as a arguement. E.g. build.sh 0.1.0' && exit 1

readonly verison=$1
readonly script_dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )/.."

env GOOS=linux GOARCH=amd64 go build -o ${script_dir}/builds/vcc-linux-v${verison}

env GOOS=darwin GOARCH=amd64 go build -o ${script_dir}/builds/vcc-darwin-v${verison}

env GOOS=windows GOARCH=amd64 go build -o ${script_dir}/builds/vcc-windows-v${verison}.exe
