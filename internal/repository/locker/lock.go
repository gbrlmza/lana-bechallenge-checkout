package locker

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// NOTE: This is a basic in memory lock service. When working with multiple instances this
// type of lock doesn't work because a locked resource in one instance is no locked in the
// other instances and with concurrency we can have more than one instance working with
// the same resource. An external distributed lock shared by all instances is the proper
// way to do it, using something like Redis, Zookeeper, DynamoDB, etc...

const (
	defaultTTL       = 5 * time.Second
	maxRetryAttempts = 3
	retryWaitTime    = time.Millisecond * 100
)

func NewLocker(ctx context.Context) *locker {
	return &locker{
		lockMap: make(map[string]lockValue, 0),
	}
}

type locker struct {
	lockMap map[string]lockValue
	mutex   sync.Mutex
}

type lockValue struct {
	Key       string
	TTL       time.Duration
	CreatedAt time.Time
}

func (l lockValue) expired() bool {
	return l.CreatedAt.Add(l.TTL).Before(time.Now())
}

func (l locker) Lock(ctx context.Context, resource string) error {
	var err error

	// Retry strategy
	for i := 1; i <= maxRetryAttempts; i++ {
		if err = l.doLock(ctx, resource); err == nil {
			// Resource successfully locked
			return nil
		}
		time.Sleep(retryWaitTime)
	}

	return err
}

func (l locker) doLock(ctx context.Context, resource string) error {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	value, ok := l.lockMap[resource]
	if ok && !value.expired() {
		return fmt.Errorf("the resource '%s' is locked", resource)
	}

	l.lockMap[resource] = lockValue{
		Key:       resource,
		TTL:       defaultTTL,
		CreatedAt: time.Now(),
	}

	return nil
}

func (l locker) Unlock(ctx context.Context, resource string) error {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	delete(l.lockMap, resource)
	return nil
}
