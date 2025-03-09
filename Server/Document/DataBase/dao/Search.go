package dao

import (
	"context"
	"github.com/Rinai-R/Gocument/Server/Document/DataBase/DB"
	"github.com/Rinai-R/Gocument/Server/Document/DataBase/conf/DB"
	"github.com/Rinai-R/Gocument/pkg/Logger"
	"github.com/Rinai-R/Gocument/pkg/models"
	"github.com/olivere/elastic/v7"
	"reflect"
)

func Search(ctx context.Context, Doc models.SearchDocument) []models.ElasticDocument {
	query := elastic.NewBoolQuery().
		Must(
			elastic.NewMatchQuery("all", Doc.Content),
		).
		Filter(
			elastic.NewMatchQuery("user_id", Doc.UserId),
		)
	search, err := DB.ES.Search().
		Index(conf.DocDB.ElasticSearch.IndexName).
		Query(query).
		Do(ctx)
	if err != nil {
		Logger.Logger.Debug("Dao:Search Error " + err.Error())
		return nil
	}
	var documents []models.ElasticDocument
	for _, item := range search.Each(reflect.TypeOf(models.ElasticDocument{})) {
		if doc, ok := item.(models.ElasticDocument); ok {
			documents = append(documents, doc)
		}
	}
	return documents
}
