package waiter

import "context"

type ConfigOpt interface {
	Apply(*config)
}

type CatchSignalCfg bool

func (cfg CatchSignalCfg) Apply(wcf *config) {
	wcf.catchSignal = bool(cfg)
}

type ParentCtxOpt struct {
	context.Context
}

func (cfg ParentCtxOpt) Apply(wcf *config) {
	wcf.parentCtx = context.Context(cfg)
}
