package usecases

import (
	"context"
	"fmt"
	"req-smr/internal/services"

	"github.com/atomix/go-client/pkg/client/log"
)

func WatchRequests() error {
	fmt.Println("START:WATCH_REQUESTS")
	db, err := services.GetDatabase()
	if err != nil {
		fmt.Printf("ERROR:GET_DATABASE %s\n", err)
		return err
	}

	logPrimitive, err := db.GetLog(context.TODO(), "request-logs")
	if err != nil {
		fmt.Printf("ERROR:GET_LOG_REFERENCE %s\n", err)
		return err
	}

	channel := make(chan *log.Event)
	err = logPrimitive.Watch(context.Background(), channel)
	if err != nil {
		fmt.Printf("ERROR:WATCH_LOG %s\n", err)
		return err
	}
	go watchLogChanges(channel)

	return nil
}

func watchLogChanges(channel <-chan *log.Event) {
	for {
		fmt.Printf("START:WATCH_LOG_EVENT")
		event := <-channel
		fmt.Printf("FINISH:WATCH_LOG_EVENT %s\n", event)
	}
}
