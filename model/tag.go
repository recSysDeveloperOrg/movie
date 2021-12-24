package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sync"
)

type TagDao struct {
}

type Tag struct {
	Id      string `bson:"_id"`
	Content string `bson:"content"`
}

var tagDao *TagDao
var tagDaoOnce sync.Once

func NewTagDao() *TagDao {
	tagDaoOnce.Do(func() {
		tagDao = &TagDao{}
	})

	return tagDao
}

func (*TagDao) FindTagContentByTagID(ctx context.Context, ids []string) (map[string]string, error) {
	tagObjectIDs := make([]primitive.ObjectID, len(ids))
	for i, id := range ids {
		tagObjectID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, err
		}
		tagObjectIDs[i] = tagObjectID
	}
	c, err := GetClient().Collection(CollectionTag).
		Find(ctx, bson.D{{"_id", bson.D{{"$in", tagObjectIDs}}}})
	if err != nil {
		return nil, err
	}

	var tags []*Tag
	if err := c.All(ctx, &tags); err != nil {
		return nil, err
	}
	res := make(map[string]string)
	for _, tag := range tags {
		res[tag.Id] = tag.Content
	}

	return res, nil
}
