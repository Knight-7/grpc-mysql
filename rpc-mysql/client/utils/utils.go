package utils

import (
	"rpc-mysql/models"
	pb "rpc-mysql/proto"
)

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
