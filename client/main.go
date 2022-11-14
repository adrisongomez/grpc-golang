package main

import (
	"context"
	"io"
	"log"
	"time"

	"github.com/adrisongomez/grpc_golang/testpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	cc, err := grpc.Dial("localhost:5070", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("couldn't connect %v", err)
	}

	defer cc.Close()

	c := testpb.NewTestServiceClient(cc)

	// DoUnary(c)
	// DoClientStreaming(c)
	// DoServerStreaming(c)
	DoBidirectionalStream(c)
}

func DoUnary(c testpb.TestServiceClient) {
	req := &testpb.GetTestRequest{
		Id: "t1",
	}

	res, err := c.GetTest(context.Background(), req)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Response: %v\n", res)
}

func DoClientStreaming(c testpb.TestServiceClient) {
	questions := []*testpb.Question{
		{
			Id:       "q9t1",
			Answer:   "Maito",
			Question: "Guy",
			TestId:   "t1",
		},
		{
			Id:       "q10t1",
			Answer:   "Dai",
			Question: "Guy",
			TestId:   "t1",
		},
		{
			Id:       "q11t1",
			Answer:   "Senju",
			Question: "Tobirama",
			TestId:   "t1",
		},
		{
			Id:       "q12t1",
			Answer:   "Sarutobi",
			Question: "Hiruzen",
			TestId:   "t1",
		},
		{
			Id:       "q13t1",
			Answer:   "Sarutobi",
			Question: "Asuma",
			TestId:   "t1",
		},
	}

	stream, err := c.SetQuestions(context.Background())

	if err != nil {
		log.Fatalf("Error while calling SetQuestions: %v \n", err)
	}

	for _, question := range questions {
		log.Println("sending question: ", question.Id)
		stream.Send(question)
		time.Sleep(2 * time.Second)
	}

	msg, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("response server: %v", msg)
}

func DoServerStreaming(c testpb.TestServiceClient) {
	req := &testpb.GetStudentsPerTestRequest{
		TestId: "t1",
	}

	stream, err := c.GetStudentsPerTest(context.Background(), req)

	if err != nil {
		log.Fatalf("error while calling GetStudented per test: %v \n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("error while reading from stream. %v\n", err)
		}

		log.Printf("response: %v", msg)
	}

}

func DoBidirectionalStream(c testpb.TestServiceClient) {
	answer := testpb.TakeTestRequest{
		Answer: "Hiii!",
	}

	numberOfQuestions := 13

	waitChannel := make(chan struct{})

	stream, err := c.TakeTest(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for i := 0; i < numberOfQuestions; i++ {
			stream.Send(&answer)
		}
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatal(err)
				break
			}
			log.Println(res)
		}

		close(waitChannel)
	}()
	<-waitChannel
}
