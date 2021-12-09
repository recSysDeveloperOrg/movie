package main

import (
	"context"
	movie "movie/idl/gen"
)

type Handler struct {
	*movie.UnimplementedMovieServiceServer
}

func (*Handler) RecommendMovies(ctx context.Context, req *movie.RecommendReq) (*movie.RecommendResp, error) {
	return nil, nil
}

func (*Handler) GetMovieDetail(ctx context.Context, req *movie.MovieDetailReq) (*movie.MovieDetailResp, error) {
	return nil, nil
}

func (*Handler) SearchMovies(ctx context.Context, req *movie.SearchReq) (*movie.SearchResp, error) {
	return nil, nil
}

func (*Handler) CreateMovie(ctx context.Context, req *movie.CreateReq) (*movie.CreateResp, error) {
	return nil, nil
}
