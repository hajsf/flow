package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func human(msg string, botChan chan MsgData) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		botChan <- MsgData{
			msg:  "err",
			data: fmt.Sprint("error: ", err),
		}
	} else {
		text := scanner.Text()
		if strings.Compare("quit", text) == 0 {
			botChan <- MsgData{
				msg:  "quit",
				data: text,
			}
		} else {
			var id string
			switch msg {
			case "user", "pswd":
				id = msg
				botChan <- MsgData{
					msg:  id,
					data: text,
				}
			default:
				botChan <- MsgData{
					msg:  "",
					data: text,
				}
			}
		}
	}
}
