#!/usr/bin/env bats

load test_helpers

setup() {
  setup_command
  export_errors
  export_yamls
  export TEMP_DIR="$(mktemp -dt bats.XXXXX)"
}

teardown() {
  rm -rf $TEMP_DIR
  rm -f ${HOME}/vcc-downloads/VMware-Tools-*.zip
  rm -f ${HOME}/vcc-downloads/nsx-lcp-*-le.zip
  echo ""
}

@test "download single file successfully to temp" {
  $VCC_CMD logout
  rm -f $TEMP_DIR/*
  local cmd="$VCC_CMD download -p vmware_horizon_clients -s cart+andrd_x8632 -v 2106 -f VMware-Horizon-Client-AndroidOS-x86-*-store.apk --accepteula -o $TEMP_DIR"
  echo $cmd
  run $cmd
  echo "$output"
  [[ "$output" != *"No output directory set."* ]]
  [[ "$output" == *"Collecting download payload"* ]]
  [[ "$output" == *"Download started to"* ]]
  [[ "$output" == *"Download finished"* ]]
  [ "$status" -eq 0 ]
  [ -f $TEMP_DIR/VMware-Horizon-Client-*.apk ]
}

@test "re-download single file successfully to temp" {
  $VCC_CMD logout
  rm -f $TEMP_DIR/*
  local cmd="$VCC_CMD download -p vmware_tools -s vmtools -v 11.3.0 -f VMware-Tools-darwin-*.gz --accepteula -o $TEMP_DIR"
  echo "$cmd"
  run $cmd
  run $cmd
  echo "$output"
  [[ "$output" != *"No output directory set."* ]]
  [[ "$output" == *"Collecting download payload"* ]]
  [[ "$output" == *"Checksum validate completed successfully. No need to re-download."* ]]
  [ "$status" -eq 0 ]
  [ -f $TEMP_DIR/VMware-Tools-darwin-*.gz ]
  local cmd="$VCC_CMD download -p vmware_tools -s vmtools -v 11.3.0 -f VMware-Tools-darwin-*.gz --accepteula -o $TEMP_DIR --forcedownload"
  echo "$cmd"
  run $cmd
  echo "$output"
  [[ "$output" != *"No output directory set."* ]]
  [[ "$output" == *"Collecting download payload"* ]]
  [[ "$output" == *"Download started to"* ]]
  [[ "$output" == *"Download finished"* ]]
  [ "$status" -eq 0 ]
  [ -f $TEMP_DIR/VMware-Tools-darwin-*.gz ]
}

@test "download single file successfully to user vcc-downloads" {
  rm -f $TEMP_DIR/*
  local cmd="$VCC_CMD download -p vmware_tools -s vmtools -v 11.3.0 -f VMware-Tools-darwin-*.zip --accepteula"
  echo $cmd
  run $cmd
  echo "$output"
  [[ "$output" == *"No output directory set."* ]]
  [[ "$output" == *"Collecting download payload"* ]]
  [[ "$output" == *"Download started to"* ]]
  [[ "$output" == *"Download finished"* ]]
  [ "$status" -eq 0 ]
  [ -f $HOME/vcc-downloads/VMware-Tools-darwin-11.3.0-*.zip ]
}

@test "download multiple files successfully to temp" {
  rm -f $TEMP_DIR/*
  run $VCC_CMD download -p vmware_tools -s vmtools -v 11.3.0 -f VMware-Tools-darwin-* --accepteula -o $TEMP_DIR
  echo "$output"
  [[ "$output" != *"No output directory set."* ]]
  [[ "$output" == *"Collecting download payload"* ]]
  [[ "$output" == *"Download started to"* ]]
  [[ "$output" == *"Download finished"* ]]
  [ "$status" -eq 0 ]
  [ -f $TEMP_DIR/VMware-Tools-darwin-*.zip ]
  [ -f $TEMP_DIR/VMware-Tools-darwin-*.tar.gz ]
}

@test "download from manifest" {
  rm -f $TEMP_DIR/*
  run $VCC_CMD download -m <(echo "$VALID_YAML") --accepteula -o $TEMP_DIR
  echo "$output"
  [[ "$output" == *"Opening manifest file:"* ]]
  [[ "$output" == *"Collecting download payload"* ]]
  [[ "$output" == *"Download started to"* ]]
  [[ "$output" == *"Download finished"* ]]
  [ "$status" -eq 0 ]
  ls -l $TEMP_DIR
  [ -f $TEMP_DIR/VMware-Tools-darwin-*.zip ]
  [ -f $TEMP_DIR/VMware-Tools-darwin-*.tar.gz ]
  [ -f $TEMP_DIR/VMware-Tools-other-*.tar.gz ]
}

@test "download from manifest missing field" {
  run $VCC_CMD download -m <(echo "$INVALID_YAML_MISSING_FIELD") --accepteula -o $TEMP_DIR
  echo "$output"
  [[ "$output" == *"Opening manifest file:"* ]]
  [[ "$output" == *"Manifest entry 0 does not have the 4 required keys!"* ]]
  [[ "$output" != *"Collecting download payload"* ]]
  [[ "$output" != *"Download started to"* ]]
  [[ "$output" != *"Download finished"* ]]
  [ "$status" -eq 1 ]
}

@test "download from manifest invalid type" {
  run $VCC_CMD download -m <(echo "$INVALID_YAML_INVALID_TYPE") --accepteula -o $TEMP_DIR
  echo "$output"
  [[ "$output" == *"Opening manifest file:"* ]]
  [[ "$output" == *"Parsing file failed with error:"* ]]
  [[ "$output" != *"Collecting download payload"* ]]
  [[ "$output" != *"Download started to"* ]]
  [[ "$output" != *"Download finished"* ]]
  [ "$status" -eq 1 ]
}

@test "download with invalid product" {
  run $VCC_CMD download -p INVALID -s vmtools -v 11.3.0 -f VMware-Tools-darwin-*.zip --accepteula
  echo "$output"
  [[ "$output" == *"$ERRORINVALIDSLUG"* ]]
  [ "$status" -eq 1 ]
}

@test "download with invalid subproduct" {
  run $VCC_CMD download -p vmware_tools -s INVALID -v 11.3.0 -f VMware-Tools-darwin-*.zip --accepteula
  echo "$output"
  [[ "$output" == *"$ERRORINVALIDSUBPRODUCT"* ]]
  [ "$status" -eq 1 ]
}

@test "download with invalid version" {
  run $VCC_CMD download -p vmware_tools -s vmtools -v INVALID -f VMware-Tools-darwin-*.zip --accepteula
  echo "$output"
  [[ "$output" == *"$ERRORINVALIDVERSION"* ]]
  [ "$status" -eq 1 ]
}

@test "download with invalid credentials" {
  $VCC_CMD logout
  run $VCC_CMD download -p vmware_tools -s vmtools -v 11.3.0 -f VMware-Tools-darwin-*.zip --accepteula --user invalid --pass invalid
  echo "$output"
  [[ "$output" == *"$ERRORAUTHENTICATIONFAILURE"* ]]
  [ "$status" -eq 1 ]
}

@test "download when not entitled" {
  $VCC_CMD logout
  rm -f ${HOME}/vcc-downloads/VMware-VMvisor-Installer-.7*.iso
  run $VCC_CMD download -p vmware_nsx -s nsx_le -v '4.*' -f nsx-lcp-*-le.zip --accepteula
  echo "$output"
  [[ "$output" == *"$ERRORNOTENTITLED"* ]]
  [ "$status" -eq 1 ]
}

@test "download with invalid output directory" {
  run $VCC_CMD download -p vmware_tools -s vmtools -v 11.3.0 -f VMware-Tools-darwin-*.zip --accepteula -o /tmp/stilton/on/toast
  echo "$output"
  [[ "$output" == *"ERROR: Output directory"* ]]
  [ "$status" -eq 1 ]
}