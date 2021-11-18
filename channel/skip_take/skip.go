package skip_take

func skipN(done <-chan struct{}, valueStream <-chan interface{}, num int) <-chan interface{} {
	takeStream := make(chan interface{})

	go func() {
		defer close(takeStream)

		for i := 0; i < num; i++ {
			select {
			case <-done:
				return
			case <-valueStream:
			}
		}

		for {
			select {
			case <-done:
				return
			case takeStream <- <-valueStream:
			}
		}
	}()
	return takeStream
}

func skipFn(done <-chan struct{}, valueStream <-chan interface{}, fn func(interface{}) bool) <-chan interface{} {
	takeStream := make(chan interface{})
	go func() {
		defer close(takeStream)

		for {
			select {
			case <-done:
				return
			case v := <-valueStream:
				if !fn(v) {
					takeStream <- v
				}
			}
		}
	}()
	return takeStream
}

func skipWhile(done <-chan struct{}, valueStream <-chan interface{}, fn func(interface{}) bool) <-chan interface{} {
	takeStream := make(chan interface{})

	go func() {
		defer close(takeStream)

		take := false
		for {
			select {
			case <-done:
				return
			case v := <-valueStream:
				if !take {
					take = !fn(v)
					if !take {
						continue
					}
				}
				takeStream <- v
			}
		}
	}()
	return takeStream
}
