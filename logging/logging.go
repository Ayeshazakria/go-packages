package logging

import (
   "fmt"
    "time"
    "runtime"
)

var debug bool

func Debug(b bool) {
    debug = b
}

func Log(statement string) {
    if !debug {
        return
    }

    fmt.Printf("%s %s\n", time.Now().Format(time.RFC3339), statement)
    fmt.Println(runtime.Caller(0))
}
