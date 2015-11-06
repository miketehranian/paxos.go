package main

import "os"

type Process struct {
	inbox []string
	env   string
	id    int
}

func (p *Process) run() {
	var err error

	// Start a goroutine here that will run the body of the process

	// Remove the completed process from env, based on the ID

	if err != nil {
		os.Exit(1)
	}
}

func (p *Process) getNextMessage() string {
	return p.inbox[0]
}

func (p *Process) sendMessage(dest, msg string) {
	// Call the env with the intended dest and the msg

}

// func (p *Process) deliver(msg string) error {
// 	append(p.inbox, msg)

// 	return p.index[0]
// }
