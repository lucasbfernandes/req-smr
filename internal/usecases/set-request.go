package usecases

import (
	"context"
	"fmt"

	"req-smr/internal/services"
)

func SetRequest(rawData []byte) error {
	fmt.Println("START:SET_REQUEST")
	db, err := services.GetDatabase()
	if err != nil {
		fmt.Printf("ERROR:GET_DATABASE %s\n", err)
		return err
	}

	log, err := db.GetLog(context.TODO(), "request-logs")
	if err != nil {
		fmt.Printf("ERROR:GET_LOG_REFERENCE %s\n", err)
		return err
	}

	_, err = log.Append(context.TODO(), rawData)
	if err != nil {
		fmt.Printf("ERROR:APPEND_LOG %s\n", err)
		return err
	}

	log.Close(context.TODO())
	fmt.Printf("FINISH:SET_APPEND INDEX: %d")
	return nil
}
