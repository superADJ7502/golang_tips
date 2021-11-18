package main

import (
	"fmt"
)

func main() {
	//valueCtx := context.WithValue(context.Background(), "key", "value")
	//ctx, cancel := context.WithCancel(valueCtx)
	//go func() {
	//	ticker := time.NewTicker(1 * time.Second)
	//	for {
	//		select {
	//		case <-ticker.C:
	//			fmt.Println(ctx.Value("key"))
	//			ticker.Reset(1 * time.Second)
	//		}
	//	}
	//}()
	//
	//time.Sleep(5 * time.Second)
	//cancel()

	s1 := "1"
	//s2 := "1"
	var s3 interface{} = "1"
	//s3 = s2

	fmt.Println(s1 == s3)
}
