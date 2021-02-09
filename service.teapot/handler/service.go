package handler

import (
	"context"
	"errors"
	"euterpe/service.teapot/dao"
	"euterpe/service.teapot/domain"
	teapotv1pb "euterpe/service.teapot/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Service struct {
	teapotv1pb.UnimplementedTeapotServer
	dao *dao.DAO
}

func New(dao *dao.DAO) *Service {
	return &Service{
		dao: dao,
	}
}

func (s *Service) GetById(ctx context.Context, req *teapotv1pb.GetByIdRequest) (*teapotv1pb.GetByIdResponse, error) {
	if req.Id == "" {
		return nil, errors.New("id should be non-empty string")
	}

	teapot, err := s.dao.FindByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &teapotv1pb.GetByIdResponse{
		// TODO: transform function to replace below
		Teapot: &teapotv1pb.Teapot{
			Id:          teapot.ID,
			Name:        teapot.Name,
			Radius:      teapot.Radius,
			Height:      teapot.Height,
			CreateTime:  timestamppb.New(teapot.CreateTime),
			UpdateTime:  timestamppb.New(teapot.UpdateTime),
			Description: "",
		},
	}, nil
}

func (s *Service) Create(ctx context.Context, req *teapotv1pb.CreateRequest) (*teapotv1pb.CreateResponse, error) {
	if req.Name == "" {
		return nil, errors.New("name should be non-empty string")
	}

	if req.Radius > 100 {
		return nil, errors.New("radius should be under 100")
	}

	if req.Height > 100 {
		return nil, errors.New("height should be under 100")
	}

	if len(req.Description) > 280 {
		return nil, errors.New("description should be tweet-length")
	}

	return &teapotv1pb.CreateResponse{}, nil
}

func (s *Service) CalculateVolume(ctx context.Context, req *teapotv1pb.CalculateVolumeRequest) (*teapotv1pb.CalculateVolumeResponse, error) {
	if req.Id == "" {
		return nil, errors.New("id should be non-empty string")
	}

	// TODO: Fetch teapot
	t := domain.Teapot{}

	return &teapotv1pb.CalculateVolumeResponse{Volume: t.Volume()}, nil
}
