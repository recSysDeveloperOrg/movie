package service

import (
	"context"
	"fmt"
	. "movie/constant"
	"movie/idl/gen/movie"
	"movie/model"
	"sync"
)

type UpdateMovieRatingService struct {
}

type UpdateMovieRatingContext struct {
	Ctx     context.Context
	Req     *movie.ModifyMovieRatingReq
	Resp    *movie.ModifyMovieRatingResp
	ErrCode *ErrorCode
}

var updateMovieRatingService *UpdateMovieRatingService
var updateMovieRatingServiceOnce sync.Once

func NewUpdateMovieRatingService() *UpdateMovieRatingService {
	updateMovieRatingServiceOnce.Do(func() {
		updateMovieRatingService = &UpdateMovieRatingService{}
	})

	return updateMovieRatingService
}

func NewUpdateMovieRatingContext(ctx context.Context, req *movie.ModifyMovieRatingReq) *UpdateMovieRatingContext {
	return &UpdateMovieRatingContext{
		Ctx: ctx,
		Req: req,
		Resp: &movie.ModifyMovieRatingResp{
			BaseResp: &movie.BaseResp{},
		},
	}
}

func (s *UpdateMovieRatingService) UpdateMovieRating(ctx *UpdateMovieRatingContext) {
	defer s.buildResponse(ctx)
	if s.checkParams(ctx); ctx.ErrCode != nil {
		return
	}
	s.updateRatings(ctx)

}

func (s *UpdateMovieRatingService) checkParams(ctx *UpdateMovieRatingContext) {
	req := ctx.Req
	if len(req.MovieId) == 0 {
		ctx.ErrCode = BuildErrCode("movieId 为空", RetParamsErr)
		return
	}

	if req.NUpdate > 5.0 || req.NUpdate < -5.0 {
		ctx.ErrCode = BuildErrCode("更新范围非法", RetParamsErr)
		return
	}

}

func (s *UpdateMovieRatingService) buildResponse(ctx *UpdateMovieRatingContext) {
	errCode := RetSuccess
	if ctx.ErrCode != nil {
		errCode = ctx.ErrCode
	}
	ctx.Resp.BaseResp.ErrNo = errCode.Code
	ctx.Resp.BaseResp.ErrMsg = errCode.Msg
}

func (s *UpdateMovieRatingService) updateRatings(ctx *UpdateMovieRatingContext) {
	oldMovieDetailMap, err := model.NewMovieDao().FindMovies(ctx.Ctx, []string{ctx.Req.MovieId})
	if err != nil {
		ctx.ErrCode = BuildErrCode(err, RetReadRepoErr)
		return
	}
	if len(oldMovieDetailMap) == 0 {
		ctx.ErrCode = BuildErrCode(fmt.Sprintf("找不到Movie:%s", ctx.Req.MovieId), RetReadRepoErr)
		return
	}
	oldMovieDetail := oldMovieDetailMap[ctx.Req.MovieId]
	totalRating := oldMovieDetail.AverageRating * float64(oldMovieDetail.UniqueCount)
	totalRating += ctx.Req.NUpdate
	if ctx.Req.IsNewUser {
		oldMovieDetail.UniqueCount++
	}
	oldMovieDetail.AverageRating = totalRating / float64(oldMovieDetail.UniqueCount)

	err = model.NewMovieDao().UpdateMovies(ctx.Ctx, oldMovieDetail)
	if err != nil {
		ctx.ErrCode = BuildErrCode(err, RetReadRepoErr)
		return
	}

}
