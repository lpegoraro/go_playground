package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

//type MockAgent interface {
//	Start(ctx context.Context, cancelFunc context.CancelFunc)
//	Stop(ctx context.Context)
//}
//
//type agent struct {
//	cancelFunc context.CancelFunc
//}

// main Function of the application
func main() {
	// {},
	rootCtx, mainCtxCancelFunc := context.WithCancel(context.WithValue(context.Background(), " #routine", 1))
	subRoutine1Ctx := context.WithValue(rootCtx, "#routine", 10)
	subRoutine2Ctx, sr2CancelFunc := context.WithTimeout(context.WithValue(rootCtx, "#routine", 20), time.Second*5)
	subRoutine3Ctx := context.WithValue(rootCtx, "#routine", 30)
	go routine1(subRoutine1Ctx)
	go cancelableRoutine(subRoutine2Ctx, mainCtxCancelFunc)
	go routine1(subRoutine3Ctx)

	timeoutPrint := false
LOOP:
	for {
		select {
		case <-subRoutine2Ctx.Done():
			if !timeoutPrint {
				fmt.Printf("routine #%d timeout\n", subRoutine2Ctx.Value("#routine"))
				timeoutPrint = true
			}
			sr2CancelFunc()
		case <-rootCtx.Done():
			sr2CancelFunc()
			fmt.Println("main context done, exiting and checking for goroutines and contexts")
			fmt.Printf("#gouroutines:%d\n", runtime.NumGoroutine())
			fmt.Printf("#contexts:%d\n", rootCtx.Value("#routine")) //nil
			fmt.Printf("#contexts:%d\n", subRoutine1Ctx.Value("#routine"))
			fmt.Printf("#contexts:%d\n", subRoutine2Ctx.Value("#routine"))
			fmt.Printf("#contexts:%d\n", subRoutine3Ctx.Value("#routine"))
			break LOOP
		}
	}
}

func routine1(ctx context.Context) {
	i := 1
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("  #routine:%d exit\n", ctx.Value(" #routine"))
			fmt.Println(ctx.Err())
			return
		default:
			fmt.Printf("routine#%d: heartbeat #%d\n", ctx.Value("#routine"), i)
			time.Sleep(time.Second * 1)
			i++
		}
	}
}

func cancelableRoutine(ctx context.Context, mainCtxCancel context.CancelFunc) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("    #routine:%d exit\n", ctx.Value(" #routine"))
			fmt.Println(ctx.Err())
			return
		default:
			fmt.Printf("    #routine:%d Exit in 10 seconds\n", ctx.Value("#routine"))
			time.Sleep(time.Second * 10)
			mainCtxCancel()
		}
	}
}
