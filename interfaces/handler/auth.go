package handler

import (
	"context"
	"net/http"

	"github.com/shota-aa/grpc-pr/util/random"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func setCookie(ctx context.Context) error {
	value, err := random.MakeRandomStr(30)
	if err != nil {
		return err
	}
	cookie := http.Cookie{
		// handlerがもつようにしたほうがいいかも
		Name: "session",
		Value: value,
		Path: "/",
		// Domain: "localhost",
		MaxAge: 604800,
		Secure: true,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}

	md := make(metadata.MD, 1)
	md.Set("set-cookie", cookie.String())
	grpc.SetHeader(ctx, md)
	return nil
}

func getSessionID(ctx context.Context) (string, error) {
	
} 