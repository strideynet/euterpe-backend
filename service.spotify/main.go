package main

import (
	"euterpe/lib/foundation"
	v1 "euterpe/service.spotify/handler"
	spotifyv1pb "euterpe/service.spotify/proto"
)

// Run starts the service
// TODO: How do we more cleanly handle context here? Should we inherit the cobra context or should foundation return one?
func main() {
	f := foundation.New()

	spotifyv1pb.RegisterSpotifyServer(f.GRPCRegistry(), v1.New())

	f.Run()
}
