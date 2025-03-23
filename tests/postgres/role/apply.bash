#!/bin/bash

main() {
  for func in tests/postgres/_functions tests/postgres/role/_functions; do
    source "$func"
  done
  
  _load_and_run_tests role
   
  exit $?
}

main "$@"

