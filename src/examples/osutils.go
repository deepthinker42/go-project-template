package main

import "os"
import "fmt"
import "path"


func Exists(name string) bool {
    _, err := os.Stat(name)
    if os.IsNotExist(err) {
        return false
    }
    return err == nil
}

func main() {
    home, _ := os.UserHomeDir()
    p := path.Join(home, ".bashrc")
    if Exists(p) {
        fmt.Printf("%v exists\n", p)
    } else {
        fmt.Printf("%v does not exist\n", p)
    }
    p = path.Join(home, ".bashrc2")
    if Exists(p) {
        fmt.Printf("%v exists\n", p)
    } else {
        fmt.Printf("%v does not exist\n", p)
    }
}

