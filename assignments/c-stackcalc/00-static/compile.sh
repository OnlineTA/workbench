#!/usr/bin/env bash

cd "$1"

make main --silent || exit 2
