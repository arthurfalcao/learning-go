package main

import (
	"net/http"

	"github.com/arthurfalcao/learning-go/web/routes"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8080", nil)
}
