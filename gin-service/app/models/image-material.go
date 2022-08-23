package models

type ImageMaterial struct {
	ID        int      `gorm:"primary_key" json:"id,omitempty"`
	Title     string   `json:"title"`
	ImgUrl    string   `json:"img_url,omitempty"`
	IsDel     int      `json:"is_del" gorm:"default:1"`
	CreatedAt JsonTime `json:"created_at,omitempty"`
	UpdatedAt JsonTime `json:"updated_at,omitempty"`
}

func (ImageMaterial) TableName() string {
	return "a_image_material"
}
