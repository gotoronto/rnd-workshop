package concurrency

import (
	"time"

	"github.com/gotoronto/rnd-workshop/lib/sleep"
)

func BufferedChannel() chan int {
	channel := make(chan int, 3)
	channel <- 1
	channel <- 2
	channel <- 3
	close(channel)
	return channel
}

func MySleep() *sleep.Sleep {
	sleep := sleep.New(time.Second)
	sleep.Sleep()
	return sleep
}

func ManySleep() *sleep.Sleep {
	sleep := sleep.New(time.Second)
	for i := 0; i < 1000; i++ {
		go sleep.Sleep()
	}
	time.Sleep(1500 * time.Millisecond)
	return sleep
}

func BetterManySleep() *sleep.Sleep {
	sleep := sleep.New(time.Second)
	for i := 0; i < 1000; i++ {
		go sleep.Sleep()
	}

	count := 0
	for {
		select {
		case <-sleep.Done():
			count++
			if count >= 1000 {
				return sleep
			}
		}
	}

	return nil
}
