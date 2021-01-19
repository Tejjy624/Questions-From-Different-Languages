package bug2

import (
	"sync"
)

func bug2(n int, foo func(int) int, out chan int) {
	var wg sync.WaitGroup
	channel := make(chan int, n) //create channel with type 'chan int'
	for i := 0; i < n; i++ {
		channel <- i
		wg.Add(1)
		go func() {
			channelvar := <-channel //recieve expression in assignment
			out <- foo(channelvar) //send the statement
			wg.Done()
		}()
	}
	wg.Wait()
	close(out)
}
