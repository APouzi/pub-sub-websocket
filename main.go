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

func wsUpgrade(w http.ResponseWriter, r *http.Request){
	print("Upgrading to WebSockets")
	conn, wr, _,err := ws.UpgradeHTTP(r,w)
	if err != nil{
		println("connection failed")
	}
	print(conn)
	for {
		payload, err := wsutil.ReadClientText(wr)
		if err != nil{
			print("err reading client text")
		}
		if strings.Contains(string(payload),"sub"){
			println("has sub")
		}
	}

}