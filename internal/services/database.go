package services

import (
	"context"
	"sync"

	"github.com/atomix/go-client/pkg/client"
)

var dbInstance *client.Database
var syncError error
var once sync.Once

func GetDatabase() (*client.Database, error) {
	once.Do(func() {
		atomix, err := client.New("atomix-controller.default.svc.cluster.local:5679")
		if err != nil {
			syncError = err
		}

		dbInstance, err = atomix.GetDatabase(context.TODO(), "raft-database")
		if err != nil {
			syncError = err
		}
	})

	return dbInstance, syncError
}
