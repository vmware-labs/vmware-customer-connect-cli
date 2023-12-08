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

@test "get subproducts successfully - drivers" {
  run $VCC_CMD get subproducts -p vmware_vsphere -t drivers_tools
  echo $output
  [[ "$output" == *"ESX"* ]]
  [ "$status" -eq 0 ]
}

@test "get subproducts drivers with version successfully" {
  run $VCC_CMD get subproducts -p vmware_vsphere -t drivers_tool
  echo $output
  [[ "$output" == *"Supported types are:"* ]]
  [ "$status" -eq 1 ]
}

@test "get subproducts drivers with version invalid type" {
  run $VCC_CMD get subproducts -p vmware_vsphere -t drivers_tools -m 8_0
  echo $output
  [[ "$output" == *"ESX"* ]]
  [ "$status" -eq 0 ]
}

@test "get subproducts with invalid product" {
  run $VCC_CMD get subproducts -p INCORRECT
  echo $output
  [[ "$output" == *"$ERRORINVALIDSLUG"* ]]
  [ "$status" -eq 1 ]
}