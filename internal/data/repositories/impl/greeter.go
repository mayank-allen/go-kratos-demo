package impl

import (
	"context"
	data2 "demo/internal/data"
	"demo/internal/data/models"
	"github.com/go-kratos/kratos/v2/log"
)

// GreeterRepo is a Greater repo.
type GreeterRepo interface {
	Save(context context.Context, greeter *models.Greeter) (*models.Greeter, error)
	Update(context.Context, *models.Greeter) (*models.Greeter, error)
	FindByID(context.Context, int64) (*models.Greeter, error)
	ListByHello(context.Context, string) ([]*models.Greeter, error)
	ListAll(context.Context) ([]*models.Greeter, error)
}

type greeterRepo struct {
	data *data2.Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewGreeterRepo(data *data2.Data, logger log.Logger) GreeterRepo {
	return &greeterRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *greeterRepo) Save(ctx context.Context, g *models.Greeter) (*models.Greeter, error) {
	return g, nil
}

func (r *greeterRepo) Update(ctx context.Context, g *models.Greeter) (*models.Greeter, error) {
	return g, nil
}

func (r *greeterRepo) FindByID(context.Context, int64) (*models.Greeter, error) {
	return nil, nil
}

func (r *greeterRepo) ListByHello(context.Context, string) ([]*models.Greeter, error) {
	return nil, nil
}

func (r *greeterRepo) ListAll(context.Context) ([]*models.Greeter, error) {
	return nil, nil
}
