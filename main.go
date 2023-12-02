package main

import (
	"os"
	"os/signal"
	"syscall"
	"twitter-oracle-backend/internal/log"
	"twitter-oracle-backend/logic/blockchain/events"
	"twitter-oracle-backend/logic/config"
)

func main() {
	_, err := events.NewEvents(config.Get().Event)
	if err != nil {
		panic(err)
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-done

	log.GetLogger().Info("gracefully shutdown...")
}
