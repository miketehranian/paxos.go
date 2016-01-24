package main

import (
	"fmt"
	"log"
	"os"
)

// Process is a thread with a process identifier, a queue of incoming messages, and an "environment" that keeps track of all proceses and queues
type Process struct {
	inbox []string // TODO make this a slice of type funcs which are go-routines
	env   Environment
	id    int
}

func (p *Process) run() {
	var err error

	// Start a goroutine here that will run the body of the process

	// Remove the completed process from env, based on the ID

	p.env.removeProc(p.id)
	if err != nil {
		log.Fatal("Exiting")
	}
}

func (p *Process) getNextMessage() string {
	return p.inbox[0]
}

func (p *Process) sendMessage(dest, msg string) {
	// Call the env with the intended dest and the msg
}

func (p *Process) deliver(msg string) {
	append(p.inbox, msg)
}
