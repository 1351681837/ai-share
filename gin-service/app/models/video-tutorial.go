package models

type VideoTutorial struct {
	Id        int      `gorm:"primaryKey" json:"id,omitempty"`
	CateId    int      `json:"cate_id"`
	Title     string   `json:"title"`
	CoverUrl  string   `json:"cover_url"`
	VideoSrc  string   `json:"video_src"`
	Sort      int      `json:"sort"`
	IsLock    int      `json:"is_lock"`
	IsDel     int      `json:"is_del"`
	CreatedAt JsonTime `json:"created_at"`
	UpdatedAt JsonTime `json:"updated_at"`
}

func (VideoTutorial) TableName() string {
	return "a_video_tutorial"
}
