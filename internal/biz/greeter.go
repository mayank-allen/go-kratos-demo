package biz

import (
	"context"
	"demo/internal/data/models"
	"demo/internal/data/repositories/impl"

	v1 "demo/api/helloworld/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// GreeterUsecase is a Greeter usecase.
type GreeterUsecase struct {
	repo impl.GreeterRepo
	log  *log.Helper
}

// NewGreeterUsecase new a Greeter usecase.
func NewGreeterUsecase(repo impl.GreeterRepo, logger log.Logger) *GreeterUsecase {
	return &GreeterUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *GreeterUsecase) CreateGreeter(ctx context.Context, g *models.Greeter) (*models.Greeter, error) {
	uc.log.WithContext(ctx).Infof("CreateGreeter: %v", *g)
	return uc.repo.Save(ctx, g)
}
