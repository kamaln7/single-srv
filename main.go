package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gabriel-vasile/mimetype"
)

var (
	verbose bool
	addr    string
)

func main() {
	flag.BoolVar(&verbose, "verbose", false, "verbose mode")
	flag.StringVar(&addr, "addr", ":8000", "address to listen on")
	flag.Parse()

	args := flag.Args()
	if len(args) != 1 {
		log.Fatalln("usage: single-srv <path to file to serve>")
	}

	path := args[0]
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if verbose {
			log.Printf("[%s] %s\n", r.RequestURI, r.RemoteAddr)
		}

		body, err := ioutil.ReadFile(path)
		if err != nil {
			w.WriteHeader(500)
			return
		}
		mime := mimetype.Detect(body)

		w.Header().Set("Content-Type", mime.String())
		w.Write(body)
	})

	log.Printf("starting http server on %s\n", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Printf("error starting http server: %v\n", err)
	}
}
