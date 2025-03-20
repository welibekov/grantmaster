#!/bin/bash

main() {
  source tests/postgres/_functions
  
  _load_and_run_tests role
   
  exit $?
}

main "$@"

