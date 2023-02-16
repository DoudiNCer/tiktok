package mysql

import "github.com/DodiNCer/tiktok/biz/model"

func QueryCommentByCommentId(commentId int64) (comment *model.Comment, err error) {
	err = DB.Where(model.Comment{Id: commentId}).Find(&comment).Error
	return comment, nil
}

func CreateComment(comment *model.Comment) (*model.Comment, error) {
	err := DB.Create(comment).Error
	return comment, err
}

func DeleteComment(commentId int64) error {
	return DB.Where(model.Comment{Id: commentId}).Updates(model.Comment{IsDeleted: 1}).Error
}
