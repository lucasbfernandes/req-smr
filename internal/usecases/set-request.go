package usecases

import (
	"context"

	"github.com/atomix/go-client/pkg/client"
)

func SetRequest() error {
	atomix, err := client.New("localhost:5679")
	if err != nil {
		return err
	}

	db, err := atomix.GetDatabase(context.TODO(), "raft-database")
	if err != nil {
		return err
	}

	counter, err := db.GetCounter(context.TODO(), "my-counter")
	if err != nil {
		return err
	}

	err = counter.Set(context.TODO(), 10)
	if err != nil {
		return err
	}

	counter.Close(context.TODO())
	return nil
}
