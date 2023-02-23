package service

import (
	"context"
	"offer_tiktok/biz/dal/db"
	"offer_tiktok/biz/model/social/message"
	"offer_tiktok/pkg/errno"
	"offer_tiktok/pkg/utils"

	"github.com/cloudwego/hertz/pkg/app"
)

type MessageService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewMessageService(ctx context.Context, c *app.RequestContext) *MessageService {
	return &MessageService{ctx: ctx, c: c}
}

func (m *MessageService) GetMessageChat(req *message.DouyinMessageChatRequest) ([]*message.Message, error) {
	messages := make([]*message.Message, 0)
	from_user_id, _ := m.c.Get("current_user_id")
	to_user_id := req.ToUserId
	pre_msg_time := req.PreMsgTime
	db_messages, err := db.GetMessageByIdPair(from_user_id.(int64), to_user_id, utils.MillTimeStampToTime(pre_msg_time))
	if err != nil {
		return messages, err
	}
	for _, db_message := range db_messages {
		messages = append(messages, &message.Message{
			Id:         db_message.ID,
			ToUserId:   db_message.ToUserId,
			FromUserId: db_message.FromUserId,
			Content:    db_message.Content,
			CreateTime: db_message.CreatedAt.UnixNano() / 1000000,
		})
	}
	return messages, nil
}

func (m *MessageService) MessageAction(req *message.DouyinMessageActionRequest) error {
	from_user_id, _ := m.c.Get("current_user_id")
	to_user_id := req.ToUserId
	// action_type := req.ActionType
	content := req.Content

	ok, err := db.AddNewMessage(&db.Messages{
		FromUserId: from_user_id.(int64),
		ToUserId:   to_user_id,
		Content:    content,
	})
	if err != nil {
		return err
	}
	if !ok {
		return errno.MessageAddFailedErr
	}
	return nil
}
