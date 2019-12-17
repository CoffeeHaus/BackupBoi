package main

import "fmt"

// this is a comment

func main() {
    if runtime.GOOS == "windows" {
    fmt.Println("Hello from Windows")
        }
    else if runtime.GOOS == 'linux' {    // also can be specified to FreeBSD
    fmt.Println('Unix/Linux type OS detected')
        }
    fmt.Println("Hello World")
}
