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
	err := DB.Model(model.Comment{}).Where("video_id = ? AND is_deleted = 0", videoId).Order("create_time desc").Find(&comments).Error
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
	sql := "SELECT COUNT(video.id) FROM video WHERE video.creator_id=? AND is_deleted = 0"
	var count int64
	err := DB.Raw(sql, userId).Scan(&count).Error
	return count, err
}

// QueryPortraitPathByUserId 根据用户ID查询用户头像

func QueryPortraitPathByUserId(userId int64) (string, error) {
	sql := "SELECT `user`.portrait_path FROM user WHERE `user`.id = ? ;"
	var PortraitPath string
	err := DB.Raw(sql, userId).Scan(&PortraitPath).Error
	return PortraitPath, err
}

// QueryBackgroundImageByUserId 根据用户ID获取用户个人页顶部大图
func QueryBackgroundImageByUserId(userId int64) (string, error) {
	sql := "SELECT `user`.background_picture_path FROM user WHERE `user`.id = ? ;"
	var backgroundImage string
	err := DB.Raw(sql, userId).Scan(&backgroundImage).Error
	return backgroundImage, err
}

// QuerySignatureByUserId 根据用户ID获取用户个人简介
func QuerySignatureByUserId(userId int64) (string, error) {
	sql := "SELECT `user`.signature FROM user WHERE `user`.id = ?;"
	var signature string
	err := DB.Raw(sql, userId).Scan(&signature).Error
	return signature, err
}

func QueryCommentCountByVideo(videoId int64) (int64, error) {
	var commentCount int64
	err := DB.Model(&model.Comment{}).Where("video_id = ? AND is_deleted = ?", videoId, 0).
		Count(&commentCount).Error
	return commentCount, err
}
