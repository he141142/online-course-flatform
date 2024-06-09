package adapters

import "context"

type WaitFunc func(ctx context.Context) error
type CleanUpTask func(ctx context.Context)

type Waiter interface {
	Wait(waits ...WaitFunc)
	AddCleanUpTask(cleanTasks ...CleanUpTask)
	Context() context.Context
	CancalFunc() context.CancelFunc
}