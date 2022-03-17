package zispc

import "sync"

var err error

var errorLock sync.RWMutex

func Error() error {
	errorLock.RLock()
	defer errorLock.RUnlock()

	return err
}

func setError(e error) {
	errorLock.Lock()
	defer errorLock.Unlock()

	err = e
}
