package scheduler

import "learn-golang/crawler/types"

type QueuedSampleScheduler struct {
	requestChan chan types.Request
	workChan chan chan types.Request
}

func (q QueuedSampleScheduler) Submit(request types.Request) {
	q.requestChan <- request
}

func (q QueuedSampleScheduler) WorkerChan() chan types.Request {
	return make(chan types.Request)
}

func (q QueuedSampleScheduler) WorkerReady(requests chan types.Request) {
	q.workChan <- requests
}

func (q QueuedSampleScheduler) Run() {
	go func() {
		var requestChan []types.Request
		var workChan []chan types.Request
		for  {
			if len(requestChan) > 0 && len(workChan) > 0 {
				workChan[0] <- requestChan[0]
				workChan = workChan[1:]
				requestChan = requestChan[1:]
			}
			select {
				case r := <- q.requestChan:
					requestChan = append(requestChan, r)
				case w := <-q.workChan:
					workChan = append(workChan, w)
			}
		}
	}()
}

