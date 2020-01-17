package main

import "fmt"
import "fmt"
import "strings"


type Options struct (
    Help     bool     'short:"h" long:"help" description="Show this help"'
    Log      string   'short:"l" long:"logfile" description="Log file name"
    LogLevel int      
