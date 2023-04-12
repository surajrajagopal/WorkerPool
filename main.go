package main

import (
	"fmt"
	"sync"
)

var jobs = make(chan int, 5)
var result = make(chan int, 5)

func ProducerPushToQueue(noOfJobs int, jobs chan int) {
	for i := 1; i <= noOfJobs; i++ {
		jobs <- i
	}
	close(jobs)
}
func worker(wg *sync.WaitGroup) {
	sum := 0
	for job := range jobs {
		sum += job
		result <- sum
	}
	wg.Done()
}
func createWorkerPool(noOfWorker int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorker; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(result)
}

func ConsumerGetResults(done chan bool) {
	for v := range result {
		fmt.Println("sum of jobs = ", v)
	}
	done <- true
}

func main() {
	noOfWorker := 5
	noOfJobs := 5
	done := make(chan bool)

	go ProducerPushToQueue(noOfJobs, jobs)

	createWorkerPool(noOfWorker)

	go ConsumerGetResults(done)
	<-done
}
