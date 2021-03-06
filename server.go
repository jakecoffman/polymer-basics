package main

import (
	"net/http"
	"log"
	"github.com/gorilla/handlers"
	"os"
	"github.com/gorilla/websocket"
)

func main() {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	http.Handle("/", handlers.LoggingHandler(os.Stderr, http.FileServer(http.Dir("."))))
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}
		defer conn.Close()
		for {
			t, msg, err := conn.ReadMessage()
			if err != nil {
				log.Println(err)
				return
			}
			log.Println("Received:", string(msg))
			err = conn.WriteMessage(t, msg)
			if err != nil {
				log.Println(err)
				return
			}
		}
	})
	host := "localhost:3030"
	log.Printf("http://%v", host)
	log.Fatal(http.ListenAndServe(host, nil))
}
