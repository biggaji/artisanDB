package db

import "errors"

type DatabaseInterface interface {
	Set(key string, value string) (DbResponse, error)
	Get(key string) (DbResponse, error)
	All() []DbResponse
	Keys() []string
	Count() int
}

type Database struct {
	store map[string]string
}

func MakeDB() *Database {
	return &Database{
		store: map[string]string{},
	}
}

func (db *Database) Set(key string, value string) (DbResponse, error) {
	db.store[key] = value;
	return DbResponse{Key: key, Value: value} , nil
}

func (db *Database) Get(key string) (DbResponse, error) {
	value, found := db.store[key]

	response := DbResponse{Key: key, Value: value}

	if !found {
		return response, errors.New("key does not exist")
	}

	return response, nil
}

func (db *Database) All() []DbResponse {

	data := []DbResponse{}

	for key,value := range db.store {
		data = append(data, DbResponse{Key: key, Value: value})
	}

	return data
}

func (db *Database) Keys() []string {

	data := []string{}

	for key := range db.store {
		data = append(data, key)
	}

	return data
}

func (db *Database) Count() int {
	return len(db.store)
}