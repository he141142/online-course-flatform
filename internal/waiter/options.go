package waiter

import "context"

type ConfigOpt interface {
	Apply(*config)
}

type catchSignalCfg bool

func (cfg catchSignalCfg) Apply(wcf *config) {
	wcf.catchSignal = bool(cfg)
}

type parentCtxOpt struct {
	context.Context
}

func (cfg parentCtxOpt) Apply(wcf *config) {
	wcf.parentCtx = context.Context(cfg)
}
