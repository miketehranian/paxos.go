package main

import "fmt"

const WINDOW = 5

type BallotNumber struct {
	round    int
	leaderId string
}

func (b *BallotNumber) String() string {
	return fmt.Sprintf("BN(%d, %s)", b.round, b.leaderId)
}

type PValue struct {
	ballotNumber, slotNumber int
	Command
}

func (p *Pvalue) String() string {
	return fmt.Sprintf("PV(%s, %s, %s)", p.ballotNumber, p.slotNumber, p.Command)
}

type Command struct {
	client, reqId, op string
}

func (c *Command) String() string {
	return fmt.Sprintf("Command(%s, %s, %s)", c.client, c.reqId, c.op)
}

type ReconfigCommand struct {
	client, reqId, config string
}

func (rc *ReconfigCommand) String() string {
	fmt.Sprintf("ReconfigCommand(%s, %s, %s)", rc.client, rc.reqId, rc.config)
}

type Config struct {
	replicas  []Replica
	acceptors []Acceptor
	leaders   []Leader
}

func (c *Config) String() string {
	fmt.Sprintf("Config(%v, %v, %v)", c.replicas, c.acceptors, c.leaders)
}
