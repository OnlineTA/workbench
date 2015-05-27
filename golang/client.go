package main

import (
  "fmt"
  "io"
  "io/ioutil"
  "log"
  "os"
  "path"
  "path/filepath"
)

func mktmp() string {
  var tmpdir string
  var err error

  tmpdir, err = ioutil.TempDir(os.TempDir(), "workbench")
  if err != nil {
    if os.IsExist(err) {
      log.Fatal("Too many temp directories for Go to handle.")
    }
    // TODO: maybe other error conditions (like out-of-space)
  }

  err = os.Mkdir(path.Join(tmpdir, "assignment"), 0700)
  if err != nil {
    log.Fatal("I/O operation failed in temp directory.")
  }

  err = os.Mkdir(path.Join(tmpdir, "submission"), 0700)
  if err != nil {
    log.Fatal("I/O operation failed in temp directory.")
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

func copyrec(src string, dst string) error {
  return filepath.Walk(src,
    func(path string, info os.FileInfo, err error) error {

      in, err := os.Open(src)
      if err != nil {
        return err
      }
      defer in.Close()

      out, err := os.Create(dst)
      if err != nil {
        return err
      }
      defer func() {
        cerr := out.Close()
        if err == nil {
          err = cerr
        }
      }()
      if _, err = io.Copy(out, in); err != nil {
        return err
      }
      err = out.Sync()
      return nil
    })
}

func client(assignment_src string, submission_src string) {
  var tmpdir string
  var assignment_dst string
  var submission_dst string

  tmpdir = mktmp()
  defer rmtmp(tmpdir)

  assignment_dst = path.Join(tmpdir, "assignment")
  submission_dst = path.Join(tmpdir, "submission")

  copyrec(assignment_src, assignment_dst)
  copyrec(submission_src, submission_dst)

  fmt.Println("HEJ!")
}
