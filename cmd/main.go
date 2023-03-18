package main

import (
	"api/internal/app"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	if len(os.Args) < 1 {
		log.Println("Nedd to point config file")
		return
	}

	app := app.App{}

	app.Init(os.Args[1])

	defer app.Close()

	go app.Run()

	// handle ctr+c.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Println("\nServer stopped")

}
