//:3000/testにgetするとhellow world
package main

import (
	"log"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
)

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Get("/test", Gettest),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":3000", api.MakeHandler()))
}

func Gettest(w rest.ResponseWriter, r *rest.Request) {
	v := r.URL.Query()
	num := v.Get("num")
	log.Print(v)
	log.Print(num)

	w.WriteHeader(http.StatusOK)
	w.WriteJson(map[string]string{"Body": "Hello World!"})
}
