package models

type VideoTutorialCate struct {
	Id        int64    `gorm:"primaryKey" json:"id"`
	Name      string   `json:"name"`
	Sort      int64    `json:"sort"`
	IsLock    int8     `gorm:"default:1" json:"is_lock"`
	CreatedAt JsonTime `json:"created_at"`
	UpdatedAt JsonTime `json:"updated_at"`
}

func (VideoTutorialCate) TableName() string {
	return "a_video_tutorial_cate"
}
