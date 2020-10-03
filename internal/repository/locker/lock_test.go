package locker_test

import (
	"context"
	"github.com/gbrlmza/lana-bechallenge-checkout/internal/repository/locker"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_locker_Lock_Success(t *testing.T) {
	// Given
	ctx := context.Background()
	l := locker.NewLocker(ctx)
	key := "my-lock-key"

	// When
	err := l.Lock(ctx, key)            // Lock
	errAlreadyLock := l.Lock(ctx, key) // Already Lock

	// Then
	assert.Nil(t, err)
	assert.EqualError(t, errAlreadyLock, "the resource 'my-lock-key' is locked")
}

func Test_locker_Lock_Success_Expired(t *testing.T) {
	// Given
	ctx := context.Background()
	l := locker.NewLocker(ctx)
	key := "my-lock-key"

	// When
	err := l.Lock(ctx, key)        // Lock
	time.Sleep(time.Second * 6)    // Let lock expire (5s default TTL)
	errNewLock := l.Lock(ctx, key) // Lock again

	// Then
	assert.Nil(t, err)
	assert.Nil(t, errNewLock)
}

func Test_locker_Unlock_Success(t *testing.T) {
	// Given
	ctx := context.Background()
	l := locker.NewLocker(ctx)
	key := "my-lock-key"

	// When
	err := l.Lock(ctx, key)         // Lock
	errUnlock := l.Unlock(ctx, key) // Unlock
	errNewLock := l.Lock(ctx, key)  // Lock again

	// Then
	assert.Nil(t, err)
	assert.Nil(t, errUnlock)
	assert.Nil(t, errNewLock)
}
