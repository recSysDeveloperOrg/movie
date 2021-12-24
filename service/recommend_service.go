package service

import (
	"context"
	"fmt"
	. "movie/constant"
	"movie/idl/gen/movie"
	"movie/idl/gen/recommend"
	"movie/model"
	"movie/rpc"
	"sync"
)

type RecommendService struct {
}

type RecommendContext struct {
	Ctx     context.Context
	Req     *movie.RecommendReq
	Resp    *movie.RecommendResp
	ErrCode *ErrorCode

	recommendMovies []*recommend.RecommendEntry
	movieID2Movie   map[string]*model.Movie
	tagID2Tag       map[string]string
}

var recommendService *RecommendService
var recommendServiceOnce sync.Once

var recommendSourceType2RecommendReasonType = map[recommend.RecommendSourceType]movie.RecommendReasonType{
	recommend.RecommendSourceType_RECOMMEND_SOURCE_TYPE_RATING: movie.RecommendReasonType_RECOMMEND_REASON_TYPE_MOVIE,
	recommend.RecommendSourceType_RECOMMEND_SOURCE_TYPE_TAG:    movie.RecommendReasonType_RECOMMEND_REASON_TYPE_TAG,
	recommend.RecommendSourceType_RECOMMEND_SOURCE_TYPE_LOG:    movie.RecommendReasonType_RECOMMEND_REASON_TYPE_LOG,
	recommend.RecommendSourceType_RECOMMEND_SOURCE_TYPE_TOP_K:  movie.RecommendReasonType_RECOMMEND_REASON_TYPE_TOP_K,
}

func NewRecommendService() *RecommendService {
	recommendServiceOnce.Do(func() {
		recommendService = &RecommendService{}
	})

	return recommendService
}

func NewRecommendContext(ctx context.Context, req *movie.RecommendReq) *RecommendContext {
	return &RecommendContext{
		Ctx: ctx,
		Req: req,
		Resp: &movie.RecommendResp{
			BaseResp: &movie.BaseResp{},
		},
	}
}

func (s *RecommendService) RecommendMovies(ctx *RecommendContext) {
	defer s.buildResponse(ctx)
	if s.checkParams(ctx); ctx.ErrCode != nil {
		return
	}
	if s.fetchRecommendMovies(ctx); ctx.ErrCode != nil {
		return
	}
	s.fillDetails(ctx)
}

func (s *RecommendService) checkParams(ctx *RecommendContext) {
	req := ctx.Req
	if req.Page < 0 {
		ctx.ErrCode = BuildErrCode("page < 0", RetParamsErr)
		return
	}
	if req.Offset < 0 {
		ctx.ErrCode = BuildErrCode("offset < 0", RetParamsErr)
		return
	}
}

func (s *RecommendService) fetchRecommendMovies(ctx *RecommendContext) {
	resp, err := rpc.RecommendMovies(ctx.Ctx, ctx.Req.UserId, ctx.Req.Page, ctx.Req.Offset)
	if err != nil {
		ctx.ErrCode = BuildErrCode(err, RetRpcCallFailed)
		return
	}

	ctx.recommendMovies = resp.Entry
}

func (s *RecommendService) fillDetails(ctx *RecommendContext) {
	if len(ctx.recommendMovies) == 0 {
		return
	}

	switch ctx.recommendMovies[0].RsType {
	case recommend.RecommendSourceType_RECOMMEND_SOURCE_TYPE_TAG:
		s.fillTagBasedDetails(ctx)
	case recommend.RecommendSourceType_RECOMMEND_SOURCE_TYPE_RATING,
		recommend.RecommendSourceType_RECOMMEND_SOURCE_TYPE_LOG,
		recommend.RecommendSourceType_RECOMMEND_SOURCE_TYPE_TOP_K:
		s.fillMovieBasedDetails(ctx)
	default:
		ctx.ErrCode = BuildErrCode(fmt.Sprintf("invalid source type:%v", ctx.recommendMovies[0].RsType),
			RetParamsErr)
	}
}

