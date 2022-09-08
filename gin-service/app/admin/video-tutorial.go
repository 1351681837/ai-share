package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/putyy/ai-share/app/api"
	"github.com/putyy/ai-share/app/form"
	"github.com/putyy/ai-share/app/library"
	"github.com/putyy/ai-share/app/models"
)

type VideoTutorial struct {
	models.VideoTutorial
	Cate  models.VideoTutorialCate `json:"wallet Cate"  gorm:"foreignKey:CateId;references:Id"`
}

func VideoTutorialList(c *gin.Context) {
	var list []VideoTutorial
	var total int64
	model := models.Db()
	data := make(map[string]interface{})

	if c.Query("title") != "" {
		model = model.Where("title like ? ", "%"+c.Query("title")+"%")
	}

	if c.Query("start_time") != "" {
		model = model.Where("create_time like >= ?", c.Query("start_time"))
	}

	if c.Query("end_time") != "" {
		model = model.Where("create_time like <= ?", c.Query("end_time"))
	}

	model.Scopes(models.PaginateScope(c)).
		Preload("Cate").
		Order("id desc").
		Find(&list)

	if total == 0 {
		models.Db().Model(models.VideoTutorial{}).Count(&total)
	}


	for index, item := range list {
		list[index].CoverUrl = library.GetImgUrl(item.CoverUrl)
	}


	data["total"] = total
	data["page"] = c.Query("page")
	data["page_size"] = c.Query("page_size")
	data["list"] = list
	data["m_mark"] = "video-tutorial"

	api.ResponseSuccess(c, data)
	return
}

func VideoTutorialInfo(c *gin.Context) {
	var info models.VideoTutorial
	models.Db().Preload("Cate").First(&info, c.Query("id"))
	info.CoverUrl = library.GetImgUrl(info.CoverUrl)
	api.ResponseSuccess(c, info)
}

func VideoTutorialEdit(c *gin.Context) {
	var formData form.VideoTutorialFrom
	if err1 := c.ShouldBind(&formData); err1 != nil {
		api.ResponseError(c, "参数错误", err1.Error())
		return
	}

	model := models.VideoTutorial{
		CateId: formData.CateId,
		Title: formData.Title,
		CoverUrl: library.GetSaveUrl(formData.CoverUrl),
		VideoSrc: library.GetSaveUrl(formData.VideoSrc),
		Sort: formData.Sort,
	}

	var err error
	if formData.Id > 0 {
		err = models.Db().Model(&model).Where("id = ?", formData.Id).Updates(model).Error
		model.Id = formData.Id
	} else {
		err = models.Db().Create(&model).Error
	}
	if err != nil {
		api.ResponseError(c, "创建失败", err.Error())
		return
	}

	api.ResponseSuccess(c, model)
}
