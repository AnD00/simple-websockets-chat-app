package dynamodb

import (
	"os"

	"github.com/pkg/errors"
)

const connectionsTableNameTemplate = "simple-websockets-chat-app-connections"

type Connection struct {
	ConnectionID string `dynamo:"connectionId,hash"`
	Room         string `dynamo:"room" index:"room-index,hash"`
}

func GetAllConnections(room string) ([]Connection, error) {
	db, err := connect()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	tableName := getConnectionsTableName()
	table := getTable(db, tableName)

	var results []Connection
	err = table.Get("room", room).Index("room-index").All(&results)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return results, nil
}

func PutConnection(connectionId string, room string) error {
	db, err := connect()
	if err != nil {
		return errors.WithStack(err)
	}

	tableName := getConnectionsTableName()
	table := getTable(db, tableName)

	putModel := Connection{
		ConnectionID: connectionId,
		Room:         room,
	}
	err = table.Put(putModel).Run()
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func DeleteConnection(connectionId string) error {
	db, err := connect()
	if err != nil {
		return errors.WithStack(err)
	}

	tableName := getConnectionsTableName()
	table := getTable(db, tableName)

	err = table.Delete("connectionId", connectionId).Run()
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func getConnectionsTableName() string {
	return connectionsTableNameTemplate + "-" + os.Getenv("STAGE_NAME")
}
