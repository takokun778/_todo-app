package main

import (
	"clean-architecture/infrastructure/database"
	"clean-architecture/infrastructure/logger"
	pbTask "clean-architecture/infrastructure/proto/task"
	pbUser "clean-architecture/infrastructure/proto/user"
	tc "clean-architecture/task/controller"
	tg "clean-architecture/task/gateway"
	tu "clean-architecture/task/usecase"
	uc "clean-architecture/user/controller"
	ug "clean-architecture/user/gateway"
	uu "clean-architecture/user/usecase"
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func main() {
	// -------------------------------
	// database
	// -------------------------------
	db, _ := database.GormConnect().DB.DB()
	defer db.Close()

	// -------------------------------
	// grpc
	// -------------------------------
	server := grpc.NewServer()

	// userGateway := ug.NewUserGatewayInMememory()
	userGateway := ug.NewUserGatewayGorm()

	userUsecase := uu.NewUserUsecase(userGateway)

	userController := uc.NewUserController(userUsecase)

	pbUser.RegisterUserServiceServer(server, userController)

	// taskGateway := tg.NewTaskGatewayInMemory()
	taskGateway := tg.NewTaskGatewayGorm()

	taskUsecase := tu.NewTaskUsecase(taskGateway)

	taskController := tc.NewtaskController(taskUsecase)

	pbTask.RegisterTaskServiceServer(server, taskController)

	listenPort, err := net.Listen("tcp", ":19003")
	if err != nil {
		panic(err)
	}

	// grpc server listen
	go func() {
		logger.StartLog()
		log.Fatal(server.Serve(listenPort))
	}()

	// -------------------------------
	// grpc gateway server
	// -------------------------------
	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:19003",
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()

	err = pbUser.RegisterUserServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}
	err = pbTask.RegisterTaskServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    ":8080",
		Handler: gwmux,
	}

	logger.Debug("http://localhost:8080/")
	log.Fatal(gwServer.ListenAndServe())
}
