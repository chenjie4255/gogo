package gogo

import (
	"errors"
	"sync"
)

type Fn func() error
type Fns []Fn

func Run(fnQueue ...Fns) error {
	errChan := make(chan error)
	isExit := false
	var mux sync.RWMutex
	defer func() {
		mux.Lock()
		defer mux.Unlock()
		isExit = true
		close(errChan)
	}()

	for i := 0; i < len(fnQueue); i++ {
		fns := fnQueue[i]
		for j := 0; j < len(fns); j++ {
			fn := fns[j]
			go func() {
				defer func() {
					e := recover()
					if e != nil {
						mux.RLock()
						defer mux.RUnlock()
						if !isExit {
							errChan <- errors.New("panic on func...")
						}
					}
				}()

				err := fn()
				mux.RLock()
				defer mux.RUnlock()
				if !isExit {
					errChan <- err
				}
			}()
		}

		// for again for reading return value of functions
		for j := 0; j < len(fns); j++ {
			select {
			case err := <-errChan:
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
