package ping

import (
	"context"
	"log"
	"sync"
)

type Pinger struct {
	ctx     context.Context
	Channel *chan string
	syncWG  *sync.WaitGroup
}

func NewPinger(ctx context.Context) Pinger {
	channel := make(chan string)
	return Pinger{
		ctx:     ctx,
		Channel: &channel,
		syncWG:  &sync.WaitGroup{},
	}
}

func main() {
	p := NewPinger(context.Background())
	p.Ping()
	p.Ping()
	p.Ping()
	p.syncWG.Wait()

}

func (p *Pinger) Start() {
	go func() {
		for {
			select {
			case _ = <-p.ctx.Done():
				return
			case msg := <-*p.Channel:
				log.Default().Print("received msg: ", msg)
				p.syncWG.Done()
			}
		}
	}()
	p.Ping()

}

func (p *Pinger) Ping() {
	p.syncWG.Add(1)
	*p.Channel <- "ping"

}
