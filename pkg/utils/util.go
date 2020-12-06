package utils

import (
	"rpc-mysql/pkg/models"
	pb "rpc-mysql/pkg/proto"
)

func ToRpcUsers(users []*models.User) []*pb.User {
	var spec []*pb.User
	for _, user := range users {
		spec = append(spec, &pb.User{
			Id:    int32(user.Id),
			Name:  user.Name,
			Age:   int32(user.Age),
			Email: user.Email,
		})
	}
	return spec
}

func ToDaoUsers(users []*pb.User) []*models.User {
	var spec []*models.User
	for _, user := range users {
		spec = append(spec, &models.User{
			Id:    int(user.Id),
			Name:  user.Name,
			Age:   int(user.Age),
			Email: user.Email,
		})
	}
	return spec
}
