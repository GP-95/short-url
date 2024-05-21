package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/GP-95/short-url/internal/cache"
	"github.com/GP-95/short-url/internal/db"
	"github.com/GP-95/short-url/internal/url"
)

func main() {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		log.Fatal("Could not find 'PORT' in env. variables")
	}

	cache, err := cache.New()
	if err != nil {
		log.Fatal("Could not connect to Redis")
	}

	db, err := db.New("./database.db")
	if err != nil {
		log.Fatal("Could not read database")
	}

	url.RegisterHandlers()

	server := http.Server{
		Addr: ":" + port,
	}

	defer func() {
		cache.Close()
		db.Close()
		server.Close()
	}()

	fmt.Println("Listening on port: " + port)
	e := server.ListenAndServe()
	if e != nil {
		log.Fatal("Could not start HTTP server: ", e)
	}

}
