package main

import (
	"log"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

const (
	processName = "kdevtmpfsi"
	checkInterval = 5 * time.Second
)

func findAndKillProcess() {
	cmd := exec.Command("sh", "-c", "ps aux | grep "+processName+" | grep -v grep")
	
	output, err := cmd.Output()
	if err != nil {
		log.Printf("Proses '%s' not found or error executing command: %v", processName, err)
		return
	}

	line := string(output)
	fields := strings.Fields(line)

	if len(fields) < 2 {
		log.Printf("Output format is invalid: %s", line)
		return
	}

	pidStr := fields[1]
	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		log.Printf("Failed to convert PID to number: %s", pidStr)
		return
	}

	log.Printf("!!! Process '%s' detected with PID: %d. Stopping process...", processName, pid)

	// Send KILL signal (kill -9) to the process
	killCmd := exec.Command("kill", "-9", strconv.Itoa(pid))
	err = killCmd.Run()
	if err != nil {
		log.Printf("Failed to stop process with PID %d: %v", pid, err)
	} else {
		log.Printf("Process with PID %d successfully stopped.", pid)
	}
}

func main() {
	log.Println("Start service process-killer...")
	log.Printf("Checking for process '%s' every %v.", processName, checkInterval)

	ticker := time.NewTicker(checkInterval)
	defer ticker.Stop()

	findAndKillProcess()
	for range ticker.C {
		findAndKillProcess()
	}
}