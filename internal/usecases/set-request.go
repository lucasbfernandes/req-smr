package usecases

import (
	"context"
	"fmt"

	"github.com/atomix/go-client/pkg/client"
)

func SetRequest() error {
	fmt.Println("start:creating_client")
	atomix, err := client.New("atomix-controller.default.svc.cluster.local:5679")
	if err != nil {
		return err
	}
	fmt.Println("finish:creating_client")

	fmt.Println("start:get_database")
	db, err := atomix.GetDatabase(context.TODO(), "raft-database")
	if err != nil {
		return err
	}
	fmt.Println("finish:get_database")

	fmt.Println("start:get_counter")
	counter, err := db.GetCounter(context.TODO(), "my-counter")
	if err != nil {
		return err
	}
	fmt.Println("finish:get_counter")

	fmt.Println("start:set_counter")
	err = counter.Set(context.TODO(), 10)
	if err != nil {
		return err
	}
	fmt.Println("finish:set_counter")

	counter.Close(context.TODO())
	return nil
}
