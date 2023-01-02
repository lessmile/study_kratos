package server

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport"
	"time"

	s "helloworld/internal/service"
)

type Util struct {
	logger     log.Logger
	ctx        context.Context
	cancelFunc context.CancelFunc
	s          *s.DemoService
}

func (u *Util) Start(ctx context.Context) error {

	u.ctx, u.cancelFunc = context.WithCancel(ctx)
	after := time.After(time.Second * 5)

	go func() {

		for {
			u.s.ListDemo(ctx, nil)

			time.Sleep(time.Second)
		}
	}()

	for {
		select {
		case <-u.ctx.Done():
			log.Fatal("done ok")
			break
		case <-after:
			u.cancelFunc()
			//u.logger.Log(log.LevelInfo, "running...", time.Now())
		}
	}
}

func (u *Util) Stop(ctx context.Context) error {

	u.cancelFunc()

	u.logger.Log(log.LevelError, "stop utils_server")

	return nil
}

func NewUtilServer(demo *s.DemoService, logger log.Logger) transport.Server {

	return &Util{
		logger,
		nil,
		nil,
		demo,
	}
}
