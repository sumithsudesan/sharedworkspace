package database

import (
    "context"
    "log"

    pb   "data-definitions/user"
)

type AuthDatabaseServiceServer struct {
    pb.UnimplementedAuthDatabaseServiceServer
    userRepository UserRepository
}

func NewAuthDatabaseServiceServer(userRepository UserRepository) *AuthDatabaseServiceServer {
    return &AuthDatabaseServiceServer{
        userRepository: userRepository,
    }
}

func (s *AuthDatabaseServiceServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
    if err := s.userRepository.CreateUser(req.Username, req.Password, req.Email); err != nil {
        return nil, err
    }
    return &pb.CreateUserResponse{Message: "User details added successfully"}, nil
}

func (s *AuthDatabaseServiceServer) GetUserByUsername(ctx context.Context, req *pb.GetUserByUsernameRequest) (*pb.GetUserByUsernameResponse, error) {
    user, err := s.userRepository.GetUserByUsername(req.Username)
    if err != nil {
        return nil, err
    }
    return &pb.GetUserByUsernameResponse{User: &pb.User{
        Id:       user.ID,
        Username: user.Username,
        Password: user.Password,
        Email:    user.Email,
    }}, nil
}

func (s *AuthDatabaseServiceServer) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
    user := &database.User{
        ID:       req.Id,
        Username: req.Username,
        Password: req.Password,
        Email:    req.Email,
    }
    if err := s.userRepository.UpdateUser(user); err != nil {
        return nil, err
    }
    return &pb.UpdateUserResponse{Message: "User details updated successfully"}, nil
}

func (s *AuthDatabaseServiceServer) DeleteUserByUsername(ctx context.Context, req *pb.DeleteUserByUsernameRequest) (*pb.DeleteUserByUsernameResponse, error) {
    if err := s.userRepository.DeleteUserByUsername(req.Username); err != nil {
        return nil, err
    }
    return &pb.DeleteUserByUsernameResponse{Message: "User details deleted successfully"}, nil
}
