package mysql

import (
	"github.com/DodiNCer/tiktok/biz/model"
)

func QueryCommentByCommentId(commentId int64) (comment *model.Comment, err error) {
	err = DB.Where(model.Comment{Id: commentId}).Where("is_deleted = 0").Find(&comment).Error
	return comment, nil
}

func CreateComment(comment *model.Comment) (*model.Comment, error) {
	err := DB.Create(comment).Error
	return comment, err
}

func DeleteComment(commentId int64) error {
	return DB.Model(model.Comment{}).Where(model.Comment{Id: commentId}).Updates(map[string]interface{}{"is_deleted": 1}).Error
}

func QueryCommentsByVideoId(videoId int64) ([]*model.Comment, error) {
	var comments []*model.Comment
	err := DB.Model(model.Comment{}).Where("video_id = ? AND is_deleted = 0", videoId).Find(&comments).Error
	return comments, err
}

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

// QueryWorkCount 获取用户作品数量
func QueryWorkCount(userId int64) (int64, error) {
	sql := "SELECT COUNT(video.id) FROM video WHERE video.creator_id=?"
	var count int64
	err := DB.Raw(sql, userId).Scan(&count).Error
	return count, err
}
