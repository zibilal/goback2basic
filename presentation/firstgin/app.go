package main

import (
	"goback2basic/presentation/firstgin/router"
	"log"
)

func main() {
	r := router.SetupEngine(router.Ping, router.PingAction)
	log.Fatal(r.Run())
}
