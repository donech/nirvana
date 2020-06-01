package service

func NewSimpleService() *SimpleService {
	return &SimpleService{}
}

type SimpleService struct {
	//db *gorm.DB
}

func (s SimpleService) ItemByID(ID int64) interface{} {
	return struct {
		ID int64
	}{ID: ID}
}
