package main

import (
  "log"
  "syscall"
  "time"
)

func main() {
  dt := time.Now().Format("2006-01-02T15-04-05")
  errMkdir := syscall.Mkdir("tmpdir-"+dt, 0644)
  if errMkdir != nil{
     panic(errMkdir)
  }

  log.Println("Created tmpdir via syscall")
}
