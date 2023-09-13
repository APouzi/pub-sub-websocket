package main

import (
	"net/http"
	"strings"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

var pubsubStart *PubSub

func main() {
	println("Golang Server Started")
	pubsubStart = InitiatePubSubFeed()
	http.HandleFunc("/ws-upgrade",pubsubStart.wsUpgrade)
	http.ListenAndServe(":8080",nil)

}


func(pbs *PubSub) wsUpgrade(w http.ResponseWriter, r *http.Request){
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