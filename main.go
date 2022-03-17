package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"
	"time"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
var addr = flag.String("addr", ":8005", "http service address")

// webSocket returns text format
func textApi(c *gin.Context) {
	//Upgrade get request to webSocket protocol
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("error get connection")
		log.Fatal(err)
	}
	defer ws.Close()
	//Read data in ws
	mt, message, err := ws.ReadMessage()
	if err != nil {
		log.Println("error read message")
		log.Fatal(err)
	}

	//Write ws data, pong 10 times
	var count = 0
	for {
		count++
		if count > 10 {
			break
		}

		message = []byte(string(message) + " " + strconv.Itoa(count))
		err = ws.WriteMessage(mt, message)
		if err != nil {
			log.Println("error write message: " + err.Error())
		}
		time.Sleep(1 * time.Second)
	}
}

//webSocket returns json format
func jsonApi(c *gin.Context) {
	//Upgrade get request to webSocket protocol
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("error get connection")
		log.Fatal(err)
	}
	defer ws.Close()
	var data struct {
		A string `json:"a"`
		B int    `json:"b"`
	}
	//Read data in ws
	err = ws.ReadJSON(&data)
	if err != nil {
		log.Println("error read json")
		log.Fatal(err)
	}

	//Write ws data, pong 10 times
	var count = 0
	for {
		count++
		if count > 10 {
			break
		}

		err = ws.WriteJSON(struct {
			A string `json:"a"`
			B int    `json:"b"`
			C int    `json:"c"`
		}{
			A: data.A,
			B: data.B,
			C: count,
		})
		if err != nil {
			log.Println("error write json: " + err.Error())
		}
		time.Sleep(1 * time.Second)
	}
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	log.Println(r.URL.Path)
	if r.URL.Path != "/" {
		http.Error(w, "Not foundeee", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	log.Println("Serve source")
	http.ServeFile(w, r, "./public/index.html")
}

func main() {
	chatbot := NewWebAssistant()

	hub := NewHub(chatbot)

	startChatHub(hub)
	log.Println("Server started on port: " + *addr)
	http.HandleFunc("/", serveHome)
	//http.HandleFunc("/ws", ServeWs(hub))
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ServeWs(hub, w, r)
	})
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	//r := gin.Default()
	//r.GET("/json", jsonApi)
	//r.GET("/text", textApi)
	//
	//r.GET("/ws", ServeWs(hub))
	//
	//// static files
	//r.Use(static.Serve("/", static.LocalFile("./public", true)))
	//
	//r.NoRoute(func(c *gin.Context) {
	//	c.File("./public/index.html")
	//})
	//
	//r.Run(":8005")
}
