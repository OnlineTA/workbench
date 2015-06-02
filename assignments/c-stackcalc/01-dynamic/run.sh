#!/usr/bin/env bash

cd "$1"

./main <<EOF
5
5
0
EOF

EXIT_CODE="$?"
if [ "$EXIT_CODE" != 0 ] ; then
  echo "Din lommeregner afsluttede med kode $EXIT_CODE."
  echo "Vi forventede at den afsluttede med kode 0."
  exit 2
fi
