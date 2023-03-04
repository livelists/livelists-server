package client

import (
	"net/http"

	"github.com/example/internal/haberdasherserver"
	"github.com/example/rpc/haberdasher"
)

func StartTwirp() {
	server := &haberdasherserver.Server{} // implements Haberdasher interface
	twirpHandler := haberdasher.NewHaberdasherServer(server)

	http.ListenAndServe(":80", twirpHandler)
}

func NewTwirpServer()
