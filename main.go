package main

import (
    "bufio"
    "flag"
    "log"
    "os"
)

func main() {
    fptr := flag.String("fpath",os.Args[1],"file path to read from")
    flag.Parse()

    f, err := os.Open(*fptr)
    if err != nil {
        log.Fatal(err)
    }

    defer func() {
        if err = f.Close(); err != nil {
            log.Fatal(err)
        }
    }()

    s := bufio.NewScanner(f)

    threadTraces := make(map[uint64][]interface{})

    parse(s, &threadTraces)

    process(&threadTraces)

    err = s.Err()
    if err != nil {
        log.Fatal(err)
    }
}
