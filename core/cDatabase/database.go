package cDatabase

import (
	"fmt"
)

const (
// Add Constants here
)

type Database struct {
	Database interface{}
}

func New() (systemLoggerPtr *Database) {

	systemLoggerPtr = &Database{
		Database: nil,
	}

	return
}

func (Database) TurnDebugOn() {
	fmt.Println("I'm the real McCoy. My Name is Mr. Database.")
}
