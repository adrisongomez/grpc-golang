package server

import (
	"context"
	"io"
	"log"

	"github.com/adrisongomez/grpc_golang/models"
	"github.com/adrisongomez/grpc_golang/repository"
	"github.com/adrisongomez/grpc_golang/studentpb"
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

func (t *TestServer) SetQuestions(stream testpb.TestService_SetQuestionsServer) error {
	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&testpb.SetQuestionResponse{Ok: true})
		}

		if err != nil {
			return err
		}

		question := models.Question{
			Id:       msg.GetId(),
			Question: msg.GetQuestion(),
			Answer:   msg.GetAnswer(),
			TestId:   msg.GetTestId(),
		}

		err = t.repo.SetQuestion(context.Background(), &question)

		if err != nil {
			return stream.SendAndClose(&testpb.SetQuestionResponse{
				Ok: false,
			})
		}

	}
}

func (t *TestServer) EnrollStudents(stream testpb.TestService_EnrollStudentsServer) error {
	for {
		msg, err := stream.Recv()
		log.Println(err)
		if err == io.EOF {
			return stream.SendAndClose(&testpb.SetQuestionResponse{Ok: true})
		}

		if err != nil {
			return err
		}

		enrollment := &models.Enrollment{
			StudentId: msg.GetStudentId(),
			TestId:    msg.GetTestId(),
		}

		err = t.repo.SetEnrollment(context.Background(), enrollment)

		log.Println(err)
		if err != nil {
			return stream.SendAndClose(&testpb.SetQuestionResponse{Ok: false})
		}
	}
}

func (t *TestServer) GetStudentsPerTest(
	req *testpb.GetStudentsPerTestRequest,
	stream testpb.TestService_GetStudentsPerTestServer,
) error {
	testId := req.GetTestId()

	students, err := t.repo.GetStudentPerTest(context.Background(), testId)
	log.Println(students)
	if err != nil {
		log.Println(err)
		return err
	}

	for _, student := range students {
		studentPb := &studentpb.Student{
			Id:   student.Id,
			Name: student.Name,
			Age:  student.Age,
		}
		err := stream.Send(studentPb)
		log.Println(err)

		if err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}

func (t *TestServer) TakeTest(stream testpb.TestService_TakeTestServer) error {

	questions, err := t.repo.GetQuestionsPerTest(context.Background(), "t1")
    log.Println(err)
	if err != nil {
		return err
	}

	i := 0
	var currentQuestion = &models.Question{}

	for {
		if i < len(questions) {
			currentQuestion = questions[i]
		}

		if i <= len(questions) {
			questionToSend := &testpb.Question{
				Id:       currentQuestion.Id,
				Question: currentQuestion.Question,
				TestId:   currentQuestion.TestId,
			}

			err := stream.Send(questionToSend)
			if err != nil {
				return err
			}
			i++
		}

		answer, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		log.Println(answer)
	}
}
