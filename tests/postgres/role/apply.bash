#!/bin/bash

export TEST_DIR=$(mktemp -d)

main() {
  source tests/postgres/role/_tests

  cat <<EOF
/////////////////////////////////
//                             //
// Starting role tests         //
// Results are under:          //
// $TEST_DIR/roles   //
//                             //
/////////////////////////////////
EOF
  tests=()
  
  while IFS= read -r test; do
    tests+=("$test")
  done <<<$(declare -F | grep -Eo "_test_number_[0-9]{1,}")

  for test_name in "${tests[@]}"; do
    local num; num=${test_name##*_}

    "$test_name" "$num" 
  done

  cat <<'EOF'
/////////////////////////////////
//                             //
// Stopping role tests         //
//                             //
/////////////////////////////////
EOF
}

main "$@"

