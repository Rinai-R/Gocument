package dao

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Rinai-R/Gocument/DataBase/DB"
	conf "github.com/Rinai-R/Gocument/DataBase/conf/DB"
	"github.com/Rinai-R/Gocument/Logger"
	"github.com/Rinai-R/Gocument/Utils/Error"
	"github.com/Rinai-R/Gocument/models"
	"github.com/olivere/elastic/v7"
)

func GetDocument(ctx context.Context, ESDocument *models.ElasticDocument) error {
	res, err := DB.ES.Get().Index(conf.DB.ElasticSearch.IndexName).Id(ESDocument.Id).Do(ctx)
	if err != nil {
		if elastic.IsNotFound(err) {
			Logger.Logger.Debug("Dao: ES Document Not Found")
			return Error.DocumentNotFound
		} else {
			Logger.Logger.Error("Dao: Internal Error" + err.Error())
			return err
		}
	}
	if res.Found {
		if err = json.Unmarshal(res.Source, &ESDocument); err != nil {
			Logger.Logger.Debug("Dao: Bind Json Error " + err.Error())
			return err
		}
		fmt.Println(ESDocument)
		Logger.Logger.Debug("Dao: Bind Json Success " + ESDocument.Id)
		return nil
	}
	Logger.Logger.Debug("Dao: ES Document Not Found")
	return Error.DocumentNotFound
}
