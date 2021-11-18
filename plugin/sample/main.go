package main

import (
	"log"
	"plugin"
)

type CustomPlugin interface {
	CallMe(name string) string
}

func main() {
	p, err := plugin.Open("plugin.so")
	if err != nil {
		panic(err)
	}
	sayHelloPlugin, err := p.Lookup("SayHelloPlugin")
	if err != nil {
		panic(err)
	}
	if sayHello, ok := sayHelloPlugin.(CustomPlugin); ok {
		log.Println(sayHello.CallMe("to get to you"))
	}
}
