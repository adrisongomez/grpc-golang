package server

import (
	"context"

	"github.com/adrisongomez/grpc_golang/models"
	"github.com/adrisongomez/grpc_golang/repository"
	"github.com/adrisongomez/grpc_golang/testpb"
)

type TestServer struct {
	Server
	testpb.UnimplementedTestServiceServer
}

func NewTestServer(repo repository.Repository) *TestServer {
	return &TestServer{
		Server: Server{repo: repo},
	}
}

func (t *TestServer) GetTest(ctx context.Context, req *testpb.GetTestRequest) (*testpb.Test, error) {
	test, err := t.repo.GetTest(ctx, req.GetId())

	if err != nil {
		return nil, err
	}

	response := &testpb.Test{
		Id:   test.Id,
		Name: test.Name,
	}

	return response, nil
}

func (t *TestServer) SetTest(ctx context.Context, test *testpb.Test) (*testpb.SetTestResponse, error) {
	newtest := models.Test{
		Id:   test.GetId(),
		Name: test.GetName(),
	}

	err := t.repo.SetTest(ctx, &newtest)

	if err != nil {
		return nil, err
	}

	response := &testpb.SetTestResponse{
		Id:   test.GetId(),
		Name: test.GetName(),
	}
	return response, nil
}
