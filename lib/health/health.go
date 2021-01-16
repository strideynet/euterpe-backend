// Package health provides a **simple** grpc_health_v1 compatible service that
// has a shutting down state and calls onto other registered services when a health check is triggered.
package health

import (
	"context"
	"gmm/lib/log"
	"sync"

	"google.golang.org/grpc/health/grpc_health_v1"
)

var logger = log.Package("health")

// Checker represents another service that we may query when wanting to check if our system is healthy
type Checker interface {
	Healthy(ctx context.Context) error
}

// Service is a simple implementation of the grpc health v1 service
type Service struct {
	mu           sync.Mutex
	shuttingDown bool

	checkers []Checker

	grpc_health_v1.UnimplementedHealthServer
}

// Check returns the overall health of the microservice, and returns failure when shut down has begun
func (s *Service) Check(
	ctx context.Context, _ *grpc_health_v1.HealthCheckRequest,
) (*grpc_health_v1.HealthCheckResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.shuttingDown {
		res := grpc_health_v1.HealthCheckResponse{
			Status: grpc_health_v1.HealthCheckResponse_NOT_SERVING,
		}

		return &res, nil
	}

	for _, checker := range s.checkers {
		err := checker.Healthy(ctx)
		if err != nil {
			res := grpc_health_v1.HealthCheckResponse{
				Status: grpc_health_v1.HealthCheckResponse_NOT_SERVING,
			}

			return &res, nil
		}
	}

	res := grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	}

	return &res, nil
}

// Register adds another function that should be called when a client requests a Check
func (s *Service) Register(c Checker) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.checkers = append(s.checkers, c)
}

// ShuttingDown marks the system as shutting down, and will start failing ready checks
func (s *Service) ShuttingDown() {
	s.mu.Lock()
	defer s.mu.Unlock()

	logger.Info("marking service as unhealthy")

	s.shuttingDown = true
}

// New returns a new instance of the health service
func New() *Service {
	return &Service{}
}
