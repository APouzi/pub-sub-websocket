package main

import (
	"net/http"
	"strings"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

func main() {
	println("Golang Server Started")
	http.HandleFunc("/ws-upgrade",wsUpgrade)
	http.ListenAndServe(":8080",nil)
}
}