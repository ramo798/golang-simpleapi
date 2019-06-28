package main

import (
	"github.com/ramo798/golang-simpleapi/sql"
)

func main() {
	// api := rest.NewApi()
	// api.Use(rest.DefaultDevStack...)

	// router, err := rest.MakeRouter(
	// // rest.Get("/test", Gettest),
	// )
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// api.SetApp(router)
	// log.Fatal(http.ListenAndServe(":3000", api.MakeHandler()))
	sql.TestInit()
}
