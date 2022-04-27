// Написать программу, которая при получении в канал сигнала SIGTERM останавливается не позднее, чем за одну секунду (установить таймаут).

package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, cancelFunction := context.WithCancel(context.Background())
	go func() {
		osSig := make(chan os.Signal)
		signal.Notify(osSig, syscall.SIGTERM)

		go func() {
			<-osSig
			cancelFunction()
		}()
	}()

	longTimeFunction(ctx)
}

func longTimeFunction(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("get TERM dignal")
			time.Sleep(time.Second * 5)
			fmt.Println("bye bye...")
			return
		case <-time.After(time.Second):
			fmt.Println("ZzZ Zzz...")
		}
	}
}
