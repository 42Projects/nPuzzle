package main

import (
	"context"
	"flag"
	"fmt"
	pb "github.com/terry-finkel/nPuzzle/proto"
	npuzzle "github.com/terry-finkel/nPuzzle/src"
	"google.golang.org/grpc"
	"log"
	"net"
)

var port = flag.Int("port", 9090, "the server port")

type server struct{}

func (s *server) Greets(ctx context.Context, handshake *pb.Handshake) (*pb.Handshake, error) {

	log.Printf("client handshake: %v\n", handshake.Message)
	return &pb.Handshake{Message: "pong!"}, nil
}

func (s *server) Solve(ctx context.Context, problem *pb.Problem) (*pb.Empty, error) {

	log.Printf("received problem: %v\n", problem)

	/* Choose heuristic function */
	var heuristic npuzzle.Heuristic
	switch problem.Heuristic {
	case "hamming":
		heuristic = npuzzle.HammingDistance
	case "manhattan":
		heuristic = npuzzle.ManhattanDistance
	case "manhattan + linear conflicts":
		heuristic = npuzzle.ManhattanPlusLinearConflicts
	}

	/* Choose between greedy search and uniform-cost search */
	var search npuzzle.Search
	switch problem.Search {
	case "greedy":
		search = npuzzle.GreedySearch
	case "uniform-cost":
		search = npuzzle.UniformCostSearch
	}

	log.Println(problem.Matrix)

	m, err := npuzzle.ParseMatrix(problem.Matrix)
	if err != nil {
		log.Println(err)
	}

	if npuzzle.IsSolvable(m) == false {
		log.Println("unsolvable")
	}

	res, totalNumberOfStates, maxNumberOfStates := m.Solve(problem.Search, heuristic, search)
	fmt.Printf("\n -------------- SOLVED! ---------------\n")
	fmt.Printf("- heuristic used: %v\n", problem.Heuristic)
	fmt.Printf("- search: %v\n", problem.Search)
	fmt.Printf("- total number of states selected: %v\n", totalNumberOfStates)
	fmt.Printf("- maximum number of states in memory: %v\n", maxNumberOfStates)
	fmt.Printf("- number of moves: %v\n", res.Cost)

	return &pb.Empty{}, nil
}

func main() {

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	s := grpc.NewServer()
	pb.RegisterNpuzzleServer(s, &server{})
	log.Printf("starting server on port %v\n", *port)
	log.Fatalf("failed to serve: %v\n", s.Serve(lis))
}
