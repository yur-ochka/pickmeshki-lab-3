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

type Position struct {
	x, y float64
}

func (p *Position) move(dx, dy float64) {
	p.x += dx
	p.y += dy
}

func sendMoveAndUpdate(client *http.Client, pos Position) error {
	if err := sendCommand(client, fmt.Sprintf("move %f %f", pos.x, pos.y)); err != nil {
		return err
	}
	return sendCommand(client, "update")
}

func main() {
	client := &http.Client{}

	startPos := 0.25
	finishPos := 0.65
	step := 0.1
	interval := time.Second

	pos := Position{startPos, startPos}

	initCommands := []string{
		"green",
		fmt.Sprintf("figure %f %f", pos.x, pos.y),
		"update",
	}

	for _, cmd := range initCommands {
		if err := sendCommand(client, cmd); err != nil {
			log.Fatal(err)
		}
	}

	time.Sleep(interval)

	direction := "right" 
	moves := map[string]struct {
		condition func() bool
		delta     Position
		next      string
	}{
		"right": {
			condition: func() bool { return pos.x < finishPos },
			delta:     Position{step, 0},
			next:      "down",
		},
		"down": {
			condition: func() bool { return pos.y < finishPos },
			delta:     Position{0, step},
			next:      "left",
		},
		"left": {
			condition: func() bool { return pos.x > startPos },
			delta:     Position{-step, 0},
			next:      "up",
		},
		"up": {
			condition: func() bool { return pos.y > startPos },
			delta:     Position{0, -step},
			next:      "right",
		},
	}

	for {
		move := moves[direction]
		if move.condition() {
			pos.move(move.delta.x, move.delta.y)
			if err := sendMoveAndUpdate(client, pos); err != nil {
				log.Fatal(err)
			}
		} else {
			direction = move.next
		}

		time.Sleep(interval)
	}
}