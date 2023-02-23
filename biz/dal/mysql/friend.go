package mysql

import "github.com/DodiNCer/tiktok/biz/model"

// 获取好友信息
func QueryFriend(uid int64) (res []*model.User, err error) {
	sql := "SELECT u.id AS 'id', u.`name` AS 'name'" +
		"FROM follower AS f LEFT JOIN follower AS t ON f.to_user_uid = t.user_uid " +
		"LEFT JOIN `user` AS u ON f.to_user_uid = u.id " +
		"WHERE (f.is_deleted = 0 AND t.is_deleted = 0) " +
		"AND (f.user_uid = ? AND t.to_user_uid = ?)"
	db := DB.Raw(sql, uid, uid)
	err = db.Error
	if err != nil {
		return nil, err
	}
	db.Scan(&res)
	return
}

// 获取粉丝人数
func QueryFollowerNum(uid int64) (int64, error) {
	var count int64
	db := DB.Model(model.Follower{})
	db = db.Where("is_deleted = 0 AND to_user_uid = ?", uid).Count(&count)
	return count, nil
}

// 获取关注人数
func QueryFollowNum(uid int64) (int64, error) {
	var count int64
	db := DB.Model(model.Follower{})
	db = db.Where("is_deleted = 0 AND user_uid = ?", uid).Count(&count)
	return count, nil
}
