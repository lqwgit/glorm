// Code generated by model_gen
// comment_test.go contains model test for the database table [blog.comment]

package examples

import (
	"github.com/zhengyun1112/glorm/orm"
	"testing"
	"time"
)

func TestInsertAndGetComment(t *testing.T) {
	orm.InitDefault("root:@/blog?parseTime=true&loc=Local")
	orm.TruncateTable("comment")

	var comment = &Comment{
		UserId:    0,
		ArticleId: 0,
		Content:   "",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := CommentDao.Insert(comment)
	if err != nil {
		t.Fatalf("failed to InsertComment, err: %+v", err)
	}
	loaded, err := CommentDao.GetByPK(comment.CommentId)
	if err != nil {
		t.Fatalf("failed to GetByPK, err: %+v", err)
	}
	if loaded == nil {
		t.Fatalf("should have loaded one Comment")
	}
}
