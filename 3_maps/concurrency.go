package main

import (
	"math/rand"
	"fmt"
	"sync"
	"time"
)

func printHello() {
	fmt.Println("Hello from goroutine")
}

func goroutineExample() {
	go printHello()
	time.Sleep(1 * time.Second) // pause to let child goroutine execute
	fmt.Println("Hello from main")
}

func chanExample() {
	c := make(chan int, 1)
	fmt.Println(len(c), cap(c))
	c <- 3
	fmt.Println(len(c), cap(c))
	a, f := <-c
	fmt.Println(a, f)
	fmt.Println(len(c), cap(c))
	// a, f = <-c // fatal error: all goroutines are asleep - deadlock!
	// fmt.Println(a, f)
	// fmt.Println(len(c), cap(c))
}

func myFunc(c chan int) {
	c <- cap(c)
}

func chanCapacity() {
	c := make(chan int, 5)
	for i := 0; i < 5; i++ {
		go myFunc(c)
	}
	fmt.Print(<-c)
}

func task(c chan int, n int) {
	c <- n + 1
}

func task2(c chan string, s string) {
	for i := 0; i < 5; i++ {
		c <- s + " "
	}
	close(c)
}

func testChan() {
	c := make(chan int)
	n := 7
	go task(c, n)
	time.Sleep(1 * time.Second)
	fmt.Print(<-c)
	fmt.Println()
	c2 := make(chan string, 5)
	go task2(c2, "Hello")
	for s := range c2 {
		fmt.Print(s)
	}
}

func removeDuplicates(c1 chan string, c2 chan string) {
	defer close(c2)
	var previous string
	for {  // could also better be 'for range'
		s, ok := <- c1
		if !ok {
			return
		}
		if s != previous {
			c2<-s
			previous = s
		}
	}
}

func testRemoveDuplicates() {
	c1 := make(chan string)
	c2 := make(chan string)
	go removeDuplicates(c1, c2)

	go func() {
		defer close(c1)
		for _, item := range "abcdefghijklmnopqrstuvwxy" {
			c1<-string(item)
		}
	}()

	for res := range c2 {
		fmt.Println(res)
	}
}

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c<-x: // until c is ready to read, write to it
			x, y = y, y+x
		case <-quit: // when anything comes to 'quit' - quit
			fmt.Println("quit")	
			return
		}
	}
}

func testFibonacci(lim int) {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < lim; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)  // wouldn't work properly with 'go fibonacci'. Why ?
	// because it wouldn't wait until it's done in this case :)) needs sync - something like reading from signal chan
}

func myFuncImp(done chan struct{}) {
	fmt.Println("Hello from Goroutine")
	close(done)  // by closing the chan it sends signal that it's finished
}

func improvedGoroutineExample() {
	done := make(chan struct{}) // type isn't important - just struct doesn't use memory
	go myFuncImp(done)
	<-done  // use chan for sinchronization. when it's closed, Go understands that nothing will be written to it and continues
}

func myFuncUlt() <-chan struct{} {  // returns read-only chan
	done := make(chan struct{})
	go func() {
		fmt.Println("Hello from hellknowwhere")
		close(done)
	}()
	return done  // returns signal chan to main func
}

func ultimateGoroutineExample() {
	<-myFuncUlt() // await reading from chan returned by myFuncUlt
}

// func checkCores() {
// 	fmt.Println(runtime.NumCPU())
// }

func worker(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Goroutine %d started work \n", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d finished work \n", i)
}

func ExampleWaitGroup() {
	wg := new(sync.WaitGroup)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go worker(i, wg)
	}

	wg.Wait() // blocks until wg.counter is zero
	fmt.Println("Goroutines finished their work")
}

func WaitGroupTask() {
	worker2 := func() {
		time.Sleep(10 * time.Millisecond)
		fmt.Println("Done")
	}
	wg := new(sync.WaitGroup)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker2()
			fmt.Println("DoneDone")
		}()
	}
	wg.Wait()
	fmt.Println("Finished")
}

