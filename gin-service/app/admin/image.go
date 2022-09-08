package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/putyy/ai-share/app/api"
	"github.com/putyy/ai-share/app/form"
	"github.com/putyy/ai-share/app/library"
	"github.com/putyy/ai-share/app/models"
)

func ImageList(c *gin.Context) {
	data := make(map[string]interface{})
	var imgList []models.ImageMaterial
	model := models.Db()

	var total int64

	if c.Query("title") != "" {
		model = model.Where("title LIKE ?", "%"+c.Query("title")+"%")
	}

	model.Scopes(models.PaginateScope(c)).
		Where("is_del = 1").
		Order("id desc").
		Find(&imgList)

	if total == 0 {
		model.Model(models.ImageMaterial{}).Count(&total)
	}

	for index, list := range imgList {
		imgList[index].ImgUrl = library.GetImgUrl(list.ImgUrl)
	}

	data["total"] = total
	data["page"] = c.Query("page")
	data["page_size"] = c.Query("page_size")
	data["list"] = imgList

	api.ResponseSuccess(c, data)
}

func ImageAdd(c *gin.Context) {

	var formData form.ImageAdminForm
	if err1 := c.ShouldBind(&formData); err1 != nil {
		api.ResponseError(c, "参数错误", err1.Error())
		return
	}

	model := models.ImageMaterial{
		Title:  formData.Title,
		ImgUrl: library.GetSaveUrl(formData.ImgUrl),
	}

	var err error
	if formData.Id > 0 {
		err = models.Db().Model(&model).Where("id = ?", formData.Id).Updates(model).Error
		model.ID = formData.Id
	} else {
		err = models.Db().Create(&model).Error
	}
	if err != nil {
		api.ResponseError(c, "创建失败", err.Error())
		return
	}

	api.ResponseSuccess(c, model)
}

func ImageInfo(c *gin.Context) {
	var ImgInfo models.ImageMaterial
	models.Db().First(&ImgInfo, c.Query("id"))
	ImgInfo.ImgUrl = library.GetImgUrl(ImgInfo.ImgUrl)
	api.ResponseSuccess(c, ImgInfo)
}
