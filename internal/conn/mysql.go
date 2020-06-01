package conn

import (
	"fmt"

	"github.com/donech/nirvana/internal/config"
	"github.com/donech/tool/xdb"
	"github.com/jinzhu/gorm"
)

type NirvanaDB *gorm.DB

func NewNirvanaDB(conf *config.Config) (NirvanaDB, func()) {
	fmt.Println(conf)
	DB, clean := xdb.Open(conf.NirvanaDB)
	return DB, clean
}
