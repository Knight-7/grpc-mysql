package utils

import (
	"rpc-mysql/models"
	pb "rpc-mysql/proto"
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
