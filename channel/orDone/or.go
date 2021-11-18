package main

import "sync"

func OrOnce(channels ...<-chan interface{}) <-chan interface{} {
	out := make(chan interface{})
	go func() {
		var once sync.Once
		for _, c := range channels {
			go func(c <-chan interface{}) {
				select {
				case <-c:
					once.Do(func() {
						close(out)
					})
				case <-out:
				}
			}(c)
		}
	}()
	return out
}

func Or(channels ...<-chan interface{}) <-chan interface{} {
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}
	orDone := make(chan interface{})
	go func() {
		defer close(orDone)

		switch len(channels) {
		case 2:
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		default:
			m := len(channels) / 2
			select {
			case <-Or(channels[:m]...):
			case <-Or(channels[m:]...):
			}
		}
	}()
	return orDone
}

func OrRecur(channels ...<-chan interface{}) <-chan interface{} {
	// 特殊情况，只有0个或者1个chan
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}

	orDone := make(chan interface{})
	go func() {
		defer close(orDone)

		switch len(channels) {
		case 2: // 特殊情况
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		case 3: // 特殊情况
			select {
			case <-channels[0]:
			case <-channels[1]:
			case <-channels[2]:
			}
		default: // 超过3个，二分法递归处理
			m := len(channels) / 2
			select {
			case <-OrRecur(append(channels[:m:m], orDone)...):
			case <-OrRecur(append(channels[m:], orDone)...):
			}
		}
	}()

	return orDone
}
