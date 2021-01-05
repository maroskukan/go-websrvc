package main

import (
	"net/http"

	"github.com/maroskukan/go-websrvc/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
