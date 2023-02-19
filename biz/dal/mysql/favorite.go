package mysql

import "github.com/DodiNCer/tiktok/biz/model"

// QueryFavoriteCount 获取用户的总点赞数
func QueryFavoriteCount(userId int64) (int64, error) {
	var TotalFavorite int64
	err := DB.Model(&model.Favorite{}).Where(model.Favorite{CreatorId: userId}).Where("is_deleted=0").Count(&TotalFavorite).Error
	return TotalFavorite, err
}
