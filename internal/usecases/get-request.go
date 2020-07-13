package usecases

import (
	"context"
	"fmt"
	"req-smr/internal/services"
)

func GetRequest() (*int64, error) {
	db, err := services.GetDatabase()
	if err != nil {
		return nil, err
	}

	counter, err := db.GetCounter(context.TODO(), "my-counter")
	if err != nil {
		return nil, err
	}
	defer counter.Close(context.TODO())

	count, err := counter.Get(context.TODO())
	if err != nil {
		return nil, err
	}

	fmt.Printf("COUNTER VALUE: %d\n", count)
	return &count, nil
}
