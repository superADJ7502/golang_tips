package main

import (
	"fmt"
	"log"
)

const (
	pluginName    = "testplugin"
	pluginVersion = 0x00010000
)

func Load(register func(name string, version uint64) error) error {
	err := register(pluginName, pluginVersion)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	log.Println("loading test plugin")
	return nil
}

func Unload() error {
	fmt.Printf("unload %s, version: 0x%x\n", pluginName, pluginVersion)
	return nil
}

func Test(data string) string {
	return "hello " + data
}

func Test2Data(data string) string {
	return "Test2Data " + data
}
