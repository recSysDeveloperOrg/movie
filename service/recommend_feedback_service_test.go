package service

import (
	"context"
	"movie/idl/gen/movie"
	"testing"
)

func TestRecommendFeedbackService_Feedback(t *testing.T) {
	ctx := NewRecommendFeedbackContext(context.Background(), &movie.FeedbackReq{
		UserId:   newUser,
		SourceId: "100000000000000000000004",
		FbType:   movie.FeedbackType_FEEDBACK_TYPE_MOVIE,
	})
	NewRecommendFeedbackService().Feedback(ctx)

	ctx.Req.FbType = movie.FeedbackType_FEEDBACK_TYPE_CATEGORY
	NewRecommendFeedbackService().Feedback(ctx)
}
