#!/usr/bin/env bats

load test_helpers

setup() {
  setup_command
  export_errors
}

@test "get versions successfully" {
  run $VCC_CMD get versions -p vmware_tools -s vmtools
  echo $output
  [[ "$output" == *"11.0.0"* ]]
  [ "$status" -eq 0 ]
}

@test "get versions with invalid product" {
  run $VCC_CMD get versions -p INCORRECT -s vmtools
  echo $output
  [[ "$output" == "$ERRORINVALIDSLUG"* ]]
  [ "$status" -eq 1 ]
}

@test "get versions with invalid subproduct" {
  run $VCC_CMD get versions -p vmware_tools -s INCORRECT
  echo $output
  [[ "$output" == "$ERRORINVALIDSUBPRODUCT"* ]]
  [ "$status" -eq 1 ]
}