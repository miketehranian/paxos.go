package main

type Commander struct {
	Process
	Leader
	acceptors    []Acceptor
	replicas     []Replicas
	ballotNumber int
	slotNumber   int
	Command
}

func NewCommander() Commander {
	commander := &Commander{&Process{env, id}, leader, acceptors, replicas, ballotNumber, slotNumber, command}
	commander.env.addProc(commander)
}

func (c *Commander) body() {
	waitfor := make(map[Acceptor]bool)

	for a := range c.acceptors {
		c.sendMessage(a, &P2aMessage(c.id, c.ballotNumber, c.slotNumber, c.command))
	}
	if _, contains := waitfor[a]; contains {
		append(waitfor, a)
	}

	for {
		msg := c.getNextMessage()

		switch msgType := msg.(type) {
		case P2bMessage:
			if _, contains := waitfor; c.ballotNumber == msg.ballotNumber && contains {
				delete(waitfor, msg.src)

				if len(waitfor) < len(c.acceptors) {
					for r := range c.replicas {
						c.sendMessage(r, &DecisionMessage{c.id, c.slotNumber, c.command})
						return
					}
				}

			} else {
				c.sendMessage(c.leader, &PreemptedMessage{c.id, msg.ballotNumber})
				return
			}
		}
	}
}
