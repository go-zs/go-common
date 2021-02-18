package context

import (
	"context"
	"time"
)

func CopyContext(ctx context.Context) context.Context {
	return valueContext{ctx}
}

type valueContext struct{ context.Context }

func (valueContext) Deadline() (deadline time.Time, ok bool) { return }
func (valueContext) Done() <-chan struct{}                   { return nil }
func (valueContext) Err() error                              { return nil }
