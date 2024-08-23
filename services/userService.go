package services

import (
	"context"
	"fmt"
	pb "grpc-users/pb"

	"github.com/google/uuid"

	"gorm.io/gorm"
)

type UserService struct {
	Db *gorm.DB
	pb.UnimplementedUserServiceServer
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	newUUID := uuid.New()
	uuidString := newUUID.String()
	mapReq := map[string]interface{}{
		"id":    uuidString,
		"name":  req.Name,
		"email": req.Email,
	}
	resp := s.Db.Table("users").Create(&mapReq)
	if resp.Error != nil {
		return nil, resp.Error
	}
	return &pb.CreateUserResponse{
		Status:  true,
		Message: "User created successfully",
		Id:      uuidString,
		Name:    req.Name,
		Email:   req.Email,
	}, nil
}

func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	var user map[string]interface{}
	fmt.Println(req.Id)
	resp := s.Db.Table("users").Where("id = ?", req.Id).Take(&user)
	if resp.Error != nil {
		return nil, resp.Error
	}
	return &pb.GetUserResponse{
		Status:  true,
		Message: "User found",
		Id:      user["id"].(string),
		Name:    user["name"].(string),
		Email:   user["email"].(string),
	}, nil
}

func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	mapReq := map[string]interface{}{
		"id":    req.Id,
		"name":  req.Name,
		"email": req.Email,
	}
	resp := s.Db.Table("users").Where("id = ?", req.Id).Updates(&mapReq)
	if resp.Error != nil {
		return nil, resp.Error
	}
	return &pb.UpdateUserResponse{
		Status:  true,
		Message: "User updated successfully",
		Id:      req.Id,
		Name:    req.Name,
		Email:   req.Email,
	}, nil
}

func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	resp := s.Db.Table("users").Where("id = ?", req.Id).Delete(&req)
	if resp.Error != nil {
		return nil, resp.Error
	}
	return &pb.DeleteUserResponse{
		Status:  true,
		Message: fmt.Sprintf("User with id %s deleted successfully", req.Id),
	}, nil
}
