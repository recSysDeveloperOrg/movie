package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"sync"
)

type MovieDao struct {
}

var movieDao *MovieDao
var movieDaoOnce sync.Once

func NewMovieDao() *MovieDao {
	movieDaoOnce.Do(func() {
		movieDao = &MovieDao{}
	})

	return movieDao
}

type Participant struct {
	Character string `bson:"character"`
	Name      string `bson:"name"`
}

type Movie struct {
	Id           string         `bson:"_id"`
	Title        string         `bson:"title"`
	PicUrl       string         `bson:"pic_url"`
	Introduction string         `bson:"introduction"`
	Participants []*Participant `bson:"participants"`
	ReleaseDate  string         `bson:"release_date"`
	Language     string         `bson:"language"`
}

var movieID2MovieCache = make(map[string]*Movie)
var movieID2MovieCacheLock = sync.RWMutex{}

func (*MovieDao) FindMovies(ctx context.Context, movieIds []string) (map[string]*Movie, error) {
	cacheMissIDs := make([]string, 0)
	res := make(map[string]*Movie)
	movieID2MovieCacheLock.RLock()
	defer movieID2MovieCacheLock.RUnlock()
	for _, movieID := range movieIds {
		if movie, ok := movieID2MovieCache[movieID]; ok {
			res[movieID] = movie
			continue
		}

		cacheMissIDs = append(cacheMissIDs, movieID)
	}

	var movies []*Movie
	c, err := GetClient().Collection(CollectionMovie).Find(ctx,
		bson.D{{"_id", bson.D{{"$in", cacheMissIDs}}}})
	if err != nil {
		return nil, err
	}
	if err := c.All(ctx, &movies); err != nil {
		return nil, err
	}
	for _, movie := range movies {
		res[movie.Id] = movie
	}

	go func(movies []*Movie) {
		movieID2MovieCacheLock.Lock()
		defer movieID2MovieCacheLock.Unlock()
		for _, movie := range movies {
			movieID2MovieCache[movie.Id] = movie
		}
	}(movies)

	return res, nil
}
