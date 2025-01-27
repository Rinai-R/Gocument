package utils

import (
	Dpb "github.com/Rinai-R/Gocument/App/Func/Document/Client/rpc"
	Upb "github.com/Rinai-R/Gocument/App/Func/User/Client/rpc"
	"github.com/Rinai-R/Gocument/models"
)

func TimestampToTime(rsp *Upb.PersonalPageResponse) *models.User {
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

func TimeTransition(rsp *Dpb.SearchResponse) []models.Document {
	var ans []models.Document
	for _, doc := range rsp.Documents {
		ans = append(ans, models.Document{
			Id:        int(doc.Id),
			Title:     doc.Title,
			IsPrivate: doc.IsPrivate,
			CreateAt:  doc.CreateAt.AsTime(),
			UpdateAt:  doc.UpdatedAt.AsTime(),
		})
	}
	return ans
}
