#!/usr/bin/env bash

make main -C "$1" --silent || exit 2
