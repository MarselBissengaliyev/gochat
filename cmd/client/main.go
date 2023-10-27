package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/MarselBissengaliyev/gochat/pkg/model"
	"github.com/MarselBissengaliyev/gochat/pkg/util"
	"golang.org/x/net/websocket"
)

var (
	port = flag.String("port", "8080", "port used for ws connection")
)

func main() {
	flag.Parse()

	ws, err := websocket.Dial(fmt.Sprintf("ws://localhost:%s", *port), "", util.MockedIp())
	if err != nil {
		log.Fatal(err)
		return
	}

	defer ws.Close()

	var m model.Message
	go func() {
		for {
			err := websocket.JSON.Receive(ws, &m)
			if err != nil {
				fmt.Println("Error receiving message: ", err.Error())
				break
			}
			fmt.Println("Message: ", m)
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		}

		m := model.Message{
			Text: text,
		}
		err := websocket.JSON.Send(ws, m)
		if err != nil {
			fmt.Println("Error sending message: ", err.Error())
			break
		}
	}
}