package main

import (
	"net/http"
	"redirect-checker/internal/config"
	"redirect-checker/internal/handlers"
)

func main() {
	// rPath, err := os.Getwd()
	// if err != nil {
	// 	panic(err)
	// }
	err := config.LoadEnvConfig()
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/", handlers.HandlerCheckRedirect)
	err = http.ListenAndServe(config.Cfg.Port, nil)
	if err != nil {
		panic(err)
	}
}
