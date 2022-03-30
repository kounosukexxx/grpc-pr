package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"sync"
	"time"

	"cloud.google.com/go/pubsub"
	firebase "firebase.google.com/go"
	"github.com/golang/protobuf/ptypes"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	pb "github.com/shota-aa/grpc-pr/pb/proto"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
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

	// firestore
	conf := &firebase.Config{ProjectID: projectID}
	// sa := option.WithCredentialsFile("credential/serviceAccount.json")
	app, err := firebase.NewApp(ctx, conf)
	// app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Fatalln(err)
	}
	firestoreClient, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer firestoreClient.Close()
	log.Println("aaaaaaaaaaaa")

	// client, err := firestore.NewClient(ctx, "kauche-practice")
	// if err != nil {
	// 	log.Fatal("aaaa")
	// 	log.Fatal(err)
	// }

	// log.Print("aa")
	doc, _, err := firestoreClient.Collection("users").Add(ctx, map[string]interface{}{
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

	// pubsub
	pubsubClient, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer pubsubClient.Close()

	// publisher
	// topicID, _ := random.MakeRandomStr(10)
	topicID1 := "test1"
	topic1 := pubsubClient.Topic(topicID1)
	t := 0
	for t < 12 {
		t++
		result := topic1.Publish(ctx, &pubsub.Message{
			Data: []byte("Hello World" + strconv.Itoa(t)),
		})
		id, err := result.Get(ctx)
		if err != nil {
			fmt.Println(err)
		}
		// topic, err := pubsubClient.CreateTopic(topicID)
		// if err != nil {
		// 	log.Fatalf("Failed to create topic: %v", err)
		// }
		fmt.Printf("Topic %s got.\n", topic1)
		fmt.Printf("Published a message; msg ID: %v\n", id)
	}

	pbuser := &pb.GetUserResponse{
		Id:        999,
		Name:      "kounosuke",
		Email:     "dqx@com",
		UpdatedAt: timestamppb.Now(),
	}
	topicID2 := "test2-proto"
	topic2 := pubsubClient.Topic(topicID2)
	// cfg, err := topic2.Config(ctx)
	// if err != nil {
	// 	fmt.Printf("topic.Config err: %v", err)
	// }
	// cfg := &pubsub.TopicConfig{
	// 	SchemaSettings: &pubsub.SchemaSettings{
	// 		Schema:   fmt.Sprintf("projects/%s/schemas/%s", projectID, "aaa"),
	// 		Encoding: pubsub.EncodingBinary,
	// 	},
	// }
	// encoding := cfg.SchemaSettings.Encoding
	// fmt.Println(encoding)
	// var msg2 []byte
	// switch encoding {
	// case pubsub.EncodingBinary:
	// 	msg2, err = proto.Marshal(pbuser)
	// 	if err != nil {
	// 		fmt.Printf("proto.Marshal err: %v", err)
	// 	}
	// case pubsub.EncodingJSON:
	// 	msg2, err = protojson.Marshal(pbuser)
	// 	if err != nil {
	// 		fmt.Printf("protojson.Marshal err: %v", err)
	// 	}
	// default:
	// 	fmt.Printf("invalid encoding: %v", encoding)
	// }
	msg2, err := proto.Marshal(pbuser)
	if err != nil {
		fmt.Printf("proto.Marshal err: %v", err)
	}
	result2 := topic2.Publish(ctx, &pubsub.Message{
		Data: msg2,
	})
	_, err = result2.Get(ctx)
	if err != nil {
		fmt.Printf("result.Get: %v", err)
	}
	a := pb.GetUserResponse{}
	err = proto.Unmarshal(msg2, &a)
	fmt.Printf("msg2: %v\n", a)

	// subscriber
	var mu sync.Mutex
	received := 0
	subID := "test1-sub"
	sub := pubsubClient.Subscription(subID)
	cctx, cancel := context.WithCancel(ctx)
	err = sub.Receive(cctx, func(ctx context.Context, msg *pubsub.Message) {
		fmt.Println("1")
		msg.Ack()
		fmt.Printf("Got message: %q\n", string(msg.Data))
		mu.Lock()
		defer mu.Unlock()
		received++
		fmt.Println("2")
		if received == 10 {
			cancel()
		}
	})
	if err != nil {
		fmt.Print(err)
	}

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
	pb.RegisterUserServiceServer(s, &server{})
	reflection.Register(s)
	log.Printf("Listening on %v", ":"+port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type server struct{}

func (s *server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	log.Printf("Received: %v", req.Id)
	now, _ := ptypes.TimestampProto(time.Now())
	return &pb.GetUserResponse{Id: req.Id, Name: "John Smith", Email: "johnsmith@example.com", UpdatedAt: now}, nil
}

func (s *server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	return nil, nil
}
