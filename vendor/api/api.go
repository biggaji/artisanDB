package api

import (
	"db"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type ArtisanServer struct {
	db db.DatabaseInterface
}

func MakeDBServer(db db.DatabaseInterface) *ArtisanServer {
	return &ArtisanServer{db}
}

func (s *ArtisanServer) Start() {
	err := rpc.Register(s)
	if err != nil {
		log.Fatal("Server could not be started", err)
	
	}

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":9090")

	log.Println("Starting server at port 9090")

	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("Could not start server: ", err)
	}

}

func (s *ArtisanServer) Set(payload, response *string) error {
	s.db.Set("hello", "world")

	*response = "world"

	return nil
}