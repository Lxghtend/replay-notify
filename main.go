package main

import (
	"log"
	"os/exec"

	"github.com/andreykaipov/goobs"
	"github.com/andreykaipov/goobs/api/events"
)

const (
	obsHost = "localhost:4455"
)

func sendSwayNotification() {
	title := "OBS Studio"
	body := "Replay buffer saved."

	cmd := exec.Command("notify-send",
		"-a", "OBS Studio",
		"-i", "media-record",
		"-u", "normal",
		"-i", "2500",
		title,
		body,
	)

	if err := cmd.Run(); err != nil {
		log.Printf("Failed to trigger notify-send: %v", err)
	}
}

func main() {
	client, err := goobs.New(obsHost)
	if err != nil {
		log.Fatalf("Failed to connect to OBS: %v", err)
	}
	defer client.Disconnect()

	for event := range client.IncomingEvents {
		switch event.(type) {
		case *events.ReplayBufferSaved:
			sendSwayNotification()
		}
	}
}
