# makeslice

make slice 5*500MB! test GO memory allocation

test go make slice out of memory
runtime: out of memory: cannot allocate *******-byte block (****** in use)

```go
`Thanks a lot for davecheney's advice:`

`the upper limit of a slice is some significant fraction of the address space of a process.` 
`For 32 bit processes, between 1-2 gb, perhaps a little less on 32 bit windows because of DLL address space fragmentation.` 
`For 64 bit processes, in excess of a terrabyte, 10^40 bits.`

`Thank you for posting your sample code. Unlike many projects on GitHub, the Go project does not use its bug tracker for general discussion or asking questions. We only use our bug tracker for tracking bugs and tracking proposals going through the Proposal Process.`
```

---------------------------

make slice  3 *500MB! test GO memory allocation

任务： 3 525019669 4 163 17

make 耗时：928 ms

任务： 1 525019669 254 157 11

make 耗时：936 ms

内存占用： 1025438 Kb  CPU核心: 8

make slice Done  完成1

内存占用： 1538158 Kb  CPU核心: 8

make slice Done  完成3

任务： 2 525019669 1 160 14

make 耗时：808 ms

内存占用： 512719 Kb  CPU核心: 8

make slice Done  完成2

GC前内存占用： 1538243 Kb  CPU核心: 8

结束内存占用： 80 Kb  CPU核心: 8
