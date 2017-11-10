package main

import (
	"log"

	"github.com/gobuffalo/envy"

	"github.com/bketelsen/talks/codemotion/gc18/gophercon/actions"
)

// ServiceName is the string name of the service, since
// it is used in multiple places, it's an exported Constant

func main() {

	port := envy.Get("PORT", "3001")
	app := actions.App()

	log.Fatal(app.Start(port))
}
