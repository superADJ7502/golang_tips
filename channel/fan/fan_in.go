package main

import (
	"reflect"
	"sync"
)

func fanIn(chans ...<-chan interface{}) <-chan interface{} {
	out := make(chan interface{})

	go func() {
		var wg sync.WaitGroup
		wg.Add(len(chans))
		for _, ch := range chans {
			go func(ch <-chan interface{}) {
				for v := range ch {
					out <- v
				}
				wg.Done()
			}(ch)
		}
		wg.Wait()
		close(out)
	}()
	return out
}

func mergeTwo(a, b <-chan interface{}) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		for a != nil || b != nil {
			select {
			case v, ok := <-a:
				if !ok {
					a = nil
					continue
				}
				c <- v
			case v, ok := <-b:
				if !ok {
					b = nil
					continue
				}
				c <- v
			}
		}
	}()
	return c
}

func fanInRecur(chans ...<-chan interface{}) <-chan interface{} {
	switch len(chans) {
	case 0:
		c := make(chan interface{})
		close(c)
		return c
	case 1:
		return chans[0]
	case 2:
		return mergeTwo(chans[0], chans[1])
	default:
		m := len(chans) / 2
		return mergeTwo(
			fanInRecur(chans[:m]...),
			fanInRecur(chans[m:]...))
	}
}

func fanInReflect(chans ...<-chan interface{}) <-chan interface{} {
	out := make(chan interface{})
	go func() {
		defer close(out)
		var cases []reflect.SelectCase
		for _, c := range chans {
			cases = append(cases, reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(c),
			})
		}

		for len(cases) > 0 {
			i, v, ok := reflect.Select(cases)
			if !ok {
				cases = append(cases[:i], cases[i+1:]...)
				continue
			}
			out <- v.Interface()
		}
	}()
	return out
}
