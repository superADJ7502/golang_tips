package unlimitcachechannel

import "sync/atomic"

type T interface{}

type UnboundedChan struct {
	bufCount int64
	In       chan<- T
	Out      <-chan T
	buffer   *RingBuffer
}

func (c UnboundedChan) Len() int {
	return len(c.In) + c.BufLen() + len(c.Out)
}

func (c UnboundedChan) BufLen() int {
	return int(atomic.LoadInt64(&c.bufCount))
}

func NewUnboundedChan(initCapacity int) *UnboundedChan {
	return NewUnboundedChanSize(initCapacity, initCapacity, initCapacity)
}

func NewUnboundedChanSize(initInCapacity, initOutCapacity, initBufCapacity int) *UnboundedChan {
	in := make(chan T, initInCapacity)
	out := make(chan T, initOutCapacity)
	ch := UnboundedChan{
		In:     in,
		Out:    out,
		buffer: NewRingBuffer(initBufCapacity),
	}

	go process(in, out, &ch)

	return &ch
}

func process(in, out chan T, ch *UnboundedChan) {
	defer close(out)
loop:
	for {
		val, ok := <-in
		if !ok {
			break loop
		}

		if atomic.LoadInt64(&ch.bufCount) > 0 {
			ch.buffer.Write(val)
			atomic.AddInt64(&ch.bufCount, 1)
		} else {
			select {
			case out <- val:
				continue
			default:
			}

			ch.buffer.Write(val)
			atomic.AddInt64(&ch.bufCount, 1)
		}

		for !ch.buffer.IsEmpty() {
			select {
			case val, ok := <-in:
				if !ok {
					break loop
				}
				ch.buffer.Write(val)
				atomic.AddInt64(&ch.bufCount, 1)
			case out <- ch.buffer.Peek():
				ch.buffer.Pop()
				atomic.AddInt64(&ch.bufCount, -1)
				if ch.buffer.IsEmpty() && ch.buffer.size > ch.buffer.initialSize {
					ch.buffer.Reset()
					atomic.StoreInt64(&ch.bufCount, 0)
				}
			}
		}
	}

	for !ch.buffer.IsEmpty() {
		out <- ch.buffer.Pop()
		atomic.AddInt64(&ch.bufCount, -1)
	}
	ch.buffer.Reset()
	atomic.StoreInt64(&ch.bufCount, 0)
}
