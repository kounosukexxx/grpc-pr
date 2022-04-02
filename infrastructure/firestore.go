package infrastructure

import (
	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/shota-aa/grpc-pr/config"
)

func NewFirestoreClient(conf *config.Config) (*firestore.Client, error) {
	fconf := &firebase.Config{ProjectID: conf.ProjectID}
	app, err := firebase.NewApp(conf.Context, fconf)
	if err != nil {
		return nil, err
	}
	client, err := app.Firestore(conf.Context)
	if err != nil {
		return nil, err
	}
	return client, nil
}
