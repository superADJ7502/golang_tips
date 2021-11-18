package main

import (
	"fmt"
	"github.com/letiantech/hotplugin"
)

func main() {
	options := hotplugin.ManagerOptions{
		Dir:    "./",
		Suffix: ".so",
	}
	hotplugin.StartManager(options)
	result := hotplugin.Call("testplugin", "Test", "my world")
	fmt.Println(result...)
}
