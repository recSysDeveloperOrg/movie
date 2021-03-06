package rpc

import (
	"context"
	"log"
	"movie/idl/gen/movie"
	"movie/idl/gen/recommend"
)

var feedbackType2FilterType = map[movie.FeedbackType]recommend.FilterType{
	movie.FeedbackType_FEEDBACK_TYPE_MOVIE:    recommend.FilterType_FILTER_TYPE_MOVIE,
	movie.FeedbackType_FEEDBACK_TYPE_CATEGORY: recommend.FilterType_FILTER_TYPE_TAG,
}

func RecommendMovies(ctx context.Context, userID string, page, offset int64) (*recommend.RecommendResp, error) {
	req := &recommend.RecommendReq{
		UserId: userID,
		Page:   page,
		Offset: offset,
	}
	log.Printf("rpc from movie to recommend, req:%+v", req)
	resp, err := RecommendClient.Recommend(ctx, req)
	if err != nil {
		log.Printf("rpc failed:%v", err)
		return nil, err
	}

	log.Printf("rpc response from recommend:%+v", resp)
	return resp, nil
}

func SendFeedback(ctx context.Context, userID string, uninterestedID string,
	uninterestedType movie.FeedbackType) (*recommend.FilterRuleResp, error) {
	req := &recommend.FilterRuleReq{
		UserId:   userID,
		SourceId: uninterestedID,
		FType:    feedbackType2FilterType[uninterestedType],
	}
	log.Printf("rpc from movie to recommend, req:%+v", req)
	resp, err := RecommendClient.AddFilterRule(ctx, req)
	if err != nil {
		return nil, err
	}

	log.Printf("rpc response from recommend:%+v", resp)
	return resp, nil
}

func AddViewLog(ctx context.Context, userID string, movieID string) (*recommend.ViewLogResp, error) {
	req := &recommend.ViewLogReq{
		UserId:  userID,
		MovieId: movieID,
	}
	log.Printf("rpc from movie to recommend, req:%+v", req)
	resp, err := RecommendClient.AddViewLog(ctx, req)
	if err != nil {
		return nil, err
	}

	log.Printf("rpc response from recommend:%+v", resp)
	return resp, nil
}
