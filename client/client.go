package main

import(
	"log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"../server/chat"
)

func main(){
	var conn *grpc.ClientConn
	//connection to server
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err !=nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	//new client
	c := chat.NewChatServiceClient(conn)

	//func Call
	// function call
	response, err := c.SayHello(context.Background(), &chat.Message{Body: "Hello From Client!"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}

	log.Printf("Response from Richet's gRPC server: %s", response.Body)

}