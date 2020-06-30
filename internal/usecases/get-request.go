package usecases

import (
	"context"
	"fmt"

	"github.com/atomix/go-client/pkg/client"
)

func GetRequest() (*int64, error) {
	atomix, err := client.New("localhost:5679")
	if err != nil {
		return nil, err
	}

	db, err := atomix.GetDatabase(context.TODO(), "raft-database")
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
