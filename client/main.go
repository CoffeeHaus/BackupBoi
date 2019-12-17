package main

import ("fmt"
        "os"
        "path/filepath"
       )

// this is a comment
func windowmain(){
    }
func linuxmain(){
    }
func main() {
    if runtime.GOOS == "windows" {
    fmt.Println("Windows")
        }
    else if runtime.GOOS == 'linux' {    // also can be specified to FreeBSD
    fmt.Println('Unix/Linux type OS detected')
        }
    fmt.Println("Hello World")
}
