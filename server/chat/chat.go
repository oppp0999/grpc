package chat
import (
	"log"
	"golang.org/x/net/context"
)

// server is used to implement helloworld.GreeterServer.
type Server struct{
// Embed the unimplemented server
	//helloworld.UnimplementedGreeterServer
}

func (s *Server) SayHello (ctx context.Context, in *Message) (*Message, error){
	log.Printf("receive message body from client: %s", in.Body)
	return &Message{Body: "Hello from the gRPC server"}, nil
}