package main

import router "github.com/sumelms/sumelms/user/pkg/router/json"

func main() {
	s, err := router.NewRouter()
	if err != nil {
		panic(err)
	}

	if err := s.Start(); err != nil {
		panic(err)
	}
}
