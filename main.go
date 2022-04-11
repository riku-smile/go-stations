package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/TechBowl-japan/go-stations/db"
	"github.com/TechBowl-japan/go-stations/handler"
	"github.com/TechBowl-japan/go-stations/handler/router"
	"github.com/TechBowl-japan/go-stations/model"
)

func main() {
	err := realMain()
	if err != nil {
		log.Fatalln("main: failed to exit successfully, err =", err)
	}
}

func realMain() error {
	// config values
	const (
		defaultPort   = ":8080"
		defaultDBPath = ".sqlite3/todo.db"
	)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = defaultDBPath
	}

	// set time zone
	var err error
	time.Local, err = time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return err
	}

	// set up sqlite3
	todoDB, err := db.NewDB(dbPath)
	if err != nil {
		return err
	}
	defer todoDB.Close()

	// set http handlers
	mux := router.NewRouter(todoDB)

	// TODO: ここから実装を行う
	http.ListenAndServe(port, mux)

	mux.Handle("/healthz", handler.NewHealthzHandler())

	mux.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		var ctreq model.CreateTODORequest
		// var ctres model.CreateTODOResponse
		if r.Method == "POST" {
			err := json.NewDecoder(r.Body).Decode(&ctreq)
			if err != nil {
				fmt.Println(err)
			}
		}

		// TODO: エンコード設定
		if ctreq.Subject == "" {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			handler.NewTODOHandler().Create(r.Context(), &ctreq)
		}

	})
	return nil
}
