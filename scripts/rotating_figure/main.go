package main

import (
	"fmt"
	"log"
	"math"
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

	centerX, centerY := 0.5, 0.5
	radius := 0.3
	angle := 0.0
	angularSpeed := 0.1

	for {
		x := centerX + radius*math.Cos(angle)
		y := centerY + radius*math.Sin(angle)

		if err := sendCommand(client, fmt.Sprintf("move %f %f", x, y)); err != nil {
			log.Fatal(err)
		}
		if err := sendCommand(client, "update"); err != nil {
			log.Fatal(err)
		}

		angle += angularSpeed
		if angle >= 2*math.Pi {
			angle = 0
		}

		time.Sleep(100 * time.Millisecond)
	}
} 