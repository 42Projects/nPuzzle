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
	"strconv"
	"strings"
	"time"
)

var port = flag.Int("port", 9090, "the server port")

type server struct{}

func (s *server) Greets(ctx context.Context, handshake *pb.Handshake) (*pb.Handshake, error) {

	log.Printf("client handshake: %v\n", handshake.Message)
	return &pb.Handshake{Message: "pong!"}, nil
}

func computeDuration(begin, end time.Time) string {

	nanoseconds := end.UnixNano() - begin.UnixNano()
	duration := strconv.FormatInt(nanoseconds, 10)
	size := len(duration)
	if size < 10 {
		duration = fmt.Sprintf("%v%v", strings.Repeat("0", 10-size), duration)
		size = 10
	}

	return fmt.Sprintf("%v.%v", duration[:size-9], duration[size-9:size-6])
}

func (s *server) Solve(ctx context.Context, problem *pb.Problem) (*pb.Result, error) {

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
	default:
		return &pb.Result{
			Success: false,
			Error: fmt.Sprintf("error: %#v isn't recognized as a valid heuristic function\nAvailable heuristics "+
				"are:\n - hamming\n - manhattan\n - manhattan + linear conflicts (default)\n", problem.Heuristic),
		}, nil
	}

	/* Choose between greedy search and uniform-cost search */
	var search npuzzle.Search
	switch problem.Search {
	case "greedy":
		search = npuzzle.GreedySearch
	case "uniform-cost":
		search = npuzzle.UniformCostSearch
	default:
		return &pb.Result{
			Success: false,
			Error: fmt.Sprintf("error: %#v isn't recognized as a valid search option\nAvailable search "+
				"options are:\n - greedy\n - uniform-cost (default)", problem.Search),
		}, nil
	}

	m, err := npuzzle.ParseMatrix(problem.Matrix)
	if err != nil {
		log.Println(err)
		return &pb.Result{
			Success: false,
			Error:   err.Error(),
		}, nil
	}

	if npuzzle.IsSolvable(m) == false {
		log.Println("failed to solve problem: unsolvable")
		return &pb.Result{
			Success: false,
			Error:   "unsolvable",
		}, nil
	}

	begin := time.Now()
	log.Printf("starting solve on %v", m)
	res, totalNumberOfStates, maxNumberOfStates := m.Solve(problem.Search, heuristic, search)
	duration := computeDuration(begin, time.Now())
	log.Printf("solved problem %#v in %v seconds", problem.Matrix, duration)

	return &pb.Result{
		Success:     true,
		Time:        duration,
		TotalStates: int32(totalNumberOfStates),
		MaxStates:   int32(maxNumberOfStates),
		Moves:       int32(res.Cost),
	}, nil
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
