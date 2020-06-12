// mkslice test out of memory
package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"strconv"
	"strings"
	"time"
)

var limit int = 3 //3 = pass!, 5 = fatal error!

var taskmsg chan string //任务消息

func main() {
	fmt.Println("make slice ", limit, "*500MB! test GO memory allocation")
	var n, i int
	n = 525019669
	taskmsg = make(chan string)
	for i = 1; i <= limit; i++ {
		go makeslice(i, n)
	}
	var goresult string
	for i = 1; i <= limit; i++ {
		goresult = <-taskmsg
		fmt.Println("make slice Done ", goresult)
	}
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("GC前内存占用： %d Kb  CPU核心: %d\n", m.Alloc/1024, runtime.NumCPU())
	runtime.GC()
	debug.FreeOSMemory()

	runtime.ReadMemStats(&m)
	fmt.Printf("结束内存占用： %d Kb  CPU核心: %d\n", m.Alloc/1024, runtime.NumCPU())
}

func makeslice(index int, n int) int {
	var mr1, mr2 runtime.MemStats
	runtime.ReadMemStats(&mr1)
	itmq := time.Now()
	var i int
	a1 := make([]byte, n)
	for i = 0; i <= len(a1)-1; i++ {
		a1[i] = byte(index*3 + i + 1019)
	}
	fmt.Println("任务：", index, cap(a1), a1[0], a1[99999], a1[525019661])
	itme := time.Now()
	ms1 := (itme.UnixNano() - itmq.UnixNano()) / 1e6
	fmt.Printf("make 耗时：%v ms\n", NumberFormat(strconv.FormatInt(ms1, 10)))

	runtime.ReadMemStats(&mr2)
	fmt.Printf("内存占用： %d Kb  CPU核心: %d\n", (mr2.Alloc-mr1.Alloc)/1024, runtime.NumCPU())
	taskmsg <- "完成" + strconv.Itoa(index)
	return cap(a1)
}

//格式化数值    1,234,567,898.55
func NumberFormat(str string) string {
	length := len(str)
	if length < 4 {
		return str
	}
	arr := strings.Split(str, ".") //用小数点符号分割字符串,为数组接收
	length1 := len(arr[0])
	if length1 < 4 {
		return str
	}
	count := (length1 - 1) / 3
	for i := 0; i < count; i++ {
		arr[0] = arr[0][:length1-(i+1)*3] + "," + arr[0][length1-(i+1)*3:]
	}
	return strings.Join(arr, ".") //将一系列字符串连接为一个字符串，之间用sep来分隔。
}
