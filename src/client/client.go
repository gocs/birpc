package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"

	pb "gocs.org/birpc/src/proto"

	"google.golang.org/grpc"
)

type game struct {
	stream pb.MouseService_MousePosClient
	done   chan bool
	posX   int64
	posY   int64
}

func newGame(conn *grpc.ClientConn) *game {
	client := pb.NewMouseServiceClient(conn)
	stream, err := client.MousePos(context.Background())
	if err != nil {
		log.Fatalf("openn stream error %v", err)
	}

	return &game{
		stream: stream,
		done:   make(chan bool),
	}
}

func main() {

	// dail server
	conn, err := grpc.Dial(":50005", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can not connect with server %v", err)
	}

	// create stream
	client := pb.NewMouseServiceClient(conn)
	stream, err := client.MousePos(context.Background())
	if err != nil {
		log.Fatalf("openn stream error %v", err)
	}

	done := make(chan bool)
	game := &game{
		stream: stream,
		done:   done,
	}

	// second goroutine receives data from stream
	// and saves result in max variable
	//
	// if stream is finished it closes done channel
	go game.receive()

	// first goroutine sends random increasing numbers to stream
	// and closes int after 10 iterations
	go game.send()

	// third goroutine closes done channel
	// if context is done
	go game.close()

	if err := ebiten.Run(game.update, 320, 240, 2, "Hello, World!"); err != nil {
		log.Fatal(err)
	}

	<-done
}

func (g *game) receive() {
	for {
		resp, err := g.stream.Recv()
		if err == io.EOF {
			close(g.done)
			return
		}
		if err != nil {
			log.Fatalf("can not receive %v", err)
		}

		fmt.Println(">", resp.GetPosX(), resp.GetPosY())
	}
}

func (g *game) send() {
	defer func() {
		if err := g.stream.CloseSend(); err != nil {
			log.Println(err)
		}
	}()
	for {
		req := pb.Pos{PosX: g.posX, PosY: g.posY}
		if err := g.stream.Send(&req); err != nil {
			log.Fatalf("can not send %v", err)
		}
	}
}

func (g *game) close() {
	<-g.stream.Context().Done()
	if err := g.stream.Context().Err(); err != nil {
		log.Println(err)
	}
	close(g.done)
}

func (g *game) update(screen *ebiten.Image) error {
	posX, posY := ebiten.CursorPosition()
	g.posX = int64(posX)
	g.posY = int64(posY)

	if ebiten.IsDrawingSkipped() {
		return nil
	}
	ebitenutil.DebugPrint(screen, fmt.Sprint("pos:", posX, posY))
	return nil
}
