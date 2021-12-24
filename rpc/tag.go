package rpc

import (
	"context"
	"errors"
	"movie/constant"
	"movie/idl/gen/tag"
)

func QueryTagByTagID(ctx context.Context, tagIDs []string) (*tag.QueryTagByTagIDResp, error) {
	req := &tag.QueryTagByTagIDReq{
		TagId: tagIDs,
	}
	resp, err := TagClient.QueryTagByTagID(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp == nil || resp.BaseResp.ErrNo != constant.RetSuccess.Code {
		return nil, errors.New("retNo not 0")
	}

	return resp, nil
}
