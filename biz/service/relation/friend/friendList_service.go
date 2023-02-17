package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"

	relation "offer_tiktok/biz/model/social/relation"
)

type FriendListService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewFriendListService(ctx context.Context, c *app.RequestContext) *FriendListService {
	return &FriendListService{ctx: ctx, c: c}
}

func (s *FriendListService) FriendList(req *relation.DouyinRelationFriendListRequest) (*relation.DouyinRelationFriendListResponse, error) {
	return &relation.DouyinRelationFriendListResponse{}, nil
}
