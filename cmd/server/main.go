package main

import "github.com/sumelms/sumelms/user/pkg/router/http"

func main() {
	s, err := http.NewHttpServer()
	if err != nil {
		panic(err)
	}

	if err := s.Start(); err != nil {
		panic(err)
	}
}
