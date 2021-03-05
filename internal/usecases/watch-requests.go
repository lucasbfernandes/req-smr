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

	go func() {
		for {
			fmt.Printf("START:WAITING_LOG_EVENT\n")
			event := <-channel

			request, err := services.ByteArrayToRequest(event.Entry.Value)
			if err != nil {
				fmt.Printf("ERROR:RECONSTRUCT_REQUEST %s\n", err)
				continue
			}

			_, err = services.ForwardRequest(request)
			if err != nil {
				fmt.Printf("ERROR:FORWARD_REQUEST %s\n", err)
				continue
			}

			fmt.Printf("SUCCEEDED:FORWARD_REQUEST %s\n", request)
		}
	}()

	return nil
}