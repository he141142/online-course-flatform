package waiter

import "context"

type WaitFunc func(ctx context.Context) error
type CleanUpTask func(ctx context.Context)

type config struct {
	catchSignal bool
	parentCtx   context.Context
}

type waiter struct {
	cleanUpTasks []CleanUpTask
	waitFn       []WaitFunc
}