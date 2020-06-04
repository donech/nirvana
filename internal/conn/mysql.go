package conn

import (
	"github.com/donech/nirvana/internal/config"
	"github.com/donech/tool/xdb"
	"github.com/jinzhu/gorm"
)

type NirvanaDB *gorm.DB

func NewNirvanaDB(conf *config.Config) (NirvanaDB, func()) {
	DB, clean := xdb.Open(conf.NirvanaDB)
	return DB, clean
}
