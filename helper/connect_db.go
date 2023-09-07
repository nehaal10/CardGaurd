package helper

import (
	"github.com/nehaal10/CardGaurd/controller"
	"gorm.io/gorm"
)

var DB *gorm.DB = controller.Init()
