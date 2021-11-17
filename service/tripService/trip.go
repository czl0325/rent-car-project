package trip

import (
	"context"
	trippb "service/proto/gen/go"
)

// Service implements trip service.
type Service struct {
	*trippb.UnimplementedTripServiceServer
}

func (s *Service) GetTrip(c context.Context, request *trippb.GetTripRequest) (response *trippb.Trip, err error) {
	trip := &trippb.Trip{
		Start:       "abc",
		StartPos:    &trippb.Location{
			Latitude:  24,
			Longitude: 118,
		},
		End:         "",
		EndPos:      &trippb.Location{
			Latitude:  24,
			Longitude: 118,
		},
		DurationSec: 10,
		FeeCent:     10000,
		Status:      trippb.TripStatus_FINISHED,
		Id:          request.Id,
	}
	return trip, nil
}
