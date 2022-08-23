package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/putyy/ai-share/app/api"
	"github.com/putyy/ai-share/app/models"
)

func ImageList(c *gin.Context) {
	data := make(map[string]interface{})
	var imgList []models.ImageMaterial
	model := models.Db()

	var total int64

	model.Scopes(models.PaginateScope(c)).
		Where("is_del = 1").
		Order("id desc").
		Find(&imgList)

	if total == 0 {
		model.Model(models.ImageMaterial{}).Count(&total)
	}

	data["total"] = total
	data["page"] = c.Query("page")
	data["page_size"] = c.Query("page_size")
	data["list"] = imgList

	api.ResponseSuccess(c, data)
}
