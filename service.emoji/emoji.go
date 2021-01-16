package emoji

import (
	"gmm/lib/foundation"
	v1handler "gmm/service.emoji/handler/v1"
	v1proto "gmm/service.emoji/proto/v1"
	"gmm/service.emoji/repo/memory"

	"github.com/spf13/cobra"
)

// CMD defines the cobra command for the service
var CMD = &cobra.Command{
	Use:   "service.emoji",
	Short: "Runs the emoji service",
	RunE: func(cmd *cobra.Command, args []string) error {
		return Run()
	},
}

// Run starts the service
// TODO: How do we more cleanly handle context here? Should we inherit the cobra context or should foundation return one?
func Run() error {
	f := foundation.New()

	memStore := memory.New()
	emojiSrv := v1handler.New(memStore)

	f.RegisterHealthCheck(emojiSrv)
	v1proto.RegisterEmojiServiceServer(f.GRPCRegistry(), emojiSrv)

	return f.Run()
}
