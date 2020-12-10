package main

import (
    "strconv"
    "strings"
)

type JDTraceEntry struct {
    jx uint64
    jy uint64
    jz uint64
    wx uint64
    wy uint64
    wz uint64
}

type ArithmeticTraceEntry struct {
    pc uint64
    instructions uint8
    dependency_wait []uint64
}

type StoreTraceEntry struct {
    bytes uint8
    address uint64
    trace_entry ArithmeticTraceEntry
}

type BarrierTraceEntry struct {
    trace_entry ArithmeticTraceEntry
}

type LoadTraceEntry struct {
    bytes uint8
    address uint64
    trace_entry ArithmeticTraceEntry
}

func NewJDTraceEntry(entry string) JDTraceEntry {
    e := strings.Split(entry, " ")
    jx, _ := strconv.ParseUint(e[1], 16, 64)
    jy, _ := strconv.ParseUint(e[2], 16, 64)
    jz, _ := strconv.ParseUint(e[3], 16, 64)
    wx, _ := strconv.ParseUint(e[4], 16, 64)
    wy, _ := strconv.ParseUint(e[5], 16, 64)
    wz, _ := strconv.ParseUint(e[6], 16, 64)
    return JDTraceEntry{jx, jy, jz, wx, wy, wz}
}

func NewLoadTraceEntry(e string) LoadTraceEntry{
    entry := []rune(e)
    bytes, _:= strconv.ParseUint(string(entry[2]), 16, 8)
    address, _ := strconv.ParseUint(string(entry[4:16]), 16, 64)
    trace_entry := NewArithmeticTraceEntry(entry[17:])
    return LoadTraceEntry{uint8(bytes), address, trace_entry}
}

func NewStoreTraceEntry(e string) StoreTraceEntry {
    entry := []rune(e)
    bytes, _:= strconv.ParseUint(string(entry[2]), 16, 8)
    address, _ := strconv.ParseUint(string(entry[4:16]), 16, 64)
    trace_entry := NewArithmeticTraceEntry(entry[17:])
    return StoreTraceEntry{uint8(bytes), address, trace_entry}
}

func NewBarrierTraceEntry(e string) BarrierTraceEntry{
    entry := []rune(e)
    trace_entry := NewArithmeticTraceEntry(entry[2:])
    return BarrierTraceEntry{trace_entry}
}

func NewArithmeticTraceEntry(entry []rune) ArithmeticTraceEntry {
    pc, _ := strconv.ParseUint(string(entry[:12]), 16, 64)
    instructions, _ := strconv.ParseUint(string(entry[13]), 16, 8)
    deps, _ := strconv.ParseUint(string(entry[15]), 16, 64)

    dependency_wait := make([]uint64,deps)

    if deps > 0 {
        for i := uint64(0); i < deps; i++ {
            dependency_wait[i], _ = strconv.ParseUint(string(entry[17+i*13:29+i*13]), 16, 64)
        }
    }

    return ArithmeticTraceEntry{pc, uint8(instructions), dependency_wait}
}
