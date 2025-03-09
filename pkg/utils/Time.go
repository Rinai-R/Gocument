package utils

import (
	Dpb "github.com/Rinai-R/Gocument/Server/Api/Func/Document/Client/rpc"
	Upb "github.com/Rinai-R/Gocument/Server/Api/Func/User/Client/rpc"
	models2 "github.com/Rinai-R/Gocument/pkg/models"
)

func TimestampToTime(rsp *Upb.PersonalPageResponse) *models2.User {
	user := models2.User{
		Id:        rsp.User.Id,
		Username:  rsp.User.Username,
		Bio:       rsp.User.Bio,
		Gender:    rsp.User.Gender,
		Avatar:    rsp.User.Avatar,
		CreatedAt: rsp.User.CreatedAt.AsTime(),
		UpdatedAt: rsp.User.UpdatedAt.AsTime(),
	}
	var documents []models2.Document
	for _, d := range rsp.User.Documents {
		document := models2.Document{
			Id:        int(d.Id),
			Title:     d.Title,
			IsPrivate: d.IsPrivate,
			CreateAt:  d.CreateAt.AsTime(),
			UpdateAt:  d.UpdatedAt.AsTime(),
		}
		documents = append(documents, document)
	}
	user.Documents = documents
	return &user
}

func TimeTransition(rsp *Dpb.SearchResponse) []models2.Document {
	var ans []models2.Document
	for _, doc := range rsp.Documents {
		ans = append(ans, models2.Document{
			Id:        int(doc.Id),
			Title:     doc.Title,
			IsPrivate: doc.IsPrivate,
			CreateAt:  doc.CreateAt.AsTime(),
			UpdateAt:  doc.UpdatedAt.AsTime(),
		})
	}
	return ans
}
