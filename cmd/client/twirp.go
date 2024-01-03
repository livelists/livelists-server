package client

import (
	"context"
	channel_pb "github.com/livelists/livelist-server/contracts/channel"
	participant_pb "github.com/livelists/livelist-server/contracts/participant"
	"github.com/livelists/livelist-server/pkg/services"
	"github.com/livelists/livelist-server/pkg/services/accessToken"
	"github.com/livelists/livelist-server/pkg/websocket"
	"log"
	"net/http"
	"strconv"
)

func withAuth(base http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		authToken := r.Header.Get("Authorization")

		if authToken == "" {
			ctx = context.WithValue(ctx, "isTokenValid", false)
		} else {
			token := accessToken.AccessToken{}

			isValid, err := token.Parse(authToken)

			if err != nil {
				ctx = context.WithValue(ctx, "isTokenValid", false)
			} else if isValid && token.IsServiceRoot() {
				ctx = context.WithValue(ctx, "isTokenValid", true)
			}
		}

		r = r.WithContext(ctx)

		base.ServeHTTP(w, r)
	})
}

func StartTwirpRPC(port uint) {
	channelSVC := services.ChannelService{}
	channelHandler := channel_pb.NewChannelServiceServer(&channelSVC)
	participantSVC := services.ParticipantService{}
	participantSVC.WS = websocket.WsRoom{}
	participantHandler := participant_pb.NewParticipantServiceServer(participantSVC)

	mux := http.NewServeMux()

	mux.Handle(channelHandler.PathPrefix(), withAuth(channelHandler))
	mux.Handle(participantHandler.PathPrefix(), withAuth(participantHandler))

	addr := ":" + strconv.Itoa(int(port))
	log.Printf("Admin TwirpRPC listening %s", addr)

	err := http.ListenAndServe(addr, mux)
	if err != nil {
		log.Fatal("Twirp server started", err)
	}
}
