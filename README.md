# container-security-from-the-bottom-up
Presentation on Linux Isolation Mechanisms Given at BSidesSLC 2021 on Friday, December 3rd, 2021

Shoutout to everyone at https://www.bsidesslc.org, we had a great time and look forward to next year!


# Presentation Overview
0x00 - Lab Setup
- Setup the virtual machine

0x01 - Introduction
- Marketing and selling 'Containers'
- Current state of container ecosystem
- ...but what underpins it all?

0x02 - Overview of Isolation
- Thirty-thousand foot view
- What we're covering today
- What we're not coverig today

0x03 - Linux Namespaces
- What they are and how they are used
- Lab worksto exercise understanding

0x04 - cgroups
- Purpose, function: resource isolation and control
- Lab work

0x05 - seccomp
- A story of syscalls and filtering
- Labs on how to use them

0x06 Closing Summary

0x07 - References & Tips


# Summary Blurb
Most companies have experimented with containerization in one form or another and there is a wealth of information at these higher abstraction levels for securing constructs like Dockerfiles, Container Images, and Container Orchestration tools like Kubernetes. Linux kernel features like namespaces, seccomp, capabilities and cgroups provide the resource isolation functionality that underpin these tools.

This workshop aims to provide participants with an understanding of low level kernel features that form the basis of modern container technologies, with an emphasis on Namespaces and seccomp. At the end of the session, participants will have experience implementing isolation features and better understand their limits as well as the opportunities for use.

Participants should bring a laptop with 2 or more cpus and 8GB or more RAM to run VMware workstation or player and the training VM that will be provided during class.

Who should attend? Anyone who has a desire to understand at a deeper level what is happening ‘behind the scenes’ with containers. This could include Engineers, DevOps, security folks or other technically-minded people who want to look behind the curtain. A basic understanding of linux is recommended and the ability to decipher or write scripts and simple programs is optional/beneficial.
