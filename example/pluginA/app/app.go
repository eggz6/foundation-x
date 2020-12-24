package app

import (
	"context"
	"log"
	"net/http"

	"github.com/guanwenbogit/foundation-x/example"
	"github.com/guanwenbogit/foundation-x/plugin"
)

var mux http.ServeMux

func Launch() {
	mux.HandleFunc("/ping", ping)
	ctx := context.WithValue(context.Background(), example.HttpMux, &mux)

	err := plugin.Run(ctx)

	if err != nil {
		log.Printf("[warning] run plugin failed. err=%v", err)
	}

	log.Fatal(http.ListenAndServe(":12345", &mux))
}

func ping(w http.ResponseWriter, req *http.Request) {
	_, _ = w.Write([]byte("pong"))
}
