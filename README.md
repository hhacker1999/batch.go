# Batch

A tiny Go package to run goroutines in a limited, batched manner — ensuring that at most `N` goroutines run at any given time.

## 💡 What is it?

This package lets you **batch goroutines** — so you can run multiple tasks concurrently, but limit how many can run in parallel.

It’s useful when:

- You’re hitting API rate limits.
- You want to avoid overloading CPU/Memory.
- You need basic concurrency control without using worker pools or semaphores.

## 🚀 Installation

```bash
go get github.com/hhacker1999/batch.go/batch
```

## 🛠 Usage

```go
package main

import (
	"fmt"
	"time"
	"github.com/hhacker1999/batch.go/batch"
)

func printSomething(id int) {
	fmt.Printf("Start %d\n", id)
	time.Sleep(2 * time.Second)
	fmt.Printf("End %d\n", id)
}

func main() {
	b := batch.New(3) // Max 3 concurrent goroutines

	for i := 0; i < 10; i++ {
		b.AddWork(printSomething, i)
	}

	time.Sleep(10 * time.Second) // wait for all to complete
}
```

## ⚙️ How it works

- You call `AddWork(yourFunc, args...)`
- If the number of active goroutines is below the limit, your function runs in a goroutine.
- If not, it **waits** until one finishes.
- Reflection is used to call arbitrary functions.

## 📦 API

### `New(size int) *Batch`

Creates a new `Batch` with a maximum number of concurrent goroutines.

### `AddWork(fn interface{}, args ...interface{})`

Runs the function `fn` with provided arguments if a slot is available. Waits otherwise.

## 🧠 Internals

This package uses:

- `sync.Mutex` to safely update count.
- `reflect` to allow calling any function type.

## ✅ Possible Improvements

- Replace blocking loop with sync.Cond
- Add a `Wait()` method to allow blocking until all tasks finish.
- Add context support to cancel pending jobs if needed.

## 🔒 License

MIT
