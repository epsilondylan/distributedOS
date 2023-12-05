package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "https://github.com/epsilondylan/distributedOS"
)

type server struct {
	sync.RWMutex
	sync.WaitGroup
	to   time.Duration
	pool map[string]string
}

func newServer(timeout time.Duration) *server {
	return &server{
		to:   timeout,
		pool: make(map[string]string),
	}
}

func (s *server) AddRoute(ctx context.Context, req *pb.RouteRequest) (*pb.Response, error) {
	s.Lock()
	defer s.Unlock()

	if _, ok := s.pool[req.Name]; ok {
		if s.pool[req.Name] == req.Addr {
			return &pb.Response{Status: "Error: Remote socket already exists"}, nil
		}
	}

	fmt.Printf("Adding route, peername=@%s@||peeraddress=%s\n", req.Name, req.Addr)
	s.pool[req.Name] = req.Addr
	return &pb.Response{Status: "Success"}, nil
}

func (s *server) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.Response, error) {
	s.Lock()
	defer s.Unlock()

	fmt.Printf("Deleting node: %v\n", req.Name)
	delete(s.pool, req.Name)
	return &pb.Response{Status: "Success"}, nil
}

func (s *server) DispatchAll(ctx context.Context, req *pb.Message) (*pb.Repeatedresponse, error) {
	var l sync.Mutex
	peers := s.fetchPeers()

	responses := make([]*pb.Keyvalue, 0)
	for name, addr := range peers {
		s.Add(1)
		go func(name, addr string) {
			defer s.Done()
			defer func() {
				if err := recover(); err != nil {
					fmt.Printf("panic: %v", err)
				}
			}()

			resp, err := s.dispatch(addr, req.Content)
			if err != nil {
				return
			}

			l.Lock()
			responses = append(responses, &pb.Keyvalue{
				Name: name,
				Value: &pb.Response{
					Status: resp,
				},
			})
			l.Unlock()
		}(name, addr)
	}
	s.Wait()

	return &pb.Repeatedresponse{Responses: responses}, nil
}

func (s *server) Dispatch(ctx context.Context, req *pb.DispatchRequest) (*pb.Message, error) {
	peer, err := s.getPeer(req.Name)
	if err != nil {
		return nil, err
	}

	resp, err := s.dispatch(peer.Addr, req.Message.Content)
	if err != nil {
		return nil, err
	}

	return &pb.Message{Content: resp}, nil
}

// Other helper functions

func (s *server) fetchPeers() map[string]string {
	p2 := make(map[string]string)
	s.RLock()
	defer s.RUnlock()
	for k, v := range s.pool {
		p2[k] = v
	}
	return p2
}

func (s *server) getPeer(name string) (*pb.Keyvalue, error) {
	p2 := &pb.Keyvalue{}
	s.RLock()
	defer s.RUnlock()
	if _, ok := s.pool[name]; !ok {
		return p2, fmt.Errorf("unknown peer: %v", name)
	}
	p2.Name = name
	p2.Value = &pb.Response{Status: "Success"}
	return p2, nil
}

func (s *server) dispatch(addr, msg string) (string, error) {
	// Create a gRPC connection to the remote server
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return "", err
	}
	defer conn.Close()

	// Create a gRPC client
	client := pb.NewPheromonesClient(conn)

	// Call the Dispatch RPC on the remote server
	dispatchResponse, err := client.Dispatch(context.Background(), &pb.DispatchRequest{
		Name: addr, // Assuming addr is the name for simplicity, replace it with the actual name
		Message: &pb.Message{
			Content: msg,
		},
	})
	if err != nil {
		return "", err
	}

	// Return the content from the response
	return dispatchResponse.Content, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
		return
	}

	s := grpc.NewServer()
	pb.RegisterPheromonesServer(s, newServer(time.Second))
	reflection.Register(s)

	fmt.Println("Server is running on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
		return
	}
}
