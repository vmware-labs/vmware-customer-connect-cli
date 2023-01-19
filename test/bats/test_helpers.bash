setup_command() {
  SCRIPT_DIR="$( cd "$( dirname "$BATS_TEST_FILENAME" )" >/dev/null 2>&1 && pwd )/../.."
  export VCC_CMD="go run $SCRIPT_DIR/main.go"
}

export_errors() {
  export ERRORINVALIDSLUG="Invalid slug provided"
  export ERRORINVALIDSUBPRODUCT="Invalid sub-product provided"
	export ERRORINVALIDVERSION="Invalid version provided"
	export ERRORNOMATCHINGVERSIONS="Version glob did not match any files"
	export ERRORNOMATCHINGFILES="No matching files for provided glob"
	export ERRORMULTIPLEMATCHINGFILES="Glob matches multiple files, must be restricted to match a single file"
	export ERROREULAUNACCEPTED="Eula has not been accepted for this sub-product"
	export ERRORNOTENTITLED="You are not entitled to download this sub-product"
	export ERRORNOVERSINGLOB="No version glob provided"
	export ERRORMULTIPLEVERSIONGLOB="Multiple version globs not supported"
	export ERRORAUTHENTICATIONFAILURE="Authentication failure. Check your username and/or password!"
}

export_yamls() {
	export INVALID_YAML_MISSING_FIELD="---
product: vmware_tools
subproduct: vmtools
version: \"11.*\"
"

	export INVALID_YAML_INVALID_TYPE="---
product: vmware_tools
subproduct: vmtools
version: 11.0
filename_globs: INVALID
"

	export VALID_YAML="---
product: vmware_tools
subproduct: vmtools
version: \"11.*\"
filename_globs:
  - VMware-Tools-darwin-*.tar.gz
  - VMware-Tools-darwin-*.zip
---
product: vmware_tools
subproduct: vmtools
version: \"10.*\"
filename_globs:
  - VMware-Tools-other-*.tar.gz
---"
}
