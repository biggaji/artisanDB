package db

import "fmt"


type DbResponse struct {
	Key string
	Value string
}


func (dbr DbResponse) String() string {
	return fmt.Sprintf("%s => %s", dbr.Key, dbr.Value)
}

