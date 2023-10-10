package main

import (
	"fmt"
	"strconv"
	"time"
)

const (
	NucExport string = "NucExport"
)

var chanMap map[string]chan int

func init() {
	chanMap = make(map[string]chan int)
	chanMap[NucExport] = make(chan int, 1)
	chanMap[NucExport] <- 1
}

func main() {
	go QueueDoSome("1", "AAAA", NucExport)
	go QueueDoSome("2", "BBBB", NucExport)
	go QueueDoSome("3", "CCCC", NucExport)
	go QueueDoSome("4", "DDDD", NucExport)
	go QueueDoSome("5", "EEEE", NucExport)
	go QueueDoSome("6", "FFFF", NucExport)

	time.Sleep(100 * time.Second)
}

// BackendExec
//
//	@param taskID
//	@param process
//	@param input
//	@param execFunc
func DoSome(taskID, process string) {
	fmt.Println("DoSome:", process, "TaskID:", taskID)
	value, _ := strconv.Atoi(taskID)
	time.Sleep(time.Duration(value) * time.Second)
}

// QueueDoSome
//
//	@param taskID
//	@param process
//	@param queueKey
func QueueDoSome(taskID, process string, queueKey string) {

	ch, ok := chanMap[queueKey]
	if !ok {
		DoSome(taskID, process)
		return
	}
	<-ch
	DoSome(taskID, process)
	ch <- 1
}
