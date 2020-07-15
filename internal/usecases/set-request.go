package usecases

import (
	"context"
	"fmt"

	"req-smr/internal/services"
)

func SetRequest() error {
	fmt.Println("START:SET_REQUEST")
	db, err := services.GetDatabase()
	if err != nil {
		fmt.Printf("ERROR:GET_DATABASE %s\n", err)
		return err
	}

	counter, err := db.GetCounter(context.TODO(), "my-counter")
	if err != nil {
		fmt.Printf("ERROR:GET_COUNTER_REFERENCE %s\n", err)
		return err
	}

	err = counter.Set(context.TODO(), 10)
	if err != nil {
		fmt.Printf("ERROR:SET_COUNTER %s\n", err)
		return err
	}

	counter.Close(context.TODO())
	fmt.Println("FINISH:SET_REQUEST")
	return nil
}
