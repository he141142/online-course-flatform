package adapters

import "context"

type WaitFunc func(ctx context.Context) error
type CleanUpTask func(ctx context.Context)

type Waiter interface {
	WaitFor(waits ...WaitFunc)
	AddCleanUpTask(cleanTasks ...CleanUpTask)
	Context() context.Context
	CancelFunc() context.CancelFunc
	Wait() error
}