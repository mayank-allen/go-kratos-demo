package biz

import (
	v1 "demo/api/user/v1"
	"demo/internal/data/models"
	"demo/internal/data/repositories/base"
)

type UserHandler struct {
	userRepository base.UserRepository
}

func NewUserHandler(repo base.UserRepository) *UserHandler {
	return &UserHandler{repo}
}
func convertFrom(userDto *v1.UserDto) models.UserModel {
	return models.UserModel{
		Name:             userDto.Name,
		Age:              userDto.Age,
		Email:            userDto.Email,
		CurrentAddress:   userDto.GetCurrentAddress(),
		PermanentAddress: userDto.GetPermanentAddress(),
	}
}
func convertTo(userModel models.UserModel) *v1.UserDto {
	return &v1.UserDto{
		Id:               int32(userModel.ID),
		Name:             userModel.Name,
		Age:              userModel.Age,
		Email:            userModel.Email,
		CurrentAddress:   userModel.CurrentAddress,
		PermanentAddress: userModel.PermanentAddress,
	}
}
func update(userDto *v1.UserDto, userModel *models.UserModel) {
	userModel.Name = userDto.Name
	userModel.Age = userDto.Age
	userModel.Email = userDto.Email
	userModel.CurrentAddress = userDto.CurrentAddress
	userModel.PermanentAddress = userDto.PermanentAddress
}
func (userHandler *UserHandler) Create(userDto *v1.UserDto) (*v1.CreateUserReply, error) {
	userModel := convertFrom(userDto)
	user, err := userHandler.userRepository.Create(&userModel)
	if err != nil {
		return nil, err
	}
	return &v1.CreateUserReply{
		UserResponse: &v1.UserResponse{
			UserDto: convertTo(*user),
		},
	}, nil
}

func (userHandler *UserHandler) Update(userDto *v1.UserDto) (*v1.UpdateUserReply, error) {
	userModel, err := userHandler.GetById(userDto.GetId())
	if err != nil {
		return nil, err
	}
	update(userDto, userModel)
	user, err := userHandler.userRepository.Update(userModel)
	if err != nil {
		return nil, err
	}
	return &v1.UpdateUserReply{
		UserResponse: &v1.UserResponse{
			UserDto: convertTo(*user),
		},
	}, nil
}

func (userHandler *UserHandler) GetAll() ([]*v1.UserDto, error) {
	users, err := userHandler.userRepository.FindAll()
	if err != nil {
		return nil, err
	}
	var userDtos []*v1.UserDto
	for _, user := range users {
		userDtos = append(userDtos, convertTo(*user))
	}
	return userDtos, nil
}

func (userHandler *UserHandler) GetById(id int32) (*models.UserModel, error) {
	user, err := userHandler.userRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (userHandler *UserHandler) Delete(id int32) error {
	if err := userHandler.userRepository.DeleteById(id); err != nil {
		return err
	}
	return nil
}
