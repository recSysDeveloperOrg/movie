package service

import (
	"context"
	"fmt"
	. "movie/constant"
	"movie/idl/gen/movie"
	"movie/rpc"
	"strings"
	"sync"
)

type RecommendFeedbackService struct {
}

type RecommendFeedbackContext struct {
	Ctx     context.Context
	Req     *movie.FeedbackReq
	Resp    *movie.FeedbackResp
	ErrCode *ErrorCode
}

var recommendFeedbackService *RecommendFeedbackService
var recommendFeedbackServiceOnce sync.Once

var validFeedbackTypeSet = map[movie.FeedbackType]struct{}{
	movie.FeedbackType_FEEDBACK_TYPE_MOVIE:    {},
	movie.FeedbackType_FEEDBACK_TYPE_CATEGORY: {},
}

func NewRecommendFeedbackService() *RecommendFeedbackService {
	recommendFeedbackServiceOnce.Do(func() {
		recommendFeedbackService = &RecommendFeedbackService{}
	})

	return recommendFeedbackService
}

func NewRecommendFeedbackContext(ctx context.Context, req *movie.FeedbackReq) *RecommendFeedbackContext {
	return &RecommendFeedbackContext{
		Ctx: ctx,
		Req: req,
		Resp: &movie.FeedbackResp{
			BaseResp: &movie.BaseResp{},
		},
	}
}

func (s *RecommendFeedbackService) Feedback(ctx *RecommendFeedbackContext) {
	defer s.buildResponse(ctx)
	if s.checkParams(ctx); ctx.ErrCode != nil {
		return
	}
	if s.sendFeedback(ctx); ctx.ErrCode != nil {
		return
	}
}

func (*RecommendFeedbackService) checkParams(ctx *RecommendFeedbackContext) {
	if strings.TrimSpace(ctx.Req.SourceId) == "" {
		ctx.ErrCode = BuildErrCode("movie/tag id is empty", RetParamsErr)
		return
	}
	if strings.TrimSpace(ctx.Req.UserId) == "" {
		ctx.ErrCode = BuildErrCode("user id is empty", RetParamsErr)
		return
	}
	if _, ok := validFeedbackTypeSet[ctx.Req.FbType]; !ok {
		ctx.ErrCode = BuildErrCode(fmt.Sprintf("invalid feedback type:%+v", ctx.Req.FbType), RetParamsErr)
		return
	}
}

func (*RecommendFeedbackService) sendFeedback(ctx *RecommendFeedbackContext) {
	if _, err := rpc.SendFeedback(ctx.Ctx, ctx.Req.UserId, ctx.Req.SourceId, ctx.Req.FbType); err != nil {
		ctx.ErrCode = BuildErrCode(err, RetRpcCallFailed)
		return
	}
}

func (*RecommendFeedbackService) buildResponse(ctx *RecommendFeedbackContext) {
	errCode := RetSuccess
	if ctx.ErrCode != nil {
		errCode = ctx.ErrCode
	}

	ctx.Resp.BaseResp.ErrNo, ctx.Resp.BaseResp.ErrMsg = errCode.Code, errCode.Msg
}
