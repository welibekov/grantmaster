#!/bin/bash

export POSTGRES_DOCKER_CONTAINER=${POSTGRES_DOCKER_CONTAINER:-postgres}
export POSTGRES_USER=${POSTGRES_USER:-dwh}
export POSTGRES_DATABASE=${POSTGRES_DATABASE:-dwh}
export GM_DATABASE_TYPE=postgres
export GM_POSTGRES_CONN_STRING='postgres://dwh:dwh@localhost:5432/dwh'

export TEST_DIR=$(mktemp -d)

_exec() {
  docker exec -ti "$POSTGRES_DOCKER_CONTAINER" "$@"
}

_psql_read() {
  _exec cat /root/output
}

_psql() {
  _exec psql --output=/root/output -U "$POSTGRES_USER" -d "$POSTGRES_DATABASE" "$@"
}

_role() {
  $GM_BIN role "$@"
}

_cmp() {
  cmp -s "$@"
}

main() {
  cat <<'EOF'
/////////////////////////
//                     //
// Starting role tests //
EOF

  _test_number_one
}

_test_number_one() {
  echo -n "> Test #1 ... "

  cat <<'EOF' >"$TEST_DIR/role.yaml"
- name: song_read
  schemas:
    - schema: song
      grants:
        - usage
EOF

  _role apply "$TEST_DIR/role.yaml" || exit 1
  _role get > "$TEST_DIR/get_roles.yaml" || exit 1

 local exitcode; exitcode=0

 if ! _cmp "$TEST_DIR/role.yaml" "$TEST_DIR/get_roles.yaml"; then
   diff "$TEST_DIR/role.yaml" "$TEST_DIR/get_roles.yaml"
   exitcode=1
 fi

 if [[ $exitcode -eq 0 ]]; then
   echo done
 else
   echo failed
 fi
}


main "$@"

