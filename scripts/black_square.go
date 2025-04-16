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
		"bgrect 0.2 0.2 0.8 0.8",
		"figure 0.3 0.3",  
		"figure 0.7 0.3",  
		"figure 0.5 0.5",  
		"figure 0.3 0.7",  
		"figure 0.7 0.7",
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