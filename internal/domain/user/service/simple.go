package service

import (
	"context"

	"github.com/donech/nirvana/internal/conn"

	"github.com/donech/core/xtrace"
	"github.com/donech/tool/xdb"
	"github.com/jinzhu/gorm"
)

type BaseApps struct {
	AppId        string `json:"app_id"`
	AppName      string `json:"app_name"`
	DebugMode    int    `json:"debug_mode"`
	AppConfig    string `json:"app_config"`
	Status       string `json:"status"`
	Webpath      string `json:"webpath"`
	Description  string `json:"description"`
	LocalVer     string `json:"local_ver"`
	RemoteVer    string `json:"remote_ver"`
	AuthorName   string `json:"author_name"`
	AuthorUrl    string `json:"author_url"`
	AuthorEmail  string `json:"author_email"`
	Dbver        string `json:"dbver"`
	RemoteConfig string `json:"remote_config"`
}

func NewSimpleService(db conn.NirvanaDB) *SimpleService {
	return &SimpleService{db: db}
}

type SimpleService struct {
	db *gorm.DB
}

func (s SimpleService) ItemByID(ID int64) interface{} {
	data := BaseApps{}
	xdb.Trace(xtrace.NewCtxWithTraceID(context.Background()), s.db).Table("base_apps").Where("app_id = ?", "base").First(&data)
	return data
}
