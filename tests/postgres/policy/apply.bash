#!/bin/bash

main() {
  for func in tests/postgres/_functions tests/postgres/policy/_functions; do
    source "$func"
  done
  
  _gm_prepare_roles || exit 0

  _load_and_run_tests policy

  exit $?
}

main "$@"

