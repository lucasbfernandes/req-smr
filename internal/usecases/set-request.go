package usecases

import (
	"context"
	"fmt"
	"req-smr/internal/models"

	"req-smr/internal/services"
)

func SetRequest(request *models.Request) error {
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

	serializedRequest, err := services.RequestToByteArray(request)
	if err != nil {
		fmt.Printf("ERROR:SERIALIZE_REQUEST %s\n", err)
		return err
	}
	fmt.Printf("GET:SERIALIZED_REQUEST %s\n", serializedRequest)

	_, err = log.Append(context.TODO(), serializedRequest)
	if err != nil {
		fmt.Printf("ERROR:APPEND_LOG %s\n", err)
		return err
	}

	log.Close(context.TODO())
	fmt.Println("FINISH:SET_APPEND")
	return nil
}
