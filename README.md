# Go Child Go!

---

## Introduction

This program is designed to create a child process which creates a folder, and child go file,
and then executes the newly created file. 

---

## Usage

To use this program, it takes one parameter argument. The argument must be in single or double
quotes, with the following syntax:

    '<FolderName>/<childname>.go'

To build, and run:

    go build -o goprocgo.exe GoProc.go
    ./goprocgo.exe 'Child/GoChild.go'

---