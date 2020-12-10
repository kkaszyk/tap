package main

import (
    "fmt"
)

func process(threadTraces *map[uint64][]interface{}) {
    for tid, threadTrace := range *threadTraces {
        fmt.Printf("%d %d\n", tid, len(threadTrace))
        for i := 0; i < len(threadTrace); i++ {
            fmt.Printf("%d %x\n", tid, threadTrace[i])
        }
    }
}


