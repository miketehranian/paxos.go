package main

type Scout struct {
	Process
	Leader       leader
	acceptors    []Acceptor
	ballotNumber int
}
