package main

import (
	"database/sql"
	handler_gph "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/PGabrielDev/clean-archi-go/configs"
	"github.com/PGabrielDev/clean-archi-go/internal/events/handler"
	"github.com/PGabrielDev/clean-archi-go/internal/infra/database"
	"github.com/PGabrielDev/clean-archi-go/internal/infra/graph"
	"github.com/PGabrielDev/clean-archi-go/internal/infra/grpc/pb"
	"github.com/PGabrielDev/clean-archi-go/internal/infra/grpc/service"
	"github.com/PGabrielDev/clean-archi-go/internal/infra/web"
	"github.com/PGabrielDev/clean-archi-go/internal/infra/web/webserver"
	usecase "github.com/PGabrielDev/clean-archi-go/internal/usecases"
	"github.com/PGabrielDev/clean-archi-go/pkg/events"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
)

func main() {

	envs, err := configs.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	stringConnection := "postgres://" + envs.DBUser + ":" + envs.DBPassword + "@" + envs.DBHost + "/" + envs.DBName + "?sslmode=disable"
	db, err := sql.Open(envs.DBDriver, stringConnection)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	defer db.Close()
	rabbitChannel, err := events.OpenChannel()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	defer rabbitChannel.Close()

	repository := database.NewOrderRepository(db)

	eventDispacher := events.NewEventDispatcher()

	rabittHandler := handler.RabbitMQHandler{RabbitMQChannel: rabbitChannel}

	eventDispacher.Register("OrderCreated", &rabittHandler)

	usecaseCreateOrder := usecase.NewCreateOrderUseCase(repository, eventDispacher)

	GraphQLPort := envs.GraphQLPort
	if GraphQLPort == "" {
		GraphQLPort = "8081"
	}

	srv := handler_gph.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{CreateOrderUseCase: usecaseCreateOrder}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", GraphQLPort)
	go http.ListenAndServe(":"+GraphQLPort, nil)
	grpcService := grpc.NewServer()
	reflection.Register(grpcService)
	serviceGRPC := service.NewOrderService(*usecaseCreateOrder)
	pb.RegisterOrderServiceServer(grpcService, serviceGRPC)
	listen, err := net.Listen("tcp", ":"+envs.GrpcPort)
	if err != nil {
		panic(err)
	}
	go grpcService.Serve(listen)
	println("TEste do cego")
	webServer := webserver.NewWebServer(envs.WebServerPort)
	orderHandler := web.NewWebOrderHandler(eventDispacher, repository)
	webServer.AddHandler("/orders", orderHandler.Create)
	webServer.Start()
}
