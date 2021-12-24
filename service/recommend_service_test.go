package service

import (
	"context"
	"movie/constant"
	"movie/idl/gen/movie"
	"movie/idl/gen/recommend"
	"testing"
)

func TestRecommendService_RecommendMovies(t *testing.T) {
	svc := NewRecommendService()
	ctx := NewRecommendContext(context.Background(),
		&movie.RecommendReq{
			UserId: "",
			Page:   0,
			Offset: 66,
		})
	svc.RecommendMovies(ctx)
	if ctx.recommendMovies[0].RsType != recommend.RecommendSourceType_RECOMMEND_SOURCE_TYPE_TOP_K {
		t.Fatal()
	}

	ctx.Req = &movie.RecommendReq{
		UserId: newUser,
		Page:   0,
		Offset: 10,
	}
	svc.RecommendMovies(ctx)
	if ctx.recommendMovies[0].RsType != recommend.RecommendSourceType_RECOMMEND_SOURCE_TYPE_TOP_K {
		t.Fatal()
	}

	ctx.Req = &movie.RecommendReq{
		UserId: ratedUser,
		Page:   0,
		Offset: 10,
	}
	svc.RecommendMovies(ctx)
	if ctx.recommendMovies[0].RsType != recommend.RecommendSourceType_RECOMMEND_SOURCE_TYPE_RATING {
		t.Fatal()
	}

	ctx.Req = &movie.RecommendReq{
		UserId: taggedUser,
		Page:   0,
		Offset: 10,
	}
	svc.RecommendMovies(ctx)
	if ctx.recommendMovies[0].RsType != recommend.RecommendSourceType_RECOMMEND_SOURCE_TYPE_TAG {
		t.Fatal()
	}

	ctx.Req = &movie.RecommendReq{
		UserId: taggedUser,
		Page:   -1,
		Offset: 10,
	}
	svc.RecommendMovies(ctx)
	if ctx.Resp.BaseResp.ErrNo != constant.RetParamsErr.Code {
		t.Fatal()
	}

	ctx.Req = &movie.RecommendReq{
		UserId: taggedUser,
		Page:   0,
		Offset: -1,
	}
	svc.RecommendMovies(ctx)
	if ctx.Resp.BaseResp.ErrNo != constant.RetParamsErr.Code {
		t.Fatal()
	}
}
