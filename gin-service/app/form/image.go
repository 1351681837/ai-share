package form

type ImageAdminForm struct {
	Id     int    `form:"id"`
	Title  string `form:"title" binding:"required"`
	ImgUrl string `form:"img_url" binding:"required"`
	IsDel  int `form:"is_del"`
}
