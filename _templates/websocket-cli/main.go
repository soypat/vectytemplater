package main

import (
	"log"
	"net/http"
	"os"
	"vecty-templater-project/model"

	"github.com/soypat/gwasm"
)

func main() {
	// Register http server for serving WASM and other resources.
	wsm, err := gwasm.NewWASMHandler("app", nil)
	if err != nil {
		log.Fatal("NewWASMHandler failed", err)
	}
	wsm.WASMReload = true
	wsm.SetOutput(os.Stdout)
	http.Handle("/", wsm)
	http.Handle("/ws", &todoServer{})
	log.Printf("listening on http://[::]%v", model.HTTPServerAddr)
	log.Fatal(http.ListenAndServe(model.HTTPServerAddr, nil))
}
