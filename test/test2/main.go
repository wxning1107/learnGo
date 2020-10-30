package main

import "fmt"

func main() {
	// Creates a circuit breaker that will trip if the function fails 10 times
	//cb := circuit.NewThresholdBreaker(10)
	//
	//events := cb.Subscribe()
	//go func() {
	//	for {
	//		e := <-events
	//		fmt.Println(e)
	//		// Monitor breaker events like BreakerTripped, BreakerReset, BreakerFail, BreakerReady
	//	}
	//}()
	//
	//cb.Call(func() error {
	//	// This is where you'll do some remote call
	//	// If it fails, return an error
	//	res, err := http.Get("abc")
	//	err = errors.New("abc")
	//	fmt.Printf("resp: %s\n", res)
	//	return err
	//
	//}, 0)
	//
	//time.Sleep(time.Second * 3)
	src := "a"
	TestString(&src)
	fmt.Println(src)

	v := 12

	fmt.Println("src:", v)

	change(v)

	fmt.Println("after change:", v)

	p := 12

	changep(&p)

	fmt.Println("after changep:", p)
}

func TestString(s *string) {
	a := "abc"
	*s = a
}

func change(v int) {
	v = 33
}

func changep(p *int) {
	*p = 55
}
