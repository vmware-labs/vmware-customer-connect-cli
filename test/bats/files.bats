#!/usr/bin/env bats

load test_helpers

setup() {
  setup_command
  export_errors
}

@test "get files successfully_text" {
  run $VCC_CMD get files -p vmware_tools -s vmtools -v 11.3.0
  echo $output
  [[ "$output" == *"  11.3.0"* ]]
  [[ "$output" == *"Eula Accepted:"* ]]
  [[ "$output" == *"Eligable to Download:  true"* ]]
  [[ "$output" == *"VMware-Tools-windows-11.3.0-18090558.zip"* ]]
  [ "$status" -eq 0 ]
}

@test "get files successfully_json" {
  run $VCC_CMD get files -p vmware_tools -s vmtools -v 11.3.0 --format json
  echo $output
  [[ "$output" == *"eula_accepted"* ]]
  [[ "$output" == *"eligible_to_download"* ]]
  [[ "$output" == *"sha256checksum"* ]]
  [ "$status" -eq 0 ]
}

@test "get files with invalid product" {
  run $VCC_CMD get files -p INVALID -s vmtools -v 11.3.0
  echo $output
  [[ "$output" == *"$ERRORINVALIDSLUG"* ]]
  [ "$status" -eq 1 ]
}

@test "get files with invalid subproduct" {
  run $VCC_CMD get files -p vmware_tools -s INVALID -v 11.3.0
  echo $output
  [[ "$output" == *"$ERRORINVALIDSUBPRODUCT"* ]]
  [ "$status" -eq 1 ]
}

@test "get files with invalid version" {
  run $VCC_CMD get files -p vmware_tools -s vmtools -v INVALID
  echo $output
  [[ "$output" == *"$ERRORINVALIDVERSION"* ]]
  [ "$status" -eq 1 ]
}

@test "get files with invalid credentials" {
  $VCC_CMD logout
  run $VCC_CMD get files -p vmware_tools -s vmtools -v INVALID --user invalid --pass invalid
  echo $output
  [[ "$output" == *"$ERRORAUTHENTICATIONFAILURE"* ]]
  [ "$status" -eq 1 ]
}