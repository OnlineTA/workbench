#!/usr/bin/env bash

cd "$1"

./main <<EOF
5
5
+
p
p
EOF

echo $?