func WhyDoWeNeedMutexFor() {
	var x int
	wg := new(sync.WaitGroup)

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			x++
		}(wg)
	}
	wg.Wait()
	fmt.Println(x)  // probably won't be 1000
}

func MutexExample() {
	var x int
	wg := new(sync.WaitGroup)
	mu := new(sync.Mutex)

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, mu *sync.Mutex) {
			defer wg.Done()
			mu.Lock()
			x++
			mu.Unlock()
		}(wg, mu)
	}
	wg.Wait()
	fmt.Println(x)
}

func SimpleTicker() {
	ticker := time.Tick(time.Second)  // make a chan which 'ticks' periodically
	for i := 0; i < 5; i++ {  // 'for range ticker' is also possible
		<-ticker
		fmt.Println("one more")
	}
}

func tinkWork() <-chan struct{} {
	done := make(chan struct{})  // chan for synchronization

	go func() {
		defer close(done)
		stop := time.NewTimer(time.Second)
		tick := time.NewTicker(time.Millisecond * 200)
		defer tick.Stop()
		
		for {
			select {
			case <-stop.C:  // if something came from stop chan
				return
			case <-tick.C:
				fmt.Println("tick-tock")  // each 0.2 sec
			}
		}
	}()
	return done  // return the channel by closing which we will release main goroutine
}

func SyncTickerExample() {
	<-tinkWork()
}

func AsTinkWork(i int, limit <-chan time.Time, wg *sync.WaitGroup) {
	defer wg.Done()
	<-limit
	fmt.Printf("worker %d done\n", i)
}

func AsyncTickerExample() {
	// might be used for balancing
	tick := time.NewTicker(time.Second)
	defer tick.Stop()
	wg := new(sync.WaitGroup)

	for i := 0; i <= 5; i++ {
		wg.Add(1)
		go AsTinkWork(i, tick.C, wg)
	}

	wg.Wait()
}

func sleepyGopher(id int, c chan int) {
	time.Sleep(time.Duration(rand.Intn(4000)) * time.Millisecond)
	c <- id
}

func AfterExample() {
	timeout := time.After(2 * time.Second)
	c := make(chan int)

	for j := 0; j < 5; j++ {
		go sleepyGopher(j, c)
	}

	for i := 0; i < 5; i++ {
		select {
		case gopherID := <-c:
			fmt.Println("gopher ", gopherID, " has finished sleeping")
		case <-timeout:
			fmt.Println("my patience ran out")
			return
		}
	}
	// may wait forever if no case appears
}

func strangeTest() {
	const N = 10
	m := make(map[int]int)
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()
			fmt.Printf("i = %d\n", i)
			m[i] = i
			mu.Unlock()
		}() // all of them would be rally run after i is set to 10, so they all use the same value
		// ('use goroutines on loop iterator variables' warning)
		// solution is to pass i to function call
	}
	wg.Wait()
	fmt.Println(len(m))
}

func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	done := make(chan int)
	var n int
	go func() {
		defer close(done)
		select {
		case n = <- firstChan:
			done <- n * n
		case n = <- secondChan:
			done <- n * 3
		case <- stopChan:
			return
		}
	}()

	return done
}

func testCalculator() {
	f := make(chan int)
	s := make(chan int)
	st := make(chan struct{})
	done := calculator(f, s, st)
	// go func() {
	// 	res := <-done
	// 	fmt.Printf("res: %d\n", res)
	// }()
	// f <- 3
	close(st)
	// <-done
	fmt.Println(<-done)
	fmt.Println("Finished ???")  // strange but yes :))
}

func summator(arguments <-chan int, done <-chan struct{}) <-chan int {
	res := make(chan int)
	
	go func() {
		defer close(res)
		sum := 0
		for {
			select {
			case n := <-arguments:
				sum += n
				fmt.Println(n, sum)
			case <-done:
				res <- sum
				return
			}
		}
	}()
	return res
}

