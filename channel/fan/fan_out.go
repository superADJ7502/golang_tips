package main

import (
	"reflect"
	"sync"
)

func fanOut(ch <-chan interface{}, out []chan interface{}) {
	go func() {
		defer func() {
			for i := range out {
				close(out[i])
			}
		}()

		for v := range ch {
			v := v
			for i := range out {
				i := i
				out[i] <- v
			}
		}
	}()
}

func fanOutAsync(ch <-chan interface{}, out []chan interface{}) {
	go func() {
		var wg sync.WaitGroup
		defer func() {
			wg.Wait()
			for i := range out {
				close(out[i])
			}
		}()

		for v := range ch {
			v := v
			for i := range out {
				i := i
				wg.Add(1)
				go func() {
					out[i] <- v
					wg.Done()
				}()
			}
		}
	}()
}

func fanOutReflect(ch <-chan interface{}, out []chan interface{}) {
	go func() {
		defer func() {
			for i := range out {
				close(out[i])
			}
		}()
		cases := make([]reflect.SelectCase, len(out))

		for i := range cases {
			cases[i].Dir = reflect.SelectSend
		}
		for v := range ch {
			v := v
			// 先完成send case构造
			for i := range cases {
				cases[i].Chan = reflect.ValueOf(out[i])
				cases[i].Send = reflect.ValueOf(v)
			}
			// 遍历select
			for range cases {
				chosen, _, _ := reflect.Select(cases)
				// 已发送过，用nil阻塞，避免再次发送
				cases[chosen].Chan = reflect.ValueOf(nil)
			}
		}
	}()
}
