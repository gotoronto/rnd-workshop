package concurrency

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Create a new function named BufferedChannel. The BufferedChannel function returns a channel that has 1, 2, 3 stored in it.
// If your test hangs, you probably have forgotten closing the channel in the BufferedChannel function.
// Creating a channel: https://tour.golang.org/concurrency/3
// Closing a channel: https://tour.golang.org/concurrency/4
func TestBufferedChannel(t *testing.T) {
	expectedOutput := []int{1, 2, 3}

	channel := BufferedChannel()

	assert.Equal(t, "chan int", reflect.TypeOf(channel).String())
	index := 0
	for number := range channel {
		assert.Equal(t, expectedOutput[index], number)
		index++
	}
}

// Import the custom sleep library from the sleep folder in same directory as this test case. Create a new function called MySleep in concurrency.go and make MySleep sleeps for a second.
//
// To import the sleep library, add
// import "github.com/gotoronto/rnd-workshop/exercise_3_prep/sleep"
// to the beginning of concurrency.go below "package concurrency"
func TestMySleep(t *testing.T) {
	start := time.Now()
	MySleep()
	elapsed := time.Since(start)

	assert.True(t, elapsed >= time.Second)
}

// Create a new function called ManySleep. In ManySleep, you will create ONE sleep.Sleep object and the sleep.Sleep object has to sleep for a second for 1000 times.
// Imagine you have 1000 workers, rather than letting one worker sleep at a time, you have to let 1000 of them sleep at the same time.
// You will have to use goroutine to achieve this.
//
// When you make one worker sleeps:
// sleep := sleep.New(time.Second)
// go sleep.Sleep()
//
// When you make two workers sleep:
// sleep := sleep.New(time.Second)
// go sleep.Sleep()
// go sleep.Sleep()
//
// You get the drift. :p
//
// Oh by the way, the function has to return the sleep.Sleep object created.
// In addition, the function has to complete in less than 2 seconds.
//
// Is your program exiting early?
// You probably have to wait till all the workers have finished sleeping before exiting the function.
func TestManySleep(t *testing.T) {
	start := time.Now()
	sleep := ManySleep()
	elapsed := time.Since(start)

	assert.True(t, elapsed < 2*time.Second)
	assert.Equal(t, "*sleep.Sleep", reflect.TypeOf(sleep).String())
	assert.Equal(t, 1000, sleep.Count())
}

// If you have used a timer to make sure all the workers have finished sleeping early, you are not using the most efficient way of checking. Don't worry, this test will teach you a better way of doing it.
//
// Implement a function called BetterManySleep. Like ManySleep, you will create ONE sleep.Sleep object and the sleep.Sleep object has to sleep for a second for 1000 times. The function must use the sleep.Done() to count how many workers have finished sleeping before exiting.
// This tutorial may help you: https://gobyexample.com/channel-synchronization
func TestBetterManySleep(t *testing.T) {
	start := time.Now()
	sleep := BetterManySleep()
	elapsed := time.Since(start)

	assert.True(t, elapsed < 1010*time.Millisecond)
	assert.Equal(t, "*sleep.Sleep", reflect.TypeOf(sleep).String())
	assert.Equal(t, 1000, sleep.Count())
}
