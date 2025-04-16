package main

import (
	"log"
	"net/http"
	"strings"
)

func main() {
	client := &http.Client{}
	commands := []string{
		"white",
		"bgrect 0.25 0.25 0.75 0.75",
		"figure 0.5 0.5",
		"green",
		"figure 0.6 0.6",
		"update",
	}

	for _, cmd := range commands {
		req, err := http.NewRequest("POST", "http://localhost:17000", strings.NewReader(cmd))
		if err != nil {
			log.Fatal(err)
		}
		
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		resp.Body.Close()
	}
} 