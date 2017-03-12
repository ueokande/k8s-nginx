package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	portStr := os.Getenv("AP_PORT")

	if len(portStr) == 0 {
		log.Fatal("AP_PORT not speicified")
	}

	port, err := strconv.ParseUint(portStr, 10, 16)
	if err != nil {
		log.Fatal("invalid AP_PORT: ", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("from %v: %v %v", r.RemoteAddr, r.Method, r.RequestURI)
		fmt.Fprint(w, "application works\n")
	})

	bind := fmt.Sprintf(":%d", port)
	log.Fatal(http.ListenAndServe(bind, nil))
}
