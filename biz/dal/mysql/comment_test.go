package mysql

import (
	"fmt"
	"github.com/DodiNCer/tiktok/biz/model"
	"testing"
)

func TestCreateComment(t *testing.T) {
	fmt.Println("CreateComment")
	comment := model.Comment{CreatorUid: 12, Text: "hello", VideoId: 12, IsDeleted: 0}
	Init()
	_, err := CreateComment(&comment)
	if err != nil {
		return
	}
}

func TestQueryCommentByCommentId(t *testing.T) {
	fmt.Println("QueryCommentByCommentId")
	Init()
	tmp, _ := QueryCommentByCommentId(31)
	fmt.Println(tmp)
}

func TestDeleteComment(t *testing.T) {
	fmt.Println("delete")
	Init()
	err := DeleteComment(1)
	if err != nil {
		return
	}
}

func TestQueryCommentsByVideoId(t *testing.T) {
	fmt.Println("TestQueryCommentsByVideoId")
	Init()
	comments, err := QueryCommentsByVideoId(2)
	if err != nil {
		return
	}

	for _, comment := range comments {
		fmt.Println(comment)
	}
}
