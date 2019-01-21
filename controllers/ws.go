package controllers

import (
	"log"
	"nkyWebServer/models"
	"time"

	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
)

type WebSocketController struct {
	beego.Controller
}

var upgrader = websocket.Upgrader{}
var mCH = make(chan bool)

func (c *WebSocketController) Get() {
	c.EnableRender = false
	ws, err := upgrader.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()
	defer beego.Info("退出ws")
	realtimeSend(ws)
}

func realtimeSend(conn *websocket.Conn) {
	for {
		time.Sleep(time.Second * 10)
		nodeDatas, err := models.GetAllNodeData()
		for _, nodedata := range nodeDatas {
			if err = conn.WriteJSON(nodedata); err != nil {
				beego.Error(err)
				return
			}
		}
	}
}
