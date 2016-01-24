package main

import "log"

type Acceptor struct {
	Process
	ballotNumber int
	accepted     []int
}

func (a *Acceptor) body() {

	log.Printf("Here I am: %s", a.id)

	for {
		msg := getNextMessage()
		switch msgType := msg.(type) {
		case P1aMessage:
			if msg.ballotNumber > a.ballotNumber {
				a.ballotNumber = msg.ballotNumber
			}
			sendMessage(msg.src, &P1bMessage{id, ballotNumber, accepted})
			return
		case P2aMessage:
			if msg.ballotNumber == a.ballotNumber {
				append(accepted, &PValue{msg.ballotNumber, msg.slotNumber, msg.command})
			}
			sendMessage(msg.src, &P2bMessage{id, ballotNumber, accepted})
			return
		default:
			log.Println("Scout: unexpected msg")
		}
	}
}
