package handler

import (
	"context"
	"errors"
	"net/http"

	"github.com/shota-aa/grpc-pr/util/random"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func setCookieToCtx(ctx context.Context) error {
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

func getSessionIDFromCtx(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", errors.New("cannot get metadata")
	}
	vs, ok := md["cookie"]
	if !ok {
		return "", errors.New("no cookie")
	}
	rawCookie := vs[0]
	if len(rawCookie) != 0 {
		return "", errors.New("no cookie content")
	}

	parser := &http.Request{Header: http.Header{"cookie": []string{rawCookie}}}
	cookie, err := parser.Cookie("session")
	if err != nil {
		return "", nil
	}
	return cookie.Value, nil
} 