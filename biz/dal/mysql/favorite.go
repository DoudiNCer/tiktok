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

// 查询用户所有的点赞信息
func QueryFavoritesByCreatorId(creatorId int64) ([]*model.Favorite, error) {
	var favorites []*model.Favorite
	err := DB.Where("creator_id = ? AND is_deleted = ?", creatorId, 0).
		Select("video_id").Find(favorites).Error
	return favorites, err
}

// 查找视频点赞总数
func QueryFavoriteNumByVideo(videoId int64) (int64, error) {
	var favoriteCount int64
	err := DB.Model(&model.Favorite{}).Where("video_id = ? AND is_deleted = ?", videoId, 0).
		Count(&favoriteCount).Error
	return favoriteCount, err
}

func QueryFavoriteIs(creatorId int64, videoId int64) (int64, error) {
	var isbool int64
	err := DB.Model(&model.Favorite{}).Where("creator_id = ? AND video_id = ? AND is_deleted = ?", creatorId, videoId, 0).
		Count(&isbool).Error
	return isbool, err
}
