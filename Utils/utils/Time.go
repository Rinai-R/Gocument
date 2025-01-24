package utils

import (
	pb "github.com/Rinai-R/Gocument/App/Func/User/Client/rpc"
	"github.com/Rinai-R/Gocument/models"
)

func TimestampToTime(rsp *pb.PersonalPageResponse) (*models.User, bool) {
	user := models.User{
		Id:        rsp.User.Id,
		Username:  rsp.User.Username,
		Bio:       rsp.User.Bio,
		Gender:    rsp.User.Gender,
		Avatar:    rsp.User.Avatar,
		CreatedAt: rsp.User.CreatedAt.AsTime(),
		UpdatedAt: rsp.User.UpdatedAt.AsTime(),
	}
	var documents []models.Document
	for _, d := range rsp.User.Documents {
		document := models.Document{
			Id:         int(d.Id),
			Title:      d.Title,
			IsPrivate:  d.IsPrivate,
			CreateTime: d.CreateAt.AsTime(),
			UpdateTime: d.UpdatedAt.AsTime(),
		}
		documents = append(documents, document)
	}
	user.Documents = documents
	return &user, true
}
