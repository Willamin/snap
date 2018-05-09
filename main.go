package main

import (
	"fmt"
	"github.com/mitchellh/go-ps"
	"math/rand"
	"os"
	"syscall"
)

func main() {
	processes, err := ps.Processes()

	if err != nil {
		panic(err)
	}

	fmt.Printf("There are %d processes.\n", len(processes))

	rand.Shuffle(len(processes), func(i, j int) {
		processes[i], processes[j] = processes[j], processes[i]
	})

	var toKill []int

	for index := 0; index < len(processes)/2; index++ {
		process := processes[index]
		toKill = append(toKill, process.Pid())
	}

	fmt.Printf("Killing %d processes.", len(toKill))

	for _, pid := range toKill {
		if pid != 1 || pid != os.Getpid() {
			process, _ := os.FindProcess(pid)
			process.Signal(syscall.SIGKILL)
		}
	}
}
