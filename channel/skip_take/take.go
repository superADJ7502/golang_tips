package skip_take

func takeN(done <-chan struct{}, valueStream <-chan interface{}, num int) <-chan interface{} {
	takeStream := make(chan interface{})

	go func() {
		defer close(takeStream)
		for i := 0; i < num; i++ {
			select {
			case <-done:
				return
			case takeStream <- <-valueStream:
			}
		}
	}()
	return takeStream
}

func takeFn(done <-chan struct{}, valueStream <-chan interface{}, fn func(interface{}) bool) <-chan interface{} {
	takeStream := make(chan interface{})

	go func() {
		defer close(takeStream)
		for {
			select {
			case <-done:
				return
			case v := <-valueStream:
				if fn(v) {
					takeStream <- v
				}
			}
		}
	}()
	return takeStream
}

func takeWhile(done <-chan struct{}, valueStream <-chan interface{}, fn func(interface{}) bool) <-chan interface{} {
	takeStream := make(chan interface{})
	go func() {
		defer close(takeStream)
		for {
			select {
			case <-done:
				return
			case v := <-valueStream:
				if !fn(v) {
					return
				}
				takeStream <- v
			}
		}
	}()
	return takeStream
}
