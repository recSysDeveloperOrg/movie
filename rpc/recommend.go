package rpc

import (
	"context"
	"errors"
	"log"
	"movie/constant"
	"movie/idl/gen/recommend"
)

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
	if resp.BaseResp == nil || resp.BaseResp.Code != constant.RetSuccess.Code {
		return nil, errors.New("response code not zero")
	}

	log.Printf("rpc response from recommend:%+v", resp)
	return resp, nil
}
