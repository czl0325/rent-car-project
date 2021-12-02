package rental

import (
	"context"
	rentalpb "service/rental/api/gen/v1"
)

type Service struct {
	*rentalpb.UnimplementedTripServiceServer
}

func (s *Service) CreateTrip(c context.Context, request *rentalpb.CreateTripRequest) (*rentalpb.Trip, error) {
	return &rentalpb.Trip{
		Id:         "1",
		AccountId:  "1",
		CarId:      "1",
		Start:      nil,
		Current:    nil,
		End:        nil,
		Status:     0,
		IdentityId: "1",
	}, nil
}

func (s *Service) GetTrip(c context.Context, request *rentalpb.GetTripRequest) (*rentalpb.Trip, error) {
	return &rentalpb.Trip{
		Id:         request.Id,
		AccountId:  "1",
		CarId:      "1",
		Start:      nil,
		Current:    nil,
		End:        nil,
		Status:     0,
		IdentityId: "1",
	}, nil
}
