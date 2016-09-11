package gogo

type Fn func() error
type Fns []Fn

func Run(fnQueue ...Fns) error {
	errChan := make(chan error)
	defer close(errChan)
	for i := 0; i < len(fnQueue); i++ {
		fns := fnQueue[i]
		for j := 0; j < len(fns); j++ {
			fn := fns[j]
			go func() {
				errChan <- fn()
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
