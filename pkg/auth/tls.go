package auth

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"rpc-mysql/pkg/config"
)

func GetServerCreds(cfg *config.Config) (grpc.ServerOption, error) {
	cred, err := credentials.NewServerTLSFromFile(cfg.GetCertFile(), cfg.GetKeyFile())
	if err != nil {
		return nil, err
	}

	return grpc.Creds(cred), nil
}

func GetClientCreds(cfg *config.Config) (grpc.DialOption, error) {
	creds, err := credentials.NewClientTLSFromFile(cfg.GetCertFile(), cfg.GetServerName())
	if err != nil {
		return nil, err
	}

	return grpc.WithTransportCredentials(creds), nil
}
