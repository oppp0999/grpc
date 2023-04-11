package main

import ( 
        "fmt"
        "log"
        "net"
        "google.golang.org/grpc"
        "./chat"
)
 
func main(){ 

        fmt.Println("go gRPC Tutorial")

        lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
        if err != nil { 
                log.Fatalf("failed to listen: %v", err)}
                
        s := chat.Server{}
        grpcServer := grpc.NewServer()

        chat.RegisterChatServiceServer(grpcServer, &s)

        if err := grpcServer.Serve(lis); err != nil{
                log.Fatalf("faild to serve: %s", err)
        }
}





/* package main

import "fmt"

func main() {
	fmt.Println("hellow world")
} */