// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNamePostThumb = "post_thumb"

// PostThumb 帖子点赞
type PostThumb struct {
	ID         int64     `gorm:"column:id;primaryKey;autoIncrement:true;comment:id" json:"id"`                          // id
	CreateTime time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"update_time"` // 更新时间
	PostID     int64     `gorm:"column:post_id;not null;comment:帖子 id" json:"post_id"`                                  // 帖子 id
	UserID     int64     `gorm:"column:user_id;not null;comment:创建用户 id" json:"user_id"`                                // 创建用户 id
}

// TableName PostThumb's table name
func (*PostThumb) TableName() string {
	return TableNamePostThumb
}
