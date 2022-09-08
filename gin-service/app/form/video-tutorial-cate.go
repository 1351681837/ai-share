package form

type VideoTutorialCateFrom struct {
	Id     int64  `form:"id"`
	Name   string `form:"name" binding:"required"`
	Sort   int64  `form:"sort,default=0"`
	IsLock int8   `form:"is_lock"`
}
