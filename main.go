package main

import (
	"net/http"

	"github.com/slns/webservice/controllers"
)

func main() {
	controllers.RegisterControllers()

	http.ListenAndServe(":3000", nil)
}
