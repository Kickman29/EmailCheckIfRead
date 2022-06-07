package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"main.go/handlers"
)

func main() {

	// init logger
	lh := log.New(os.Stdout, "read-checker ", log.LstdFlags)
	h := handlers.NewHandler(lh)

	ld := log.New(os.Stdout, "dashboard ", log.LstdFlags)
	d := handlers.NewDashboard(ld)

	// init serve multiplexer
	hsm := mux.NewRouter()
	dsm := mux.NewRouter()

	log.Println("Starting services")

	// init new handlers for a request on given addresses
	hsm.Handle("/{name:[a-zA-Z]+}", &h)
	// dashboard + css styling

	// FIXME: CSS FILES ARE NOT BEING SERVED

	// dsm.HandleFunc("/style.css", func(rw http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(rw, r, "/resources/style.css")
	// })
	dsm.HandleFunc("/style.css", handlers.CSSStyling)
	dsm.HandleFunc("/clearlog", handlers.ClearLog)
	dsm.Handle("/", &d)

	log.Print("Listener started at port :80...")
	go http.ListenAndServe(":80", hsm)

	log.Println("Dashboard started at port :9090...")
	go http.ListenAndServe(":9090", dsm)

	// wait for 290 years
	time.Sleep(time.Duration(int64(9223372036854775807)))

}
