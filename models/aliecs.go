package models

import (
	"github.com/jinzhu/gorm"
	"github.com/luyaops/cmdb/database"
)

// Model
type AliEcs struct {
	gorm.Model
	InstanceId     string `gorm:"unique not null VARCHAR(128)"`
	InstanceName   string `gorm:"VARCHAR(191)"`
	InnerIpAddress string `gorm:"not null VARCHAR(128)"`
}

// Method

func (ecs *AliEcs) GetEcsById() *AliEcs {
	database.DB.First(ecs)
	return ecs
}
