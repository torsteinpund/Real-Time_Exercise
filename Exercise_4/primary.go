package main

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"strconv"
	"time"
)

const (
	UDP_PORT       = ":30001"
	TIMEOUT        = 3 * time.Second
	BROADCAST_FREQ = 1 * time.Second
)

func startPrimary(counter int) {
	addr, err := net.ResolveUDPAddr("udp", UDP_PORT)
	if err != nil {
		fmt.Println("Error resolving UDP address:", err)
		os.Exit(1)
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println("Error dialing UDP:", err)
		os.Exit(1)
	}
	defer conn.Close()

	// Start a backup process

	spawnBackup(counter)

	// Give the backup some time to start before sending messages
	time.Sleep(1 * time.Second)

	for {
		counter++
		fmt.Println("Primary counting:", counter)

		message := strconv.Itoa(counter)
		_, err := conn.Write([]byte(message))
		if err != nil {
			fmt.Println("Error sending UDP message:", err)
			os.Exit(1)
		}

		time.Sleep(BROADCAST_FREQ)
	}
}

func startBackup() {
	addr, err := net.ResolveUDPAddr("udp", UDP_PORT)
	if err != nil {
		fmt.Println("Error resolving UDP address:", err)
		os.Exit(1)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Error listening on UDP:", err)
		os.Exit(1)
	}
	//defer conn.Close()

	fmt.Println("Backup is now listening on", UDP_PORT)

	buf := make([]byte, 1024)
	lastCounter := 0

	for {
		conn.SetReadDeadline(time.Now().Add(TIMEOUT))
		n, _, err := conn.ReadFromUDP(buf)

		if err != nil {
			fmt.Println("Primary process lost! Backup taking over...")
			conn.Close()
			time.Sleep(1 * time.Second)
			startPrimary(lastCounter)
			return
		}

		lastCounter, _ = strconv.Atoi(string(buf[:n]))
		fmt.Println("Backup received heartbeat:", lastCounter)
	}
}

func spawnBackup(counter int) {
	cmd := exec.Command("gnome-terminal", "--", "bash", "-c", os.Args[0]+" backup "+strconv.Itoa(counter))
	//cmd := exec.Command(os.Args[0], "backup", strconv.Itoa(counter))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Failed to start backup:", err)
		os.Exit(1)
	} else {
		fmt.Println("Backup process started successfully.")
	}
}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "backup" {
		fmt.Println("Backup process started, monitoring primary...")
		startBackup()
	} else {
		fmt.Println("Primary process started.")
		startPrimary(0)
	}
}