func (s *RecommendService) fillMovieBasedDetails(ctx *RecommendContext) {
	recMovies := ctx.recommendMovies
	movieIDs := make([]string, 0, 2*len(recMovies))
	for _, recMovie := range recMovies {
		movieIDs = append(movieIDs, recMovie.MovieId)
		// 特殊处理TopK推荐SourceID为空
		if recMovie.SourceId == "" {
			continue
		}
		movieIDs = append(movieIDs, recMovie.SourceId)
	}
	movieID2Movie, err := model.NewMovieDao().FindMovies(ctx.Ctx, movieIDs)
	if err != nil {
		ctx.ErrCode = BuildErrCode(err, RetReadRepoErr)
		return
	}
	ctx.movieID2Movie = movieID2Movie
}

func (s *RecommendService) fillTagBasedDetails(ctx *RecommendContext) {
	recMovies := make([]string, len(ctx.recommendMovies))
	for i, recMovieID := range ctx.recommendMovies {
		recMovies[i] = recMovieID.MovieId
	}
	movieID2Movie, err := model.NewMovieDao().FindMovies(ctx.Ctx, recMovies)
	if err != nil {
		ctx.ErrCode = BuildErrCode(err, RetReadRepoErr)
		return
	}
	ctx.movieID2Movie = movieID2Movie

	tagIDs := make([]string, 0, len(ctx.recommendMovies))
	for _, recMovieID := range ctx.recommendMovies {
		tagIDs = append(tagIDs, recMovieID.SourceId)
	}
	tagID2TagContent, err := model.NewTagDao().FindTagContentByTagID(ctx.Ctx, tagIDs)
	if err != nil {
		ctx.ErrCode = BuildErrCode(ctx, RetReadRepoErr)
		return
	}
	ctx.tagID2Tag = tagID2TagContent
}

func (s *RecommendService) buildResponse(ctx *RecommendContext) {
	errCode := RetSuccess
	if ctx.ErrCode != nil {
		errCode = ctx.ErrCode
	}
	ctx.Resp.BaseResp.ErrNo = errCode.Code
	ctx.Resp.BaseResp.ErrMsg = errCode.Msg

	if len(ctx.recommendMovies) == 0 {
		return
	}
	rsType := ctx.recommendMovies[0].RsType
	if rsType == recommend.RecommendSourceType_RECOMMEND_SOURCE_TYPE_TAG {
		s.buildTagBasedResponse(ctx)
		return
	}
	s.buildMovieBasedResponse(ctx, rsType)
}

func (s *RecommendService) buildTagBasedResponse(ctx *RecommendContext) {
	ctx.Resp.Movies = s.buildMovies(ctx, func(sourceID string) *movie.RecommendReason {
		return &movie.RecommendReason{
			ReasonType: movie.RecommendReasonType_RECOMMEND_REASON_TYPE_TAG,
			TagReason:  ctx.tagID2Tag[sourceID],
		}
	})
}

func (s *RecommendService) buildMovieBasedResponse(ctx *RecommendContext, rsType recommend.RecommendSourceType) {
	ctx.Resp.Movies = s.buildMovies(ctx, func(sourceID string) *movie.RecommendReason {
		sourceModelMovie := ctx.movieID2Movie[sourceID]
		// 对于topK推荐来说，sourceID为空，需要特殊处理
		if sourceModelMovie == nil {
			sourceModelMovie = &model.Movie{}
		}
		return &movie.RecommendReason{
			ReasonType: recommendSourceType2RecommendReasonType[rsType],
			MovieReason: &movie.Movie{
				Id:     sourceModelMovie.Id,
				Title:  sourceModelMovie.Title,
				PicUrl: sourceModelMovie.PicUrl,
			},
		}
	})
}

func (*RecommendService) buildMovies(ctx *RecommendContext,
	reasonFillFunc func(sourceID string) *movie.RecommendReason) []*movie.Movie {
	movies := make([]*movie.Movie, 0)
	for _, recPair := range ctx.recommendMovies {
		modelMovie := ctx.movieID2Movie[recPair.MovieId]
		respMovie := &movie.Movie{
			Id:            modelMovie.Id,
			Title:         modelMovie.Title,
			PicUrl:        modelMovie.PicUrl,
			Introduction:  &modelMovie.Introduction,
			ReleaseDate:   &modelMovie.ReleaseDate,
			Language:      &modelMovie.Language,
			Reason:        reasonFillFunc(recPair.SourceId),
			AverageRating: &modelMovie.AverageRating,
		}
		for _, participant := range modelMovie.Participants {
			respMovie.Participant = append(respMovie.Participant, &movie.Participant{
				Character: participant.Character,
				Name:      participant.Name,
			})
		}
		movies = append(movies, respMovie)
	}

	return movies
}
