#!/usr/bin/env bats

load test_helpers

setup() {
  setup_command
}

@test "get products successfully" {
  run $VCC_CMD get products
  echo $output
  [[ "$output" == *"vmware_tools"* ]]
  [ "$status" -eq 0 ]
}