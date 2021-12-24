package service

import (
	"context"
	"movie/idl/gen/movie"
	"testing"
)

func TestQueryMovieDetailService_QueryMovieDetail(t *testing.T) {
	ctx := NewQueryMovieDetailContext(context.Background(), &movie.MovieDetailReq{
		Id:     "100000000000000000000004",
		UserId: ratedUser,
	})
	NewQueryMovieDetailService().QueryMovieDetail(ctx)
	t.Logf("%+v", ctx.Resp)
}
