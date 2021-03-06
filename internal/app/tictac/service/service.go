package service

import (
	"context"

	"github.com/cage1016/ms-sample/internal/app/tictac/model"
	"github.com/go-kit/kit/log"
)

// Middleware describes a service (as opposed to endpoint) middleware.
type Middleware func(TictacService) TictacService

// Service describes a service that adds things together
// Implement yor service methods methods.
// e.x: Foo(ctx context.Context, s string)(rs string, err error)
type TictacService interface {
	// [method=post,expose=true]
	Tic(ctx context.Context) (err error)
	// [method=get,expose=true]
	Tac(ctx context.Context) (res int64, err error)
}

// the concrete implementation of service interface
type stubTictacService struct {
	logger log.Logger
	repo   model.TicTacRespository
}

// New return a new instance of the service.
// If you want to add service middleware this is the place to put them.
func New(repo model.TicTacRespository, logger log.Logger) (s TictacService) {
	var svc TictacService
	{
		svc = &stubTictacService{logger: logger, repo: repo}
		svc = LoggingMiddleware(logger)(svc)
	}
	return svc
}

// Implement the business logic of Tic
func (ti *stubTictacService) Tic(ctx context.Context) (err error) {
	return ti.repo.Add(ctx)
}

// Implement the business logic of Tac
func (ti *stubTictacService) Tac(ctx context.Context) (res int64, err error) {
	return ti.repo.Get(ctx)
}
