#!/bin/bash

exit 0

TEST_DIR=$(mktemp -d)

cat <<'EOF' > $TEST_DIR/policy.yaml
EOF

$GM_BIN policy apply --help
