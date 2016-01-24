package main

// Paxos uses a large variety of message types. They are collected below.

type Message struct {
}

type P1aMessage struct {
	Message
	ballotNumber int
}

type P1bMessage struct {
	Message
	ballotNumber int
	accepted     bool
}

type P2aMessage struct {
	Message
	ballotNumber int
	slotNumber   int
	Command
}

type P2bMessage struct {
	Message
	ballotNumber int
	slotNumber   int
}

type PreemptedMessage struct {
	Message
	ballotNumber int
}

type AdoptedMessage struct {
	Message
	accepted bool
}

type DecisionMessage struct {
	Message
	slotNumber int
	Command
}

type RequestMessage struct {
	Message
	Command
}

type ProposeMessage struct {
	Message
	slotNumber int
	Command
}
