# About
This toy application leverages linux system calls to perform a very basic operation: creating a directory. When used in conjunction with seccomp lab 2, you will see what happens when a syscall is attempted after a seccomp 'syscall firewall' is implemented.

# How to use
- Ensure that `build-essential` is installed if running on Ubuntu (other linux distros: ensure that developer tools are installed!)
- Install the `libseccomp-dev` package (ubuntu) or equivalenton other platforms
- Modify the seccomp permissions to remove `mkdirat`
- go get
- go build -o seccomp2
