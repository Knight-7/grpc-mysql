package auth

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type Authorizer struct {
	Login    string
	Password string
	OpenTLS  bool
}

func (a *Authorizer) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"login": "login",
		"pass":  "pass",
	}, nil
}

func (a *Authorizer) RequireTransportSecurity() bool {
	return a.OpenTLS
}

func (a *Authorizer) Auth(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return fmt.Errorf("token not found")
	}

	var login, pass string

	if val, ok := md["login"]; ok {
		login = val[0]
	}
	if val, ok := md["pass"]; ok {
		pass = val[0]
	}

	if login != a.Login || pass != a.Password {
		return status.Errorf(codes.Unauthenticated, "token invalid")
	}

	return nil
}
