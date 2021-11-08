package main

import (
	"github.com/foekall/cattle-management/pkg/routes"
)

func main() {
	// r := gin.New()
	routes.RegisterCattleManagementRoutes()
	// http.Handle("/", r)
	// log.Fatal(http.ListenAndServe("localhost:8080", r))
}
