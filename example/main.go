package main

import (
	"fmt"

	"github.com/reactivex/rxgo"
	"github.com/reactivex/rxgo/handlers"
)

func main() {
	obs, err := rxgo.Range(0, 2)

	if err != nil {
		return
	}

	obs.
		Map(func(v interface{}) interface{} {
			val, ok := v.(int)
			if !ok {
				fmt.Println(ok)
			}

			obs, err := rxgo.Range(val, 2)
			if err != nil {
				fmt.Println(err)
			}

			return obs
		}).
		Switch().
		Subscribe(handlers.NextFunc(func(v interface{}) {
			fmt.Println("------", v)
		})).
		Block()
}