func testSummator() {
	a := make(chan int)
	d := make(chan struct{})
	res := summator(a, d)
	for i := 1; i < 10; i++ {
		a <- i
	}
	close(d)
	fmt.Println(<-res)
}

func merge2Channels(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	// my own solution - terrible! :))
	type pair struct {
		a int
		b int
	}
	pairs := []pair{}
	results := make(map[int]int)
	go func() {
		wg := &sync.WaitGroup{}
		mu := &sync.Mutex{}
		defer close(out)  // delete for the task - would raise error
		for i := 0; i < n; i++ {
			x1 := <-in1
			x2 := <-in2
			pairs = append(pairs, pair{x1, x2})
			fmt.Println(pairs)
			fmt.Printf("inner func %d, %d working\n", x1, x2)
			wg.Add(2)
			getRes := func(x int, mu *sync.Mutex) {
				defer wg.Done()
				if _, ok := results[x]; ok {
					return
				}
				res := fn(x)
				fmt.Printf("inner res of %d is %d\n", x, res)
				mu.Lock()
				results[x] = res
				mu.Unlock()
			}
			go getRes(x1, mu)
			go getRes(x2, mu)
		}
		wg.Wait()
		for _, el := range pairs {
			res := results[el.a] + results[el.b]
			fmt.Printf("res: %d\n", res)
			out<-res
		}
	}()
}

func tFunc(n int) int {
	fmt.Printf("tFunc %d started\n", n)
	time.Sleep(time.Duration(rand.Intn(4000)) * time.Millisecond)
	return n * 2
}

func testMerge() {
	fn := tFunc
	in1 := make(chan int)
	in2 := make(chan int)
	out := make(chan int)
	const LIM = 5
	// merge2Channels(fn, in1, in2, out, LIM)
	merge2ChannelsV2(fn, in1, in2, out, LIM)
	merge2ChannelsV3(fn, in1, in2, out, LIM)
	for i := 0; i < LIM; i++ {
		in1<-i
		in2<-i*3
	}
	close(in1)
	close(in2)
	// for range out {
	for j := 0; j < LIM; j++ {
		fmt.Println(<-out)
	}
}

type MergeWorker struct {
	sync.Mutex
	sync.WaitGroup
}

func (w *MergeWorker) Do(work func(int) int, x int, res *int) {
	w.Add(1)
	go func() {
		y := work(x)
		w.Lock()
		*res += y
		w.Unlock()
		w.Done()
	}()
}

func merge2ChannelsV2(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	go func() {
		res := make([]int, n)
		var w MergeWorker
		for i := range res {
			w.Do(fn, <-in1, &res[i])
			w.Do(fn, <-in2, &res[i])
		}
		w.Wait()
		for i := range res {
			out <- res[i]
		}
	}()
}

func merge2ChannelsV3(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	// results in test are not consistent and always correct
	c1 := make(chan chan int, n)
	c2 := make(chan chan int, n)
	procIn := func(in <-chan int, c chan chan int) {
		for i:=0; i<n; i++ {
			fc := make(chan int)
			c<-fc  // pass middle channels to res chan in order so they would be readed in the same order!!!
			go func(resChan chan int, x int) {
				resChan<-fn(x)
			}(fc, <-in)
		}
	}
	go procIn(in1, c1)
	go procIn(in2, c2)

	go func() {
		for i:=0; i<n; i++ {
			out<- <-<-c1 + <-<-c2  // as soon as there be result in the chan next in queue, it would be added and sent to out chan
		}
	}()
}

func main() {
	// goroutineExample()
	// chanExample()
	// chanCapacity()
	// testChan()
	// testRemoveDuplicates()
	// testFibonacci(10)
	// improvedGoroutineExample()
	// ultimateGoroutineExample()
	// checkCores()
	// ExampleWaitGroup()
	// WaitGroupTask()
	// MutexExample()
	// WhyDoWeNeedMutexFor()
	// SimpleTicker()
	// SyncTickerExample()
	// AsyncTickerExample()
	// AfterExample()
	// strangeTest()
	// testCalculator()
	// testSummator()
	testMerge()
}
