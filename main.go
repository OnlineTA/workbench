package main

import (
  "flag"
  "fmt"
  "os"
)

func usage() {
  fmt.Fprintf(os.Stderr,
    "Usage: %s ASSIGNMENT SUBMISSION\n", os.Args[0])
  flag.PrintDefaults()
}

func main() {
  var isServer bool
  var assignment string
  var submission string

  flag.Usage = usage
  flag.BoolVar(&isServer, "server", false,
    "run in server mode (here be dragons)")
  flag.Parse()

  args := flag.Args()
  if len(args) != 2 {
    flag.Usage()
    os.Exit(1)
  }

  assignment = args[0]
  submission = args[1]

  if ! isServer {
    client(assignment, submission)
  }
}
