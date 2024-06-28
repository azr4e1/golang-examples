package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func doSomething(ctx context.Context) {
	fmt.Printf("doSomething: myKey's value is %s\n", ctx.Value("myKey"))

	anotherCtx := context.WithValue(ctx, "myKey", "anotherValue")
	doAnother(anotherCtx)

	fmt.Printf("doSomething: myKey's value is %s\n", ctx.Value("myKey"))
}

func doAnother(ctx context.Context) {
	fmt.Printf("doAnother: myKey's value is %s\n", ctx.Value("myKey"))
}

func doSomethingLong(ctx context.Context) {
	for i := 1; ; {
		select {
		// The context.Context type provides a method called Done that can be checked to see whether a context has ended or not. This method returns a channel that is closed when the context is done, and any functions watching for it to be closed will know they should consider their execution context completed and should stop any processing related to the context
		case <-ctx.Done():
			if err := ctx.Err(); err != nil {
				fmt.Printf("doSomethingLong err: %s\n", err)
			}
			fmt.Printf("doSomethingLong: interrupted\n")
			fmt.Printf("status: %d", i)
			return
		default:
			i = (i * 2) % 1000
		}
	}
}

func main() {

	// You can use this as a placeholder when you’re not sure which context to use.
	// ctx := context.TODO()

	//  it’s designed to be used where you intend to start a known context.
	ctx := context.Background()

	// add data
	// values stored in a context are immutable. It wraps your parent context inside another one with the new value
	ctx = context.WithValue(ctx, "myKey", "myValue")

	// The context.WithCancel function and the cancel function it returns are most useful when you want to control exactly when a context ends
	ctxC, cancelCtxC := context.WithCancel(ctx)

	// handle interrupt signal
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGTERM, os.Interrupt)
	go func() {
		<-c
		cancelCtxC()
	}()

	doSomething(ctx)

	fmt.Println("Press Ctrl+C to stop")
	doSomethingLong(ctxC)

	time.Sleep(1 * time.Second)
	ctxD, cancelCtxD := context.WithDeadline(ctx, time.Now().Add(5*time.Second))
	// When a context is canceled from a deadline, the cancel function is still required to be called in order to clean up any resources that were used
	defer cancelCtxD()

	fmt.Println("\nNow some contexts with deadlines")
	doSomethingLong(ctxD)

	time.Sleep(1 * time.Second)
	ctxT, cancelCtxT := context.WithTimeout(ctx, 5*time.Second)
	// When a context is canceled from a deadline, the cancel function is still required to be called in order to clean up any resources that were used
	defer cancelCtxT()

	fmt.Println("\nNow some contexts with timeout")
	doSomethingLong(ctxT)
}
