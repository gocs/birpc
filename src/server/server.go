package main

import (
	"io"
	"log"
	"net"

	pb "github.com/gocs/birpc/src/proto"
	"google.golang.org/grpc"
)

var posX int64
var posY int64

type server struct{}

func (s server) MousePos(srv pb.MouseService_MousePosServer) error {

	log.Println("start new server")
	ctx := srv.Context()

	resp := pb.Pos{PosX: posX, PosY: posX}
	if err := srv.Send(&resp); err != nil {
		log.Printf("send error %v", err)
	}

	for {

		// exit if context is done
		// or continue
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		// receive data from stream
		req, err := srv.Recv()
		if err == io.EOF {
			// return will close stream from server side
			log.Println("exit")
			return nil
		}
		if err != nil {
			log.Printf("receive error %v", err)
			continue
		}

		// continue if message is empty
		// if req.GetPosX() == "" {
		// 	continue
		// }

		// update max and send it to stream
		resp := pb.Pos{PosX: posX, PosY: posX}
		if err := srv.Send(&resp); err != nil {
			log.Printf("send error %v", err)
		}
		log.Println(">", req.GetPosX(), req.GetPosY())
	}
}

func main() {
	// create listiner
	lis, err := net.Listen("tcp", ":50005")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// create grpc server
	s := grpc.NewServer()
	pb.RegisterMouseServiceServer(s, server{})

	// and start...
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
