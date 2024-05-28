package impl

import (
	"demo/internal/data"
	"demo/internal/data/models"
	"demo/internal/data/repositories/base"
)

type UserRepositoryImpl struct {
	data *data.Data
}

func NewUserRepository(data *data.Data) base.UserRepository {
	return &UserRepositoryImpl{data}
}

func (userRepositoryImpl *UserRepositoryImpl) Create(user *models.UserModel) (*models.UserModel, error) {
	result := userRepositoryImpl.data.Db.Save(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (userRepositoryImpl *UserRepositoryImpl) Update(user *models.UserModel) (*models.UserModel, error) {
	return userRepositoryImpl.Create(user)

}

func (userRepositoryImpl *UserRepositoryImpl) DeleteById(id int32) error {
	result := userRepositoryImpl.data.Db.Delete(&models.UserModel{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (userRepositoryImpl *UserRepositoryImpl) FindById(id int32) (*models.UserModel, error) {
	userModel := &models.UserModel{}
	result := userRepositoryImpl.data.Db.Unscoped().First(userModel, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return userModel, nil
}

func (userRepositoryImpl *UserRepositoryImpl) FindAll() ([]*models.UserModel, error) {
	var userModels []*models.UserModel
	result := userRepositoryImpl.data.Db.Unscoped().Find(&userModels)
	if result.Error != nil {
		return nil, result.Error
	}
	return userModels, nil
}
