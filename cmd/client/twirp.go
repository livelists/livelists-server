package client

import (
	"fmt"
	pb "github.com/livelists/livelist-server/contracts/channel"
	"github.com/livelists/livelist-server/pkg/services"
	"log"
	"net/http"
)

const serverAddr = ":8080"

func StartTwirpRPC() {
	fmt.Println("fg")
	channelSVC := services.ChannelService{}
	channelHandler := pb.NewChannelServiceServer(&channelSVC)
	mux := http.NewServeMux()
	mux.Handle(channelHandler.PathPrefix(), channelHandler)

	log.Printf("RPC listening %s", serverAddr)

	err := http.ListenAndServe(serverAddr, mux)
	if err != nil {
		log.Fatal("Twirp server started", err)
	}
}
