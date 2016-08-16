// Code generated by model_gen
// user_test.go contains model test for the database table [blog.user]

package examples

import (
	"github.com/zhengyun1112/glorm/orm"
	"testing"
	"time"
)

func TestInsertAndGetUser(t *testing.T) {
	orm.InitDefault("root:@/blog?parseTime=true&loc=Local")
	orm.TruncateTable("user")

	var user = &User{
		Name:      "",
		Password:  "",
		IsMarried: 0,
		Age:       0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := UserDao.Insert(user)
	if err != nil {
		t.Fatalf("failed to InsertUser, err: %+v", err)
	}
	loaded, err := UserDao.GetByPK(user.UserId)
	if err != nil {
		t.Fatalf("failed to GetByPK, err: %+v", err)
	}
	if loaded == nil {
		t.Fatalf("should have loaded one User")
	}
}
