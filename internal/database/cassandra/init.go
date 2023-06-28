package cassandra

import (
	"fmt"

	"github.com/gocql/gocql"
)

var (
	DbClient *gocql.Session
)

func init() {
	var err error
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "post_service"
	DbClient, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}
    fmt.Println("\nCassandra is initialized! 👍")
}