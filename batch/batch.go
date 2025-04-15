package batch

import (
	"reflect"
	"sync"
)

type Batch struct {
	size       int
	count      int
	countMutex *sync.Mutex
}

func New(size int) *Batch {
	if size <= 0 {
		panic("Invalid size value")
	}

	return &Batch{
		size:       size,
		countMutex: &sync.Mutex{},
	}
}

func (b *Batch) AddWork(work interface{}, args ...interface{}) {

	// Note: Block until goroutines are free
	for {
		b.countMutex.Lock()
		if b.count < b.size {
			b.count += 1
			b.countMutex.Unlock()
			break
		}

		b.countMutex.Unlock()
	}

	val := reflect.ValueOf(work)
	if val.Kind() != reflect.Func {
		panic("Work can only be a function")
	}

	if len(args) != val.Type().NumIn() {
		panic("Argument count mismatch")
	}

	valArgs := make([]reflect.Value, len(args))
	for i, v := range args {
		valArgs[i] = reflect.ValueOf(v)
	}

	go func(a []reflect.Value) {
		val.Call(a)
		b.countMutex.Lock()
		b.count -= 1
		b.countMutex.Unlock()
	}(valArgs)
}
