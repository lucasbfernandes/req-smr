package usecases

import (
	"context"

	"req-smr/internal/services"
)

func SetRequest() error {
	db, err := services.GetDatabase()
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
