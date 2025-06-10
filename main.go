// /main.go
package main

import (
	"Voidside/config"
	"Voidside/routes"
)

func main() {
	config.ConnectDatabase()  // instantiate DB and store globally
	r := routes.SetupRouter() // Return Gin engine with routes mounted.
	r.Run(":8080")            // start server on default port
}
