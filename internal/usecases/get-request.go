package usecases

import (
	"context"
	"fmt"
	"req-smr/internal/services"
)

func GetRequest() (*int64, error) {
	fmt.Println("START:GET_REQUEST")
	db, err := services.GetDatabase()
	if err != nil {
		fmt.Printf("ERROR:GET_DATABASE %s\n", err)
		return nil, err
	}

	counter, err := db.GetCounter(context.TODO(), "my-counter")
	if err != nil {
		fmt.Printf("ERROR:GET_COUNTER_REFERENCE %s\n", err)
		return nil, err
	}
	defer counter.Close(context.TODO())

	count, err := counter.Get(context.TODO())
	if err != nil {
		fmt.Printf("ERROR:GET_COUNTER_VALUE %s\n", err)
		return nil, err
	}

	fmt.Printf("FINISH:GET_REQUEST %d\n", count)
	return &count, nil
}
