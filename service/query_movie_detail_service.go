package service

import (
	"context"
	"log"
	. "movie/constant"
	"movie/idl/gen/movie"
	"movie/model"
	"movie/rpc"
	"strings"
	"sync"
)

type QueryMovieDetailService struct {
}

type QueryMovieDetailContext struct {
	Ctx     context.Context
	Req     *movie.MovieDetailReq
	Resp    *movie.MovieDetailResp
	ErrCode *ErrorCode

	movie *model.Movie
}

var queryMovieDetailService *QueryMovieDetailService
var queryMovieDetailServiceOnce sync.Once

func NewQueryMovieDetailService() *QueryMovieDetailService {
	queryMovieDetailServiceOnce.Do(func() {
		queryMovieDetailService = &QueryMovieDetailService{}
	})

	return queryMovieDetailService
}

func NewQueryMovieDetailContext(ctx context.Context, req *movie.MovieDetailReq) *QueryMovieDetailContext {
	return &QueryMovieDetailContext{
		Ctx: ctx,
		Req: req,
		Resp: &movie.MovieDetailResp{
			BaseResp: &movie.BaseResp{},
		},
	}
}

func (s *QueryMovieDetailService) QueryMovieDetail(ctx *QueryMovieDetailContext) {
	defer s.buildResponse(ctx)
	if s.checkParams(ctx); ctx.ErrCode != nil {
		return
	}
	if s.queryDetail(ctx); ctx.ErrCode != nil {
		return
	}
	s.addViewLog(ctx)
}

func (*QueryMovieDetailService) checkParams(ctx *QueryMovieDetailContext) {
	if strings.TrimSpace(ctx.Req.Id) == "" {
		ctx.ErrCode = BuildErrCode("empty movie id", RetParamsErr)
		return
	}
}

func (*QueryMovieDetailService) queryDetail(ctx *QueryMovieDetailContext) {
	modelMovie, err := model.NewMovieDao().FindMovies(ctx.Ctx, []string{ctx.Req.Id})
	if err != nil {
		ctx.ErrCode = BuildErrCode(err, RetReadRepoErr)
		return
	}
	if len(modelMovie) != 1 {
		ctx.ErrCode = BuildErrCode("movie ID should be unique", RetReadRepoErr)
		return
	}
	ctx.movie = modelMovie[ctx.Req.Id]
}

func (*QueryMovieDetailService) addViewLog(ctx *QueryMovieDetailContext) {
	// user not login, that's fine
	if strings.TrimSpace(ctx.Req.UserId) == "" {
		return
	}
	go func() {
		if _, err := rpc.AddViewLog(ctx.Ctx, ctx.Req.UserId, ctx.Req.Id); err != nil {
			log.Printf("Add view log failed:%+v", err)
		}
	}()
}

func (*QueryMovieDetailService) buildResponse(ctx *QueryMovieDetailContext) {
	errCode := RetSuccess
	if ctx.ErrCode != nil {
		errCode = ctx.ErrCode
	}

	ctx.Resp.BaseResp.ErrNo, ctx.Resp.BaseResp.ErrMsg = errCode.Code, errCode.Msg
	ctx.Resp.Movie = &movie.Movie{
		Id:            ctx.movie.Id,
		Title:         ctx.movie.Title,
		PicUrl:        ctx.movie.PicUrl,
		Introduction:  &ctx.movie.Introduction,
		ReleaseDate:   &ctx.movie.ReleaseDate,
		Language:      &ctx.movie.Language,
		AverageRating: &ctx.movie.AverageRating,
	}
	for _, participant := range ctx.movie.Participants {
		ctx.Resp.Movie.Participant = append(ctx.Resp.Movie.Participant, &movie.Participant{
			Character: participant.Character,
			Name:      participant.Name,
		})
	}
}
