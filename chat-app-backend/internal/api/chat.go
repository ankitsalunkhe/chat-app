package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func (h Handlers) GetChat(e echo.Context) error {
	userId, _ := h.auth.VerifySession(e)
	fmt.Println("Hi", userId)

	ws, err := upgrader.Upgrade(e.Response(), e.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			e.Logger().Error(err)
		}
		fmt.Printf("Recieved in backend :%s\n", msg)
	}
}
