// The teapot service exists to test new ideas and demonstrate concepts
package main

import (
	"euterpe/lib/foundation"
	"euterpe/service.teapot/handler"
	"euterpe/service.teapot/proto"
)

// Run starts the service
// TODO: How do we more cleanly handle context here? Should we inherit the cobra context or should foundation return one?
func main() {
	f := foundation.New()

	teapotv1pb.RegisterTeapotServer(f.GRPCRegistry(), handler.New())

	f.Run()
}
