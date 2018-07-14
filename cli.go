package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/alexdavid/sigma/api"
)

func main() {
	command := os.Args[1]
	if command == "get-messages" {
		chatId, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}

		chats, err := api.GetMessages(chatId, time.Now().Add(time.Hour*-80))
		if err != nil {
			log.Fatal(err)
		}

		json, err := json.Marshal(chats)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(json))
		return
	}

	if command == "get-chats" {
		chats, err := api.GetChats()
		if err != nil {
			log.Fatal(err)
		}

		json, err := json.Marshal(chats)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(json))
		return
	}

	log.Fatal("Unkonwn command")
}
