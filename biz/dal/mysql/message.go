package mysql

import "github.com/DodiNCer/tiktok/biz/model"

// 插入新消息
func CreateMessage(messages []*model.Message) error {
	return DB.Create(messages).Error
}

// 获取与某用户有关的消息
func QueryMessage(uid int64) ([]*model.Message, error) {
	db := DB.Model(model.Message{})
	db = db.Where("listener_id = ? or reciver_id = ?", uid, uid)

	var res []*model.Message
	if err := db.Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// 获取两人交流的最后一条消息
func QueryLastMessage(userId, friendId int64) (message *model.Message, msgType int64, err error) {
	db := DB.Model(model.Message{})
	db = db.Where("(reciver_id = ? AND listener_id = ?) OR (reciver_id = ? AND listener_id = ?)", userId, friendId, friendId, userId)
	var msg *model.Message
	if err = db.Last(&msg).Error; err != nil {
		return nil, 0, err
	}
	msgType = 0
	// 无消息
	if msg == nil {
		return nil, 0, nil
	}
	// 当前用户不是接收者
	if msg.Listener_Id != userId {
		msgType = 1
	}
	return msg, msgType, err
}
