package main

import (
	"fmt"
	"bsu.ru/messenger/server/database/models"
	"encoding/json"
)

type Notifier struct {
	Clients map[string](chan []byte)
	MainChannel chan mysql.ChatUpdateInfo
}

func (n *Notifier) Broadcast(done <-chan interface{}) {
	for {
		//var data mysql.ChatUpdateInfo
		select {
		case <-done:
			return
		case data := <-n.MainChannel:
			js, _ := json.Marshal(data)
			fmt.Println("Message")
			for _, v := range data.Users {
				ch, has := n.Clients[v]
				if has {
					ch <- js
				}
			}
		}
	}
}