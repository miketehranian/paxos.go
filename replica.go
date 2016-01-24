package main

import "process"

// s = slot
// c = command

// Replica invariants:
// R1: There are no two different commands decided for the same slot
// R2: All commands up to slot_out are in the set of decisions
// R3: For all replicas p, p.state is the result of applying the commands
// R4: p.slot_out cannot decrease over time

type Replica struct {
	// the index of the next slot in which the replica has not yet proposed any command.
	slotIn int

	// the index of the next slot for which it needs to learn a decisison before it
	// can update its copy of the application state
	slotOut int

	// received requests that are not proposed or decided
	requests Request

	Process
}

func New(env string, id int) Replica {

	// return &Process{string, id}
}
