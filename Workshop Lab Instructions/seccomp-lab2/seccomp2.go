package main

import (
  "fmt"
  libseccomp "github.com/seccomp/libseccomp-golang"
  "log"
  "syscall"
  "time"
)

func allowSyscalls(syscalls []string){
  // Setup the 'default deny' policy for syscalls not in the list
  returnCode := int16(syscall.EPERM)
  seccompAction := libseccomp.ActErrno.SetReturnCode(returnCode)
  filter, filterErr := libseccomp.NewFilter(seccompAction)
  if filterErr != nil {
     log.Println("Unable to create seccomp filter, Err:", filterErr)
  }
  fmt.Println("Configured default-deny for any non-approved syscalls")

  // Load the approved list of syscalls
  for _, syscallN := range syscalls{
     fmt.Println("Allowing ", syscallN)
     scmpSyscall, scmpSyscallErr := libseccomp.GetSyscallFromName(syscallN)
     if scmpSyscallErr != nil{
        panic(scmpSyscallErr)
     }

  filter.AddRule(scmpSyscall, libseccomp.ActAllow)
  }

  filterLoadErr := filter.Load()
  if filterLoadErr != nil{
     log.Println("Unable to load syscall filter! Err:", filterLoadErr)
  }
}

func main() {

  allowedSyscalls := []string{"rt_sigaction", "mkdirat", "clone", "rt_sigprocmask", "rt_sigreturn", "futex", "fcntl",
     "openat", "read", "write", "mmap", "close", "execve", "sigaltstack", "arch_prctl", "gettid",
     "sched_getaffinity", "exit_group"}
  allowSyscalls(allowedSyscalls)

  dt := time.Now().Format("2006-01-02T15-04-05")
  errMkdir := syscall.Mkdir("tmpdir-"+dt, 0644)
  if errMkdir != nil{
     log.Println("Unable to create directory via syscall!")
     panic(errMkdir)
  }

  log.Println("Created tmpdir via syscall")
}
