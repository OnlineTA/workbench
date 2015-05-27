package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "os"
)

func mktmp() string {
  var tmpdir string
  var err error

  tmpdir, err = ioutil.TempDir(os.TempDir(), "workbench")
  if err != nil {
    if os.IsExist(err) {
      log.Fatalf("Too many temp directories for Go to handle.")
    }
    // TODO: maybe other error conditions (like out-of-space)
  }

  return tmpdir
}

func rmtmp(tmpdir string) {
  var err error

  err = os.RemoveAll(tmpdir) // TODO: shred
  if err != nil {
    log.Fatalf("Couldn't remove %s.", tmpdir)
  }

  // TODO: err is redundant, but we should be shreding anyhow..
}

func client(assignment string, submission string) {
  var tmpdir string

  tmpdir = mktmp()
  defer rmtmp(tmpdir)

  fmt.Println("HEJ!")
}
