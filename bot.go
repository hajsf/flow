package main

import (
	"fmt"
	"os"
)

func bot(toBot MsgData, humanChan chan MsgData) {
	switch toBot.msg {
	case "register":
		humanChan <- MsgData{
			msg:  "user",
			data: "Enter user name: ",
		}
	case "user":
		if toBot.data == "Hasan" {
			humanChan <- MsgData{
				msg:  "pswd",
				data: "Enter password: ",
			}
		} else {
			fmt.Println("wrong user, bye")
			os.Exit(0)
		}
	case "pswd":
		if toBot.data == "123" {
			humanChan <- MsgData{
				msg:  "",
				data: "Welcome Hasan, what are you looking for ",
			}
		} else {
			humanChan <- MsgData{
				msg:  "pswd",
				data: "Wrong pasword, retry: ",
			}
		}
	case "quit":
		fmt.Println("Try later")
		os.Exit(0)
	default:
		humanChan <- MsgData{
			msg:  "data",
			data: "Repeat..: ",
		}
	}
}
