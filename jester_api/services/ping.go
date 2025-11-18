package services

type PingService interface {
	Message() string
}

type pingService struct{}

func NewPingService() PingService {
	return &pingService{}
}

func (s *pingService) Message() string {
	return "pong"
}