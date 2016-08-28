package main

import (
	"log"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
)

func main() {
	statusMw := &rest.StatusMiddleware{}

	api := rest.NewApi()
	api.Use(statusMw)
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Get("/status", func(w rest.ResponseWriter, r *rest.Request) {
			w.WriteJson(statusMw.GetStatus())
		}),
		rest.Get("/", func(w rest.ResponseWriter, r *rest.Request) {
			w.WriteJson(map[string]string{"body": "Hello World"})
		}),
		rest.Get("/golang", func(w rest.ResponseWriter, r *rest.Request) {
			w.WriteJson(map[string]string{"body": "This is from a Golang app"})
		}),
	)

	if err != nil {
		log.Fatal(err)
	}

	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}
