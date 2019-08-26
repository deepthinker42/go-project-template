package main

import (
        "fmt"
        "os"
        "flag"
       )

func main() {
    var DIRS = []string{"api", "assets", "build", "cmd", "configs", "deployments", "docs", "examples", "githooks", "init", "internal", "pkg", "scripts", "test", "third_party", "tools", "vendor", "web", "website"}

    project := flag.String("p", "", "Project name")
    flag.Parse()
    if *project == "" {
        fmt.Println("usage:  go-dir-layout -p <project>")
        os.Exit(1)
    }

    os.Mkdir(*project, 0755)
    for _, d := range DIRS {
        fmt.Println(d)
        var path = *project + "/" + d
        os.Mkdir(path, 0755)
        if d == "cmd" {
            path += "/" + *project
            os.Mkdir(path, 0755)
        }
    }
}
