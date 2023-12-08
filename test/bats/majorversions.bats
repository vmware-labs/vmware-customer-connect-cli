#!/usr/bin/env bats

load test_helpers

setup() {
  setup_command
  export_errors
}

@test "get major versions successfully" {
  run $VCC_CMD get majorversions -p vmware_vsphere
  echo $output
  [[ "$output" == *"8_0"* ]]
  [ "$status" -eq 0 ]
}

@test "get major versions with invalid product" {
  run $VCC_CMD get majorversions -p vmware_open_shift
  echo $output
  [[ "$output" == "$ERRORINVALIDSLUG"* ]]
  [ "$status" -eq 1 ]
}
