package service

import (
	"context"

	"net/http"

	"github.com/dapr/go-sdk/actor"
	"github.com/dapr/go-sdk/service/common"
	daprd "github.com/dapr/go-sdk/service/http"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)

type Service struct {
	Svc     common.Service
	Mux     *chi.Mux
	Address string
}

func NewService(address string, mux *chi.Mux) *Service {
	svc := daprd.NewServiceWithMux(address, mux)
	return &Service{svc, mux, address}
}

func (s *Service) Stop() error {
	return s.Svc.Stop()
}

func (s *Service) ServeInvoke(pattern string, handler func(ctx context.Context, in *common.InvocationEvent) (out *common.Content, err error)) {
	if err := s.Svc.AddServiceInvocationHandler(pattern, handler); err != nil {
		log.Fatal().Err(err).Msg("Dapr service; cannot register invocation")
	}
}

func (s *Service) ServeHealth() {
	if err := s.Svc.AddHealthCheckHandler("healthz", func(context.Context) error {
		log.Debug().Msg("health check")
		return nil
	}); err != nil {
		log.Fatal().Err(err).Msg("Dapr service; cannot register health handler")
	}
}

func (s *Service) Serve(pattern string, registerer func() http.HandlerFunc) {
	s.Mux.Handle(pattern, registerer())
}

func (s *Service) ServeEvent(subscription *common.Subscription, handler func(ctx context.Context, e *common.TopicEvent) (retry bool, err error)) {
	if err := s.Svc.AddTopicEventHandler(subscription, handler); err != nil {
		log.Fatal().Err(err).Msg("Dapr service; cannot register event")
	}
}

func (s *Service) RegisterActorImplFactoryContext(f actor.FactoryContext) {
	s.Svc.RegisterActorImplFactoryContext(f)
}

func (s *Service) Start() {
	cont := make(chan bool)
	go func() {
		log.Info().Str("address", s.Address).Msg("started DAPR service")

		cont <- true

		if err := s.Svc.Start(); err != nil && err != http.ErrServerClosed {
			log.Panic().Err(err).Msg("PUBSUB Server stopped")
		}
	}()

	<-cont
}
