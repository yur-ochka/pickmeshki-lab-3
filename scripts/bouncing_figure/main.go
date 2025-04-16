package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

func sendCommand(client *http.Client, cmd string) error {
	req, err := http.NewRequest("POST", "http://localhost:17000", strings.NewReader(cmd))
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	resp.Body.Close()
	return nil
}

func main() {
	client := &http.Client{}

	initCommands := []string{
		"white",
		"figure 0.5 0.5",
		"update",
	}

	for _, cmd := range initCommands {
		if err := sendCommand(client, cmd); err != nil {
			log.Fatal(err)
		}
	}

	time.Sleep(500 * time.Millisecond)

	x, y := 0.5, 0.5
	dx, dy := 0.1, 0.1

	for {
		x += dx
		y += dy

		if x <= 0.1 || x >= 0.9 {
			dx = -dx
		}
		if y <= 0.1 || y >= 0.9 {
			dy = -dy
		}

		if err := sendCommand(client, fmt.Sprintf("move %f %f", x, y)); err != nil {
			log.Fatal(err)
		}
		if err := sendCommand(client, "update"); err != nil {
			log.Fatal(err)
		}

		time.Sleep(100 * time.Millisecond)
	}
} 