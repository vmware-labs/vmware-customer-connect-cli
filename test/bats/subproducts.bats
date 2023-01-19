#!/usr/bin/env bats

load test_helpers

setup() {
  setup_command
  export_errors
}

@test "get subproducts successfully" {
  run $VCC_CMD get subproducts -p vmware_tools
  echo $output
  [[ "$output" == *"VMware Tools"* ]]
  [ "$status" -eq 0 ]
}

@test "get subproducts with invalid product" {
  run $VCC_CMD get subproducts -p INCORRECT
  echo $output
  [[ "$output" == *"$ERRORINVALIDSLUG"* ]]
  [ "$status" -eq 1 ]
}