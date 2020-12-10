package main

import (
    "bufio"
    "fmt"
    "strconv"
    "strings"
)

func parse(s *bufio.Scanner, threadTraces *map[uint64][]interface{}) {
    var tid uint64
    tid = 0
    count := 0

    threadTraceMap := *threadTraces

    for s.Scan() {
        e := s.Text()
        
        switch e[0] {
        case 'T':
            trace_entry := strings.Split(e, " ")
            tid, _ = strconv.ParseUint(trace_entry[1],16,64)
            threadTraceMap[tid] = make([]interface{},0, 100)
        case 'J':
            NewJDTraceEntry(e)
        case 'P':
            threadTraceMap[tid] = append(threadTraceMap[tid], NewArithmeticTraceEntry([]rune(e)[2:]))
        case 'L':
            threadTraceMap[tid] = append(threadTraceMap[tid], NewLoadTraceEntry(e))
        case 'S':
            threadTraceMap[tid] = append(threadTraceMap[tid], NewStoreTraceEntry(e))
        case 'B':
            threadTraceMap[tid] = append(threadTraceMap[tid], NewBarrierTraceEntry(e))
        default:
        }
        count += 1
    }
    fmt.Println(count)
}
