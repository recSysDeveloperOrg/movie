package main

import (
	"context"
	"log"
	"movie/idl/gen/movie"
	"movie/service"
)

type Handler struct {
	*movie.UnimplementedMovieServiceServer
}

func (*Handler) RecommendMovies(ctx context.Context, req *movie.RecommendReq) (*movie.RecommendResp, error) {
	log.Printf("%+v", req)
	rCtx := service.NewRecommendContext(ctx, req)
	service.NewRecommendService().RecommendMovies(rCtx)
	log.Printf("%+v", rCtx.Resp)

	return rCtx.Resp, nil
}

func (*Handler) GetMovieDetail(ctx context.Context, req *movie.MovieDetailReq) (*movie.MovieDetailResp, error) {
	return nil, nil
}

func (*Handler) SearchMovies(ctx context.Context, req *movie.SearchReq) (*movie.SearchResp, error) {
	log.Printf("req:%+v", req)
	sCtx := service.NewSearchServiceContext(ctx, req)
	service.NewSearchService().SearchMovies(sCtx)
	log.Printf("resp:%+v", sCtx.Resp)

	return sCtx.Resp, nil
}

func (*Handler) RecommendFeedback(ctx context.Context, req *movie.FeedbackReq) (*movie.FeedbackResp, error) {
	return nil, nil
}
