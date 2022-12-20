package main

import (
	"context"
	"playground/ping"
)

func main() {
	ctx := context.Background()
	pinger := ping.NewPinger(ctx)
	pinger.Start()
	pinger.Ping()
	*pinger.Channel <- "pong"
	ctx.Done()
}
