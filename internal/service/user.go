package service

import (
	"context"
	"demo/internal/biz"
	"demo/internal/data/models"
	"fmt"

	pb "demo/api/user/v1"
)

type UserService struct {
	pb.UnimplementedUserServer
	uh biz.UserHandler
}

func NewUserService(userHandler *biz.UserHandler) *UserService {
	return &UserService{uh: *userHandler}
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	fmt.Println("CreateUser", req)
	response, err := s.uh.Create(req.GetUser())
	if err != nil {
		return nil, err
	}
	return response, nil
}
func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserReply, error) {
	response, err := s.uh.Update(req.GetUserDto())
	if err != nil {
		return nil, err
	}
	return response, nil
}
func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserReply, error) {
	err := s.uh.Delete(req.GetId())
	success := err == nil
	var errMsg string
	if !success {
		errMsg = err.Error()
	}
	return &pb.DeleteUserReply{Success: success, Error: errMsg}, nil
}
func convertTo(userModel models.UserModel) *pb.UserDto {
	return &pb.UserDto{
		Id:               int32(userModel.ID),
		Name:             userModel.Name,
		Age:              userModel.Age,
		Email:            userModel.Email,
		CurrentAddress:   userModel.CurrentAddress,
		PermanentAddress: userModel.PermanentAddress,
	}
}
func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	response, err := s.uh.GetById(req.GetId())
	if err != nil {
		return nil, err
	}
	return &pb.GetUserReply{UserDto: convertTo(*response)}, nil
}
func (s *UserService) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserReply, error) {
	response, err := s.uh.GetAll()
	if err != nil {
		return nil, err
	}
	return &pb.ListUserReply{UserDtos: response}, nil
}
