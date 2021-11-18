package main

import "reflect"

func roundRobin(ch <-chan interface{}, out []chan interface{}) {
	go func() {
		defer func() {
			for i := 0; i < len(out); i++ {
				close(out[i])
			}
		}()

		//round-robin
		var i = 0
		var n = len(out)
		for v := range ch {
			v := v
			out[i] <- v
			i = (i + 1) % n
		}
	}()
}

func roundRobinReflect(ch <-chan interface{}, out []chan interface{}) {
	go func() {
		defer func() {
			for i := 0; i < len(out); i++ {
				close(out[i])
			}
		}()

		cases := make([]reflect.SelectCase, len(out))
		for i := range cases {
			cases[i].Dir = reflect.SelectSend
			cases[i].Chan = reflect.ValueOf(out[i])
		}

		for v := range ch {
			v := v
			for i := range cases {
				cases[i].Send = reflect.ValueOf(v)
			}
			_, _, _ = reflect.Select(cases)
		}
	}()
}
