package search

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"movie/config"
)

type movieDetail struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

var esClient *elastic.Client
var cfg *config.ElasticSearch

func InitES() error {
	cfg = config.GetConfig().Es
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(cfg.Url))
	if err != nil {
		return err
	}
	esClient = client

	return nil
}

func GetMovies(ctx context.Context, page, size int, keyword string) (byTitle []string, byDetail []string, err error) {
	titleQuery := elastic.NewMatchQuery("title", keyword)
	titleQueryRes, err := esClient.Search().Query(titleQuery).Index(cfg.Index).From(page * size).Size(size).Do(ctx)
	if err != nil {
		return nil, nil, err
	}
	var mDetail movieDetail
	for _, hit := range titleQueryRes.Hits.Hits {
		b, err := hit.Source.MarshalJSON()
		if err != nil {
			return nil, nil, err
		}
		if err := json.Unmarshal(b, &mDetail); err != nil {
			return nil, nil, err
		}
		byTitle = append(byTitle, mDetail.ID)
	}

	detailQuery := elastic.NewMatchQuery("detail", keyword)
	detailQueryRes, err := esClient.Search().Query(detailQuery).Index(cfg.Index).From(page * size).Size(size).Do(ctx)
	if err != nil {
		return nil, nil, err
	}
	for _, hit := range detailQueryRes.Hits.Hits {
		b, err := hit.Source.MarshalJSON()
		if err != nil {
			return nil, nil, err
		}
		if err := json.Unmarshal(b, &mDetail); err != nil {
			return nil, nil, err
		}
		byDetail = append(byDetail, mDetail.ID)
	}

	return byTitle, byDetail, nil
}
