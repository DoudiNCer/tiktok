package mysql

import "github.com/DodiNCer/tiktok/biz/model"

// QueryFavoriteCount 获取用户的总点赞数
func QueryFavoriteCount(userId int64) (int64, error) {
	var TotalFavorite int64
	err := DB.Model(&model.Favorite{}).Where(model.Favorite{CreatorId: userId}).Where("is_deleted=0").Count(&TotalFavorite).Error
	return TotalFavorite, err
}

// QueryTotalFavorited 获取用户的总被点赞数
func QueryTotalFavorited(userId int64) (int64, error) {
	sql := "SELECT COUNT(favorite.id) FROM favorite " +
		"INNER JOIN(SELECT video.id FROM video WHERE video.creator_id= ? AND is_deleted = 0)AS videos " +
		"ON favorite.video_id = videos.id AND favorite.is_deleted = 0 ;"
	var totalFavorited int64
	err := DB.Raw(sql, userId).Scan(&totalFavorited).Error
	if err != nil {
		return 0, err
	}
	return totalFavorited, err

}
