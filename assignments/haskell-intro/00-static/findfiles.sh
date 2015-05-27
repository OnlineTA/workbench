#!/usr/bin/env bash

TMPDIR=$1

function findHaskell {
  find "$TMPDIR" -name "intro.hs" &> /dev/null || {
    echo "We can't find your Haskell code!"
    echo "Remember that your code file should be called 'intro.hs'."
    echo "There is a very low chance of passing without Haskell code!"
  }
}

find "$TMPDIR" -name "report.txt" &> /dev/null || {
  echo "We can't find your report!"
  echo "Remember that your report should be called 'report.txt'."
  echo "There is a very low chance of passing without a report!"

  # Don't exit, let's check if there's some haskell code around..

  findHaskell || {
    echo "No report, no haskell, no cry."
    exit 2
  }

  exit 0
}

findHaskell || exit 2
