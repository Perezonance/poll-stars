package server

import (
	"fmt"
	"net/http"

	"github.com/alexperez/poll-stars/storage"
	"go.uber.org/zap"
)

type server struct{
	s storage.Storage
	l *zap.Logger
}

//NewServer returns a server struct
func NewServer(s storage.Storage, l *zap.Logger) *server {
	return &server{s, l}
}

func (s *server) VoteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Entered VoteHandler")
}