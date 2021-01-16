package v1

import (
	"context"
	"gmm/service.emoji/domain"
	v1pb "gmm/service.emoji/proto/v1"
	"gmm/service.emoji/repo"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	emojiRepo repo.Repo
	v1pb.UnimplementedEmojiServiceServer
}

func New(repo repo.Repo) *Service {
	return &Service{
		emojiRepo: repo,
	}
}

func (s *Service) List(ctx context.Context, _ *v1pb.ListRequest) (*v1pb.ListResponse, error) {
	emojis, err := s.emojiRepo.List(ctx)
	if err != nil {
		return nil, err
	}

	var res []*v1pb.Emoji
	for _, emoji := range emojis {
		res = append(res, emoji.ToProto())
	}

	return &v1pb.ListResponse{Emojis: res}, nil
}

func (s *Service) Create(ctx context.Context, request *v1pb.CreateRequest) (*v1pb.CreateResponse, error) {
	if request.GetEmoji() == nil {
		return nil, status.Error(codes.InvalidArgument, "emoji to create must be provided")
	}

	emoji, err := domain.ProtoToEmoji(request.GetEmoji())
	if err != nil {
		return nil, err
	}

	err = s.emojiRepo.Create(ctx, emoji)
	if err != nil {
		return nil, err
	}

	return &v1pb.CreateResponse{}, nil
}

func (s *Service) Delete(ctx context.Context, request *v1pb.DeleteRequest) (*v1pb.DeleteResponse, error) {
	id := request.GetId()
	if id == "" {
		return nil, status.Error(codes.InvalidArgument, "id must be non-empty string")
	}

	err := s.emojiRepo.Delete(ctx, id)
	if err != nil {
		return nil, err
	}

	return &v1pb.DeleteResponse{}, nil
}

func (s *Service) FindById(ctx context.Context, request *v1pb.FindByIdRequest) (*v1pb.FindByIdResponse, error) {
	id := request.GetId()
	if id == "" {
		return nil, status.Error(codes.InvalidArgument, "id must be non-empty string")
	}

	emoji, err := s.emojiRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if emoji == nil {
		return nil, status.Error(codes.NotFound, "emoji by that id not found")
	}

	return &v1pb.FindByIdResponse{Emoji: emoji.ToProto()}, nil
}

func (s *Service) Healthy(_ context.Context) error {
	// Should probably check the underlying data-stores health?
	return nil
}
