package main

import (
	"net/http"

	"github.com/ridhoafwani/gingormpostgres/routers"
)

func main() {
	routes := routers.NewMainRouter()

	server := &http.Server{
		Addr:    ":8888",
		Handler: routes,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
