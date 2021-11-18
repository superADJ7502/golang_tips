package synctool

import (
	"encoding/json"
	"sync"
)

type Student struct {
	Name   string
	Age    int32
	Remark [1024]byte
}

func Pool() {
	var studentPool = sync.Pool{
		New: func() interface{} {
			return new(Student)
		},
	}
	var buf, _ = json.Marshal(Student{Name: "Geektutu", Age: 25})
	stu := studentPool.Get().(*Student)
	err := json.Unmarshal(buf, stu)
	if err != nil {
		return
	}
	studentPool.Put(stu)
}
