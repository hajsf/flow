package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

//Human data storage
type MsgData struct {
	msg  string
	data string
}

func (u *MsgData) String() string {
	return fmt.Sprintf("MessageData{msg: %q, data: %q}", u.msg, u.data)
}

func main() {
	humanChan := make(chan MsgData)
	botChan := make(chan MsgData)

	go func() {
		for {
			select {
			//receive massage from human and redirect it to bot
			case toBot := <-botChan:
				go bot(toBot, humanChan)
			//receive massage from bot and redirect it to human
			case toHuman := <-humanChan:
				fmt.Println(toHuman.data)
				go human(toHuman.msg, botChan)
			}
		}
	}()

	bot(MsgData{msg: "register", data: ""}, humanChan)
	// listen for Ctrl+C
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	//lock until the end of bot flow

	//collectedData := <-bot.Start()
	//fmt.Println(collectedData)
}
