// Copyright 2018 Kirill Lapshin. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, GoldenSplit!"))
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", indexHandler)

	server := &http.Server{
		Handler:      router,
		Addr:         "0.0.0.0:80",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Output(1, "Server started...")
	log.Fatal(server.ListenAndServe())
}
