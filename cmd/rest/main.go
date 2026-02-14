package main

import "test-task/internal/app"

func main() {
	a, err := app.New()
	if err != nil {
		panic(err)
	}
	_ = a
}
