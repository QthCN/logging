package logging

import (
	"fmt"
	"path"
	"sync"
	"time"
)

var lock sync.Mutex

type SimpleHandler struct {
}

func (h *SimpleHandler) record(level string, s string, file string, line int) {
	_, f := path.Split(file)
	t := time.Now().Format("2006-01-02 15:04:05")
	lock.Lock()
	fmt.Printf("[%s] %s %s:%d - %s\n", t, level, f, line, s)
	lock.Unlock()
}
