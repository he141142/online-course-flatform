package waiter

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"drake.elearn-platform.ru/internal/adapters"
)

type config struct {
	catchSignal bool
	parentCtx   context.Context
}

type eWaiter struct {
	cleanUpTasks []adapters.CleanUpTask
	waitFn       []adapters.WaitFunc
	cancelFn     context.CancelFunc
	ctx          context.Context
}

func NewWaiter(configOpts ...ConfigOpt) adapters.Waiter {
	cfg := &config{
		catchSignal: false,
		parentCtx:   context.Background(),
	}

	for _, opt := range configOpts {
		opt.Apply(cfg)
	}
	eWaiter := &eWaiter{}
	eWaiter.ctx = cfg.parentCtx
	eWaiter.ctx, eWaiter.cancelFn = context.WithCancel(eWaiter.ctx)
	if cfg.catchSignal {
		eWaiter.ctx, eWaiter.cancelFn = signal.NotifyContext(eWaiter.ctx, os.Interrupt,
			syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	}

	return eWaiter
}

func (w *eWaiter) Wait(waits ...adapters.WaitFunc) {
	w.waitFn = append(w.waitFn, waits...)
}

func (w *eWaiter) AddCleanUpTask(cleanTasks ...adapters.CleanUpTask) {
	w.cleanUpTasks = append(w.cleanUpTasks, cleanTasks...)
}

func (w *eWaiter) Context() context.Context {
	return w.ctx
}

func (w *eWaiter) CancalFunc() context.CancelFunc {
	return w.cancelFn
}