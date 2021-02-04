package handler

import (
	"context"
	"errors"
	teapotv1pb "euterpe/service.teapot/proto"
)

type Service struct {
	teapotv1pb.UnimplementedTeapotServer
}

func New() *Service {
	return &Service{}
}

func (s *Service) GetById(ctx context.Context, req *teapotv1pb.GetByIdRequest) (*teapotv1pb.GetByIdResponse, error) {
	if req.Id == "" {
		return nil, errors.New("id should be non-empty string")
	}
	return &teapotv1pb.GetByIdResponse{}, nil
}

func (s *Service) Create(ctx context.Context, req *teapotv1pb.CreateRequest) (*teapotv1pb.CreateResponse, error) {
	if req.Name == "" {
		return nil, errors.New("name should be non-empty string")
	}

	if req.Capacity > 100 {
		return nil, errors.New("capacity should be under 100")
	}

	if len(req.Description) > 280 {
		return nil, errors.New("description should be tweet-length")
	}

	return &teapotv1pb.CreateResponse{}, nil
}
