package lock

import (
	"context"
	"sync"
)

type Unlock func(ctx context.Context)

type Accounts struct {
	Read  []string
	Write []string
}

type Locker interface {
	Lock(ctx context.Context, accounts Accounts) (Unlock, error)
}
type LockerFn func(ctx context.Context, accounts Accounts) (Unlock, error)

func (fn LockerFn) Lock(ctx context.Context, accounts Accounts) (Unlock, error) {
	return fn(ctx, accounts)
}

var NoOpLocker = LockerFn(func(ctx context.Context, accounts Accounts) (Unlock, error) {
	return func(ctx context.Context) {}, nil
})

type defaultLocker struct {
	globalLock sync.RWMutex
	locks      map[string]*sync.Mutex
	ledger     string
}

func (d *defaultLocker) Lock(ctx context.Context, accounts Accounts) (Unlock, error) {
	d.globalLock.RLock()
	lock, ok := d.locks[d.ledger]
	d.globalLock.RUnlock()
	if !ok {
		d.globalLock.Lock()
		lock, ok = d.locks[d.ledger] // Double check, the lock can have been acquired by another go routing between RUnlock and Lock
		if !ok {
			lock = &sync.Mutex{}
			d.locks[d.ledger] = lock
		}
		d.globalLock.Unlock()
	}

	unlocked := false
	lock.Lock()
	return func(ctx context.Context) {
		if unlocked {
			return
		}
		lock.Unlock()
		unlocked = true
	}, nil
}

func NewDefaultLocker(ledger string) *defaultLocker {
	return &defaultLocker{
		locks:  map[string]*sync.Mutex{},
		ledger: ledger,
	}
}
