package services

import (
	"context"
	"fmt"
	"sync"

	"github.com/atomix/go-client/pkg/client"
)

var dbInstance *client.Database
var syncError error
var once sync.Once

func GetDatabase() (*client.Database, error) {
	fmt.Println("START:GET_DATABASE")
	once.Do(func() {
		fmt.Println("START:GET_ATOMIX_CONNECTION")
		atomix, err := client.New("atomix-controller.default.svc.cluster.local:5679")
		if err != nil {
			fmt.Printf("ERROR:GET_ATOMIX_CONNECTION %s\n", err)
			syncError = err
		}
		fmt.Println("FINISH:GET_ATOMIX_CONNECTION")

		fmt.Println("START:GET_DATABASE_CONNECTION")
		dbInstance, err = atomix.GetDatabase(context.TODO(), "raft-database")
		if err != nil {
			fmt.Printf("ERROR:GET_DATABASE_CONNECTION %s\n", err)
			syncError = err
		}
		fmt.Println("FINISH:GET_DATABASE_CONNECTION")
	})

	fmt.Printf("FINISH:GET_DATABASE %+v\n", dbInstance)
	return dbInstance, syncError
}
