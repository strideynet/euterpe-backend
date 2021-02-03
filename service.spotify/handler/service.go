package handler

import (
	"context"
	spotifyv1pb "euterpe/service.spotify/proto"
	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
)

// TODO: State generation?
var state = "state_should_be_generated"

type Service struct {
	spotifyv1pb.UnimplementedSpotifyServer
	authenticator spotifyauth.Authenticator
}

func New() *Service {
	return &Service{
		authenticator: spotifyauth.New("http://localhost:8080/callback", spotifyauth.ScopeUserReadCurrentlyPlaying, spotifyauth.ScopeUserReadRecentlyPlayed),
	}
}

func (s *Service) GetAuthURL(ctx context.Context, req *spotifyv1pb.GetAuthURLRequest) (*spotifyv1pb.GetAuthURLResponse, error) {
	url := s.authenticator.AuthURL(state)

	return &spotifyv1pb.GetAuthURLResponse{Url: url}, nil
}

func (s *Service) RedeemCode(ctx context.Context, req *spotifyv1pb.RedeemCodeRequest) (*spotifyv1pb.RedeemCodeResponse, error) {
	token, err := s.authenticator.Exchange(ctx, req.Code)
	if err != nil {
		return nil, err
	}

	client := spotify.New(spotify.HTTPClientOpt(s.authenticator.Client(ctx, token)))

	user, err := client.CurrentUser(ctx)
	if err != nil {
		return nil, err
	}

	// store token

	return &spotifyv1pb.RedeemCodeResponse{User: &spotifyv1pb.User{
		DisplayName: user.DisplayName,
		Email:       user.Email,
		Id:          user.ID,
	}}, nil
}

func (s *Service) Healthy(_ context.Context) error {
	// Should probably check the underlying data-stores health?
	return nil
}
