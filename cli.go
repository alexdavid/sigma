package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/alexdavid/sigma/sigma"
)

func main() {
	command := os.Args[1]
	if command == "get-messages" {
		chatId, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}

		messages, err := sigma.GetMessages(chatId, time.Now().Add(time.Hour*-80))
		if err != nil {
			log.Fatal(err)
		}

		json, err := json.Marshal(messages)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(json))
		return
	}

	if command == "get-attachments" {
		messageId, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}

		attachments, err := sigma.GetAttachments(messageId)
		if err != nil {
			log.Fatal(err)
		}

		json, err := json.Marshal(attachments)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(json))
		return
	}

	if command == "get-chats" {
		chats, err := sigma.GetChats()
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
