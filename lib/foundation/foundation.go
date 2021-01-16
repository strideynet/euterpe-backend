// Package foundation provides a framework for a typical microservice, providing health checking and monitoring
// whilst possible to write a service without foundation, it is not recommended as you are likely to diverge from
// organisation standards.
package foundation

import (
	"fmt"
	"gmm/lib/health"
	"gmm/lib/log"
	"gmm/lib/log/grpcdebug"
	"gmm/lib/reqid"
	"gmm/lib/sentry"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

var logger = log.Package("foundation")

// Foundation allows setup for basic dependencies of most services e.g health checking
type Foundation struct {
	grpcS   *grpc.Server
	healthS *health.Service
}

// New returns a new instance of foundation, with sub-dependencies set up
func New() *Foundation {
	return &Foundation{
		grpcS:   setupGRPCServer(),
		healthS: health.New(),
	}
}

func setupGRPCServer() *grpc.Server {
	return grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			reqid.UnaryInterceptor,
			grpcdebug.UnaryInterceptor,
			grpc_prometheus.UnaryServerInterceptor,
			sentry.UnaryInterceptor,
		),

		grpc.ChainStreamInterceptor(
			reqid.StreamInterceptor,
			grpc_prometheus.StreamServerInterceptor,
			sentry.StreamInterceptor,
		),
	)
}

// GRPCRegistry returns the grpc server so the microservice can register its services.
func (f *Foundation) GRPCRegistry() grpc.ServiceRegistrar {
	return f.grpcS
}

// RegisterHealthCheck passes another health check to be called through to the underlying health service
func (f *Foundation) RegisterHealthCheck(checker health.Checker) {
	f.healthS.Register(checker)
}

// TODO: Make these configurable
const grpcPort = 8080
const prometheusPort = 9090
const shutdownWaitSeconds = 2

// Run starts the main service as well as several debugging/healthcheck services such as prometheus.
// It should only be called once and should be called once any desired services have been attached.
func (f *Foundation) Run() error {
	logger.Info("foundation Run() called")
	// Register meta services used for debugging/healthchecking
	grpc_health_v1.RegisterHealthServer(f.grpcS, f.healthS)
	reflection.Register(f.grpcS)

	// Channel to synchronise waiting for end.
	shut := make(chan error, 1)

	// Run main grpc listener
	go func() {
		listen, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
		if err != nil {
			logger.Fatal(err.Error())
		}

		logger.Info("starting grpc listener", zap.Int("port", grpcPort))
		err = f.grpcS.Serve(listen)

		shut <- err
	}()

	// Run prometheus exporter
	go func() {
		logger.Info("starting prometheus listener", zap.Int("port", prometheusPort))

		err := http.ListenAndServe(fmt.Sprintf(":%d", prometheusPort), promhttp.Handler())
		shut <- err
	}()

	// Run shutdown await
	go func() {
		cancelChan := make(chan os.Signal, 1)
		signal.Notify(cancelChan, syscall.SIGTERM, syscall.SIGINT)

		// Wait for shutdown
		<-cancelChan
		logger.Warn("caught sigterm, starting shutdown")

		// Mark service unhealthy to loadbalancers.
		f.healthS.ShuttingDown()

		// Give loadbalancers time to shift traffic
		logger.Info("waiting before shutdown", zap.Int("seconds", shutdownWaitSeconds))
		time.Sleep(shutdownWaitSeconds * time.Second)

		// Actually stop serving!
		logger.Info("sending final stop signals to services")
		f.grpcS.GracefulStop()

		close(shut)
	}()

	return <-shut
}
