package publicModels

type Model struct {
	ModelID    uint `gorm:"column:model_id" json:"id"`
	MaterialID uint `gorm:"column:material_id" json:"material_id"`
}
