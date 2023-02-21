package mysql

import "github.com/DodiNCer/tiktok/biz/model"

// 查询用户点赞视频数
func QueryNumOfVideoFavoriteByUser(uid int64) int64 {
	var count int64
	db := DB.Model(&model.Favorite{})
	db.Where("favorite.creator_id = ? AND is_deleted != TRUE", uid).Count(&count)
	return count
}

// 查询用户视频被点赞数
func QueryNumOfFavoriteGotByUser(uid int64) (count int64, err error) {
	sql := "SELECT COUNT(1) FROM video " +
		"LEFT JOIN favorite ON video.id = favorite.video_id " +
		"WHERE video.creator_id = ? " +
		"AND video.is_deleted != TRUE AND favorite.is_deleted != TRUE"
	db := DB.Raw(sql, uid)
	err = db.Error
	if err != nil {
		return 0, err
	}
	db.Scan(count)
	return
}
