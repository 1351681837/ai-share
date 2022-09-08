package form

type VideoTutorialFrom struct {
	Id       int  `form:"id"`
	CateId   int    `form:"cate_id"`
	Title    string `form:"title" binding:"required"`
	CoverUrl string `form:"cover_url"`
	VideoSrc string `form:"video_src"`
	Sort     int    `form:"sort,default=0"`
	IsLock   int    `form:"is_lock"`
	IsDel    int    `form:"is_del"`
}
