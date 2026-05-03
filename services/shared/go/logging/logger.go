package logging

import "log"

type Logger struct{}

func (Logger) Info(message string) {
	log.Printf("[INFO] %s", message)
}
