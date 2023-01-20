package main

import (
	"qsoft_test_task/internal/apiserver"
)

func main() {
	config := apiserver.NewConfig()
	s := apiserver.NewAPIServer(config)
	s.Start()
}