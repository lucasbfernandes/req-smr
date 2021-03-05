package services

import (
	"context"
	"fmt"
)

func GetRequests() (string, error) {
	fmt.Println("START:GET_REQUEST")
	db, err := GetDatabase()
	if err != nil {
		fmt.Printf("ERROR:GET_DATABASE %s\n", err)
		return "", err
	}

	log, err := db.GetLog(context.TODO(), "request-logs")
	if err != nil {
		fmt.Printf("ERROR:GET_COUNTER_REFERENCE %s\n", err)
		return "", err
	}
	defer log.Close(context.TODO())

	entry, err := log.LastEntry(context.TODO())
	if err != nil {
		fmt.Printf("ERROR:GET_LOG_LAST_ENTRY %s\n", err)
		return "", err
	}

	entryValue := string(entry.Value)
	fmt.Printf("FINISH:GET_LAST_ENTRY %s\n", entryValue)
	return entryValue, nil
}
