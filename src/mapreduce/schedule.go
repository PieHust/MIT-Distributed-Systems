package mapreduce

import (
	"fmt"
	"sync"
)

//
// schedule() starts and waits for all tasks in the given phase (mapPhase
// or reducePhase). the mapFiles argument holds the names of the files that
// are the inputs to the map phase, one per map task. nReduce is the
// number of reduce tasks. the registerChan argument yields a stream
// of registered workers; each item is the worker's RPC address,
// suitable for passing to call(). registerChan will yield all
// existing registered workers (if any) and new ones as they register.
//
func schedule(jobName string, mapFiles []string, nReduce int, phase jobPhase, registerChan chan string) {
	var ntasks int
	var n_other int // number of inputs (for reduce) or outputs (for map)
	switch phase {
	case mapPhase:
		ntasks = len(mapFiles)
		n_other = nReduce
	case reducePhase:
		ntasks = nReduce
		n_other = len(mapFiles)
	}

	fmt.Printf("Schedule: %v %v tasks (%d I/Os)\n", ntasks, phase, n_other)

	// All ntasks tasks have to be scheduled on workers. Once all tasks
	// have completed successfully, schedule() should return.
	//
	// Your code here (Part III, Part IV).
	var doTaskArgs = DoTaskArgs{JobName: jobName, Phase: phase, NumOtherPhase: n_other}

	var wg sync.WaitGroup
	wg.Add(ntasks)
	for i := 0; i < ntasks; i++ {
		if mapPhase == mapPhase {
			doTaskArgs.File = mapFiles[i]
		}
		doTaskArgs.TaskNumber = i
		doTaskArgs := doTaskArgs
		fmt.Println("------------" + phase + "--------------------")
		workerAddress := <-registerChan
		fmt.Printf("我得到地址了%v %s", i, workerAddress)
		var exec func(string, DoTaskArgs)
		exec = func(workerAddress string, doTaskArgs DoTaskArgs) {
			result := call(workerAddress, "Worker.DoTask", doTaskArgs, nil)
			if result == false {
				wg.Add(1)
				go exec(<-registerChan, doTaskArgs)
			}
			wg.Done()
			registerChan <- workerAddress
		}
		go exec(workerAddress, doTaskArgs)
	}
	wg.Wait()
	fmt.Printf("Schedule: %v done\n", phase)
}
