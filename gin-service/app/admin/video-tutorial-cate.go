package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/putyy/ai-share/app/api"
	"github.com/putyy/ai-share/app/form"
	"github.com/putyy/ai-share/app/models"
)

func VideoTutorialCateList(c *gin.Context) {
	model := models.Db()
	var list []models.VideoTutorialCate
	data := make(map[string]interface{})
	var total int64

	if c.Query("title") != "" {
		model = model.Where("title like ?", "%"+c.Query("title")+"%")
	}

	if c.Query("start_time") != "" {
		model = model.Where("create_time like >= ?", c.Query("start_time"))
	}

	if c.Query("end_time") != "" {
		model = model.Where("create_time like <= ?", c.Query("end_time"))
	}

	model.Scopes(models.PaginateScope(c)).
		Order("id desc").
		Find(&list)

	if total == 0 {
		model.Model(models.VideoTutorialCate{}).Count(&total)
	}

	data["total"] = total
	data["page"] = c.Query("page")
	data["page_size"] = c.Query("page_size")
	data["list"] = list
	data["m_mark"] = "video-tutorial-cate"

	api.ResponseSuccess(c, data)
}

func VideoTutorialCateInfo(c *gin.Context) {
	var info models.VideoTutorialCate
	models.Db().First(&info, c.Query("id"))
	api.ResponseSuccess(c, info)
}

type AllName struct {
	ID    int    `json:"id"`
	Name string `json:"name"`
}

func VideoTutorialCateAllName(c *gin.Context) {
	var list []AllName
	models.Db().Model(&models.VideoTutorialCate{}).Find(&list)
	api.ResponseSuccess(c, list)
}

func VideoTutorialCateEdit(c *gin.Context) {
	var formData form.VideoTutorialCateFrom
	if err1 := c.ShouldBind(&formData); err1 != nil {
		api.ResponseError(c, "参数错误", err1.Error())
		return
	}

	model := models.VideoTutorialCate{
		Name: formData.Name,
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
