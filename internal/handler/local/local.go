package local

import (
	"context"
	"log"
	"net/http"
	"os"
	c "ssh_connections_manager/internal/entity/client"
	"ssh_connections_manager/internal/utils"

	"github.com/gorilla/websocket"
)


var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}



var client *c.Client



func ServeWS(w http.ResponseWriter, r *http.Request) {

	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Ошибка апргейда: %V", err)
		return
	}

	defer conn.Close()

	client = c.NewClient(conn, os.Getenv("HOME"))

	//client.Send(client.CurrentDir)

	for {
		_, command, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsCloseError(err, websocket.CloseGoingAway, websocket.CloseNormalClosure) {
				log.Println("Клиент закрыл соединение:", err)
			} else {
				log.Println("Ошибка чтения сообщения:", err)
			}
			return
		}

		output, err := client.Execute(context.Background(), string(command))
		if err != nil {
			log.Println("Ошибка выполнения команды:", err)
			if err = conn.WriteMessage(websocket.TextMessage, []byte("Ошибка выполнения команды")); err != nil {
				log.Println("Ошибка отправки сообщения об ошибки выполнения команды")
			}
			continue
		}

		err = conn.WriteMessage(websocket.TextMessage, output)
		if err != nil {
			log.Println("Ошибка отправки сообщения:", err)
			break
		}
	}
}


func PromptHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()


	output := utils.GetPrompt()

	err = conn.WriteMessage(websocket.TextMessage, []byte(output))
	if err != nil {
		log.Println("Ошибка отправки промпта:", err)
		return
	}
}


func UserNameHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
		return
	}
	defer conn.Close()

	output := utils.GetUserName()

	err = conn.WriteMessage(websocket.TextMessage, []byte(output))
	if err != nil {
		log.Println("Ошибка отправки промпта:", err)
		return
	}
}

func HostHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	defer conn.Close()


	output := utils.GetHost()

	err = conn.WriteMessage(websocket.TextMessage, []byte(output))
	if err != nil {
		log.Println("Ошибка отправки промпта:", err)
		return
	}
}

func DirHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	defer conn.Close()


	output := client.ShowDir()

	err = conn.WriteMessage(websocket.TextMessage, []byte(output))
	if err != nil {
		log.Println("Ошибка отправки промпта:", err)
		return
	}
}