# Exercise 1 - Getting Started with Go

If you do not finish during the lecture period, please finish it as homework.

## Setup

- Install Go from <http://golang.org> inside a virtual disk
  (.vhdx on Windows or .sparseimage on Mac) in a /software/go subdirectory.

  - Win: <https://www.windowscentral.com/how-create-and-set-vhdx-or-vhd-windows-10>
  - Mac: <https://support.apple.com/de-de/guide/disk-utility/dskutl11888/mac>

- Create a Go workspace on the disk in the <<DISK>>/codebase/gopath directory
  - <https://www.youtube.com/watch?v=XCsL89YtqCs>
- Create a shell script (.sh / .cmd) to make your changes to the `GOPATH` and PATH environment variables persistent.
- Create a Github project with your personal account containing a `HelloWorld.go` program
- Use go get github.com/\<\<YOUR REPO\>\> to copy the repository into your local `GOPATH`.
- Add `HelloWorld.go` to the checkout path and `commit / push` the file.
- Test the `HelloWorld` program with "go run HelloWorld"
- Optional: Install Visual Studio Code, IntelliJ or any other Editor with Go support inside your virtual disk.

## After this Exercise

- You should know how to compile and run Go code
- You should know about the meaning of the `GOPATH` and PATH variables
- You should have a portable Go installation inside a separate disk or directory on your computer
- Your personal Github project is cloned into your workspace. You are able to add, remove and commit files.
