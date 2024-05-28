package base

import (
	"demo/internal/data/models"
)

type UserRepository interface {
	BaseRepository[models.UserModel, int32]
}
