package garbage

import (
	"fmt"
	"runtime"
	"time"
)

func printStat(mem runtime.MemStats) {
	runtime.ReadMemStats(&mem)
	fmt.Println("mem.Allocation:", mem.Alloc)
	fmt.Println("mem.TotalAlloc:", mem.TotalAlloc)
	fmt.Println("mem.HeapAlloca:", mem.HeapAlloc)
	fmt.Println("mem.numGC", mem.NumGC)
	fmt.Println(mem)
}

func GColl() {
	var mem runtime.MemStats
	printStat(mem)

	for i := 0; i < 10; i++ {
		s := make([]byte, 100000000)
		if s == nil {
			fmt.Println("Operation fail")
		}
		time.Sleep(5 * time.Second)
	}
	printStat(mem)
}
