package recordstore

type Service struct{}

func New() Service {
	return Service{}
}

func (s Service) CreateRecord() interface{} {
	return nil
}

func (s Service) StoreRecord(rec interface{}) {
}
