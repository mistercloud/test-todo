package main

import (
	"context"
	"flag"
	"github.com/mistercloud/test-todo/goapi/internal/entity"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"net"
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	todoRepo "github.com/mistercloud/test-todo/goapi/internal/repo"

	gw "github.com/mistercloud/test-todo/goapi/pkg/api"
)

var (
	// command-line options:
	// gRPC server endpoint
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:9090", "gRPC server endpoint")
)

type Server struct {
	gw.UnimplementedTodoServer
}

var storage = todoRepo.NewToDoMemory()

func (s *Server) AddItem(ctx context.Context, in *gw.AddItemRequest) (*gw.Item, error) {
	item, err := storage.Add(in.Title)
	return &gw.Item{Id: int32(item.Id), Title: item.Title}, err
}

func (s *Server) RemoveItem(ctx context.Context, in *gw.RemoveItemRequest) (*emptypb.Empty, error) {
	err := storage.Remove(entity.ItemId(in.GetId()))
	return &emptypb.Empty{}, err
}

func (s *Server) GetItems(ctx context.Context, _ *emptypb.Empty) (*gw.ItemList, error) {
	list := storage.List()
	ret := gw.ItemList{}
	for _, item := range list {
		retItem := gw.Item{Id: int32(item.Id), Title: item.Title}
		ret.Items = append(ret.Items, &retItem)
	}
	return &ret, nil
}

func run() error {

	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		glog.Fatal("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Greeter service to the server
	gw.RegisterTodoServer(s, &Server{})
	// Serve gRPC server
	glog.Infoln("Serving gRPC on 0.0.0.0:9090")
	go func() {
		glog.Fatal(s.Serve(lis))
	}()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err = gw.RegisterTodoHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}

	gwServer := &http.Server{
		Addr:    `:8080`,
		Handler: cors(mux),
	}

	log.Printf("Serving gRPC-Gateway on %s\n", gwServer.Addr) // запускаем HTTP сервер
	return gwServer.ListenAndServe()
}

func cors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE, PUT")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, ResponseType")
		if r.Method == "OPTIONS" {
			return
		}
		h.ServeHTTP(w, r)
	})
}

func main() {

	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
