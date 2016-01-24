package main

type Environment struct {
	procs map[int]Process
}

func (e *Environment) sendMessage() {

}

func (e *Environment) addProc(proc Process) {
	procs[proc.id] = proc
}

func (e *Environment) removeProc(pid int) {
	delete(procs, int)
}
