package main

import (
	"context"
	"log"
	"net"
	"os"
	"time"

	firebase "firebase.google.com/go"
	"github.com/golang/protobuf/ptypes"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	user "github.com/shota-aa/grpc-pr/pb"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	ctx := context.Background()
	projectID := "kauche-practice"
	conf := &firebase.Config{ProjectID: projectID}
	sa := option.WithCredentialsFile("credential/serviceAccount.json")
	app, err := firebase.NewApp(ctx, conf, sa)
	// app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Fatalln(err)
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()
	log.Println("aaaaaaaaaaaa")

	// client, err := firestore.NewClient(ctx, "kauche-practice")
	// if err != nil {
	// 	log.Fatal("aaaa")
	// 	log.Fatal(err)
	// }

	log.Print("aa")
	doc, _, err := client.Collection("users").Add(ctx, map[string]interface{}{
		"name": "hello1",
	})
	log.Print("ww")
	if err != nil {
		log.Fatal(err)
	}
	println(doc.ID)

	// client, err := firestore.NewClient(ctx, "kauche-practice")
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }
	// defer client.Close()
	// iter := client.Collection("spots").Documents(ctx)
	// defer iter.Stop()

	// for {
	// 	doc, err := iter.Next()
	// 	if err == iterator.Done {
	// 		break
	// 	}

	// 	if err != nil {
	// 		log.Fatalf("Failed to iterate: %v<br>", err)
	// 	}

	// 	fmt.Printf("%v\n", doc.Data())
	// }

	// logging
	zap, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("failed to set: %v", err)
	}
	zap_opt := grpc_zap.WithLevels( // --- ②
		func(c codes.Code) zapcore.Level {
			var l zapcore.Level
			switch c {
			case codes.OK:
				l = zapcore.InfoLevel

			case codes.Internal:
				l = zapcore.ErrorLevel

			default:
				l = zapcore.DebugLevel
			}
			return l
		},
	)

	s := grpc.NewServer(grpc_middleware.WithUnaryServerChain(
		grpc_ctxtags.UnaryServerInterceptor(),
		grpc_zap.UnaryServerInterceptor(zap, zap_opt),
	))
	user.RegisterUserServiceServer(s, &server{})
	reflection.Register(s)
	log.Printf("Listening on %v", ":"+port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type server struct{}

func (s *server) GetUser(ctx context.Context, req *user.GetUserRequest) (*user.User, error) {
	log.Printf("Received: %v", req.Id)
	now, _ := ptypes.TimestampProto(time.Now())
	return &user.User{Id: req.Id, Name: "John Smith", Email: "johnsmith@example.com", UpdatedAt: now}, nil
}
