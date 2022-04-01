package home

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func ShowHomePage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Simple Chat Bot",
	})
}

//func ServeWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
//	//log.Println("Serve ws")
//	//conn, err := upgrader.Upgrade(w, r, nil)
//	//if err != nil {
//	//	log.Println(err)
//	//	return
//	//}
//	//client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256)}
//	//client.hub.register <- client
//	//go client.writePump()
//	//client.readPump()
//}

//webSocket returns json format
func StartWebSocket(c *gin.Context) {
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

func JsonWebsocket(c *gin.Context) {
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
