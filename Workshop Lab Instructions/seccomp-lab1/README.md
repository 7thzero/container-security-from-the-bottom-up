# About
This toy application leverages linux system calls to perform a very basic operation: creating a directory. When used in conjunction with seccomp lab 1, you can profile the application to identify which syscalls are used by the application as a whole

# How to use
- Ensure that `build-essential` is installed if running on Ubuntu (other linux distros: ensure that developer tools are installed!)
- Install the `libseccomp-dev` package (ubuntu) or equivalenton other platforms
- go get
- go build -o seccomp1
