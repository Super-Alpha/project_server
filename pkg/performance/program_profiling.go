package performance

import (
	"bytes"
	"flag"
	"log"
	"math/rand"
	"os"
	"runtime"
	"runtime/pprof"
	"sync"
)

var cpuProfile = flag.String("cpuProfile", "cpu.pprof", "write cpu profile to `file`")
var memProfile = flag.String("memProfile", "mem.pprof", "write mem profile to `file`")

func ProgramProfiling() {
	flag.Parse()

	if *cpuProfile != "" {
		f, err := os.Create(*cpuProfile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}

		defer f.Close()

		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	var wg sync.WaitGroup
	wg.Add(200)
	for i := 0; i < 200; i++ {
		go cycleNum(30000, &wg)
	}

	writeBytes()

	wg.Wait()

	if *memProfile != "" {
		f, err := os.Create(*memProfile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer f.Close()
		runtime.GC()

		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("cound not write memory profile: ", err)
		}
	}
}

func cycleNum(num int, wg *sync.WaitGroup) {
	slice := make([]int, 0)
	for i := 0; i < num; i++ {
		for j := 0; j < num; j++ {
			j = i + j
			slice = append(slice, j)
		}
	}
	wg.Done()
}

func writeBytes() *bytes.Buffer {
	var buff bytes.Buffer
	for i := 0; i < 30000; i++ {
		buff.Write([]byte{'0' + byte(rand.Intn(10))})
	}
	return &buff
}
