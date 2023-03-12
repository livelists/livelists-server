package client

import (
	channel_pb "github.com/livelists/livelist-server/contracts/channel"
	participant_pb "github.com/livelists/livelist-server/contracts/participant"
	"github.com/livelists/livelist-server/pkg/services"
	"log"
	"net/http"
)

const serverAddr = ":8080"

func StartTwirpRPC() {
	channelSVC := services.ChannelService{}
	channelHandler := channel_pb.NewChannelServiceServer(&channelSVC)
	participantSVC := services.ParticipantService{}
	participantHandler := participant_pb.NewParticipantServiceServer(&participantSVC)

	mux := http.NewServeMux()
	mux.Handle(channelHandler.PathPrefix(), channelHandler)
	mux.Handle(participantHandler.PathPrefix(), participantHandler)

	log.Printf("RPC listening %s", serverAddr)

	err := http.ListenAndServe(serverAddr, mux)
	if err != nil {
		log.Fatal("Twirp server started", err)
	}
}
