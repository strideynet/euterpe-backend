// The teapot service exists to test new ideas and demonstrate concepts
package main

import (
	"euterpe/lib/db"
	"euterpe/lib/foundation"
	"euterpe/lib/log"
	"euterpe/service.teapot/dao"
	"euterpe/service.teapot/handler"
	"euterpe/service.teapot/proto"
	"go.uber.org/zap"
)

const serviceName = "teapot"

var l = log.Package("service.teapot:main")

func main() {
	f := foundation.New()

	pg, err := db.New(db.ServiceSettingsFromEnv(serviceName))
	if err != nil {
		l.Fatal("failed to configure db", zap.Error(err))
	}
	h := handler.New(dao.New(pg))

	teapotv1pb.RegisterTeapotServer(f.GRPCRegistry(), h)

	f.Run()
}
