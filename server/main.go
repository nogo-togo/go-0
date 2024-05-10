package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"
	"strconv"

	pb "github.com/nogo-togo/go-1/game"
	"google.golang.org/grpc"
)

const (
	rock = iota
	paper
	scissor
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedGameServer
}

func (s *server) Play(ctx context.Context, in *pb.GameRequest) (*pb.GameReply, error) {

	command := "Let's play, chose Rock (0), Paper (1), scissor(2)"
	log.Println(command)
	inpString := in.GetName()
	log.Printf("Received: %v", inpString)

	i, err := strconv.Atoi(inpString)
	if err != nil || i < 0 || i > 2 {
		// handle error
		fmt.Println(err)
		//os.Exit(2)
		return &pb.GameReply{Message: fmt.Sprintf("Invalid number, try again!\n%s", command)}, nil
	}
	log.Println(s, i)
	return &pb.GameReply{Message: playRockPaperScissor(i) + fmt.Sprintf("\nlet's play again!\n%s", command)}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGameServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func playRockPaperScissor(input int) string {
	answers := []int{rock, paper, scissor}

	log.Printf("read user_input: %d-\n", input)

	compInput := rand.Intn(len(answers))
	log.Println("Computer choose", compInput)

	return evaluateChoices(input, compInput)
}

func evaluateChoices(userInput int, compInput int) string {
	if userInput == compInput {
		return fmt.Sprintf("Doh, You %d vs Computer %d Play again!\n\n", userInput, compInput)
		//main()
	} else if userInput == 0 && compInput == 1 {
		return fmt.Sprintf("Computer won!")
	} else if userInput == 0 && compInput == 2 {
		return fmt.Sprintf("You  won!")
	} else if userInput == 1 && compInput == 0 {
		return fmt.Sprintf("You  won!")
	} else if userInput == 1 && compInput == 2 {
		return fmt.Sprintf("Computer won!")
	} else if userInput == 2 && compInput == 0 {
		return fmt.Sprintf("Computer won!")
	} else if userInput == 2 && compInput == 1 {
		return fmt.Sprintf("You  won!")
	} else {
		return fmt.Sprintf("I just don't care anymore!")
	}
}
