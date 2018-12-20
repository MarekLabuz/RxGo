package main

import (
	"fmt"
	"time"

	"github.com/reactivex/rxgo"
	"github.com/reactivex/rxgo/handlers"
)

func main() {
	until := rxgo.
		Just(1).
		Repeat(3, rxgo.WithDuration(time.Second)).
		Last()

	rxgo.
		Just(1).
		Repeat(10, rxgo.WithDuration(time.Second)).
		TakeUntil(until).
		Subscribe(handlers.NextFunc(func(_ interface{}) {
			fmt.Println("Still emitting...")
		}))

	until.
		Subscribe(handlers.DoneFunc(func() {
			fmt.Println("Stop!")
		})).Block()
}
