package gotrylock

import (
	"errors"
	"runtime"
	"sync/atomic"
	"time"
)

var (
	ErrUnlockOfUnlockedMutex = errors.New("unlock of unlocked mutex")
)

type TryMutex struct {
	flag uint32
}

// Lock locks TryMutex as usual Mutex
func (tryMutex *TryMutex) Lock() {
	for !atomic.CompareAndSwapUint32(&tryMutex.flag, 0, 1) {
		runtime.Gosched()
	}
}

// Unlock unlocks TryMutex as usual Mutex.
// Panics with ErrUnlockOfUnlockedMutex, if TryMutex already unlocked
func (tryMutex *TryMutex) Unlock() {
	if !atomic.CompareAndSwapUint32(&tryMutex.flag, 1, 0) {
		panic(ErrUnlockOfUnlockedMutex)
	}
}

func (tryMutex *TryMutex) Locked() bool {
	return tryMutex.flag == 1
}

// TryLock tryes to lock mutex to at least duration in timeout parameter.
// Returns true on success.
func (tryMutex *TryMutex) TryLock(timeout time.Duration) (ok bool) {
	timeoutFlag := uint32(0)
	timeoutFlagPointer := &timeoutFlag
	time.AfterFunc(timeout, func() {
		atomic.CompareAndSwapUint32(timeoutFlagPointer, 0, 1)
	})
	for {
		if atomic.CompareAndSwapUint32(timeoutFlagPointer, 1, 0) {
			return false
		} else if atomic.CompareAndSwapUint32(&tryMutex.flag, 0, 1) {
			return true
		}
		runtime.Gosched()
	}
}
