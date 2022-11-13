package server

import (
	"context"

	"github.com/adrisongomez/grpc_golang/models"
	"github.com/adrisongomez/grpc_golang/repository"
	"github.com/adrisongomez/grpc_golang/studentpb"
)

type StudentServer struct {
	Server
	studentpb.UnimplementedStudentServiceServer
}

func NewStudentServer(repo repository.Repository) *StudentServer {
	return &StudentServer{
		Server: Server{repo: repo},
	}
}

func (s *StudentServer) GetStudent(ctx context.Context, req *studentpb.GetStudentRequest) (*studentpb.Student, error) {
	student, err := s.repo.GetStudent(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	student_response := studentpb.Student{
		Id:   student.Id,
		Name: student.Name,
		Age:  student.Age,
	}

	return &student_response, nil
}

func (s *StudentServer) SetStudent(ctx context.Context, req *studentpb.Student) (*studentpb.SetStudentResponse, error) {
	student := &models.Student{
		Id:   req.GetId(),
		Name: req.GetName(),
		Age:  req.GetAge(),
	}

	err := s.repo.SetStudent(ctx, student)
	if err != nil {
		return nil, err
	}
	response := &studentpb.SetStudentResponse{
		Id: student.Id,
	}
	return response, nil
}

