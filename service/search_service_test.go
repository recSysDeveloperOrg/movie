package service

import (
	"context"
	"movie/constant"
	"movie/idl/gen/movie"
	"testing"
)

func TestSearchService_SearchMovies(t *testing.T) {
	sCtx := NewSearchServiceContext(context.Background(), &movie.SearchReq{
		Page:    0,
		Offset:  10,
		Keyword: "玩具",
	})
	NewSearchService().SearchMovies(sCtx)
	if sCtx.Resp.BaseResp.ErrNo != constant.RetSuccess.Code {
		t.Fatal()
	}

	sCtx.Req.Keyword = ""
	NewSearchService().SearchMovies(sCtx)
	if sCtx.Resp.BaseResp.ErrNo != constant.RetParamsErr.Code {
		t.Fatal()
	}

	sCtx.Req.Keyword = "玩具"
	sCtx.Req.Page = -1
	NewSearchService().SearchMovies(sCtx)
	if sCtx.Resp.BaseResp.ErrNo != constant.RetParamsErr.Code {
		t.Fatal()
	}

	sCtx.Req.Page = 0
	sCtx.Req.Offset = -1
	NewSearchService().SearchMovies(sCtx)
	if sCtx.Resp.BaseResp.ErrNo != constant.RetParamsErr.Code {
		t.Fatal()
	}
}
