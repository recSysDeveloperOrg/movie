package service

import (
	"context"
	. "movie/constant"
	"movie/idl/gen/movie"
	"movie/model"
	"movie/search"
	"strings"
	"sync"
)

type SearchService struct {
}

type SearchServiceContext struct {
	Ctx     context.Context
	Req     *movie.SearchReq
	Resp    *movie.SearchResp
	ErrCode *ErrorCode

	movieID2Movie    map[string]*model.Movie
	byTitleMovieIDs  []string
	byDetailMovieIDs []string
}

var searchService *SearchService
var searchServiceOnce sync.Once

func NewSearchService() *SearchService {
	searchServiceOnce.Do(func() {
		searchService = &SearchService{}
	})

	return searchService
}

func NewSearchServiceContext(ctx context.Context, req *movie.SearchReq) *SearchServiceContext {
	return &SearchServiceContext{
		Ctx: ctx,
		Req: req,
		Resp: &movie.SearchResp{
			BaseResp: &movie.BaseResp{},
		},
	}
}

func (s *SearchService) SearchMovies(ctx *SearchServiceContext) {
	defer s.buildResponse(ctx)
	if s.checkParams(ctx); ctx.ErrCode != nil {
		return
	}
	if s.doSearchMovies(ctx); ctx.ErrCode != nil {
		return
	}
}

func (*SearchService) checkParams(ctx *SearchServiceContext) {
	if ctx.Req.Page < 0 {
		ctx.ErrCode = BuildErrCode("page < 0 is not allowed", RetParamsErr)
		return
	}
	if ctx.Req.Offset < 0 {
		ctx.ErrCode = BuildErrCode("offset < 0 is not allowed", RetParamsErr)
		return
	}
	if strings.TrimSpace(ctx.Req.Keyword) == "" {
		ctx.ErrCode = BuildErrCode("keyword should not be empty", RetParamsErr)
		return
	}
}

func (*SearchService) doSearchMovies(ctx *SearchServiceContext) {
	byTitleIDs, byDetailIDs, err := search.GetMovies(ctx.Ctx, int(ctx.Req.Page),
		int(ctx.Req.Offset), ctx.Req.Keyword)
	if err != nil {
		ctx.ErrCode = BuildErrCode(err, RetSearchESErr)
		return
	}
	movieID2Movie, err := model.NewMovieDao().FindMovies(ctx.Ctx, append(byDetailIDs, byTitleIDs...))
	if err != nil {
		ctx.ErrCode = BuildErrCode(err, RetReadRepoErr)
		return
	}
	ctx.movieID2Movie = movieID2Movie
	ctx.byTitleMovieIDs = byTitleIDs
	ctx.byDetailMovieIDs = byDetailIDs
}

func (s *SearchService) buildResponse(ctx *SearchServiceContext) {
	errCode := RetSuccess
	if ctx.ErrCode != nil {
		errCode = ctx.ErrCode
	}
	ctx.Resp.BaseResp.ErrNo, ctx.Resp.BaseResp.ErrMsg = errCode.Code, errCode.Msg

	s.movie2SearchEntry(ctx, ctx.byTitleMovieIDs, movie.MovieSearchEntryType_MOVIE_SEARCH_ENTRY_TYPE_TITLE)
	s.movie2SearchEntry(ctx, ctx.byDetailMovieIDs, movie.MovieSearchEntryType_MOVIE_SEARCH_ENTRY_TYPE_DETAIL)
}

func (*SearchService) movie2SearchEntry(ctx *SearchServiceContext,
	movies []string, msType movie.MovieSearchEntryType) {
	for _, id := range movies {
		modelMovie := ctx.movieID2Movie[id]
		movieSearchEntry := &movie.MovieSearchEntry{
			MsEntryType: msType,
			Movie: &movie.Movie{
				Id:            modelMovie.Id,
				Title:         modelMovie.Title,
				PicUrl:        modelMovie.PicUrl,
				Introduction:  &modelMovie.Introduction,
				ReleaseDate:   &modelMovie.ReleaseDate,
				Language:      &modelMovie.Language,
				AverageRating: &modelMovie.AverageRating,
			},
		}
		for _, participant := range modelMovie.Participants {
			movieSearchEntry.Movie.Participant = append(movieSearchEntry.Movie.Participant, &movie.Participant{
				Character: participant.Character,
				Name:      participant.Name,
			})
		}
		ctx.Resp.Movies = append(ctx.Resp.Movies, movieSearchEntry)
	}
}
