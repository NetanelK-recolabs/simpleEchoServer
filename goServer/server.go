package main

import (
    //"context"
    "fmt"
    "log"
    "net"
	//"net/http"

    // This import path is based on the name declaration in the go.mod,
    // and the gen/proto/go output location in the buf.gen.yaml.
	echov1 "echoapis/echo/v1/gen/proto/go/echo/v1"
	"google.golang.org/grpc"
)

/* Simple Go(lang) echo server */

func main() {
    if err := run(); err != nil {
        log.Fatal(err)
    }
}

func run() error {
    listenOn := "127.0.0.1:65432"
    listener, err := net.Listen("tcp", listenOn)
    if err != nil {
        return fmt.Errorf("failed to listen on %s: %w", listenOn, err)
    }

    server := grpc.NewServer()
    echov1.RegisterEchoServiceServer(server, &echoServiceServer{})
    log.Println("Listening on", listenOn)
    if err := server.Serve(listener); err != nil {
        return fmt.Errorf("failed to serve gRPC server: %w", err)
    }

    return nil
}

// echoServiceServer implements the EchoServiceServer API.
type echoServiceServer struct {
    echov1.UnimplementedEchoServiceServer
}

/*func main() {
    http.HandleFunc("/", HelloServer)
    http.ListenAndServe(":8080", nil)
}*/
