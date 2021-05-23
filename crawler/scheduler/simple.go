package scheduler

import (
	"learn-golang/crawler/types"
)

// 所有Worker共用一个输入
type SimpleScheduler struct {
	workerChan chan types.Request
}

func (s *SimpleScheduler) Submit(r types.Request) {
	// send types down to worker chan
	//避免出现循环等待的情况
	//in -> worker -> out -> 处理 -> in -> worker -> out......
	//当10个worker同时在处理10个request时, 第11个request在in中没有worker接它, 此时in堵塞了
	//当这10个worker中,有一个处理完成准备out的时候,发现in阻塞了, 它也out不了了,
	go func() { s.workerChan <- r }()
}

func (s *SimpleScheduler) WorkerChan() chan types.Request {
	return s.workerChan
}

func (s *SimpleScheduler) WorkerReady(chan types.Request) {
}

func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan types.Request)
}
