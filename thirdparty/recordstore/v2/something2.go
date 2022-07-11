package v2

type Service struct{}

type Table struct {
	name string
}

func New() Service {
	return Service{}
}

func (s Service) GetTable(name string) Table {
	return Table{name: name}
}

func (t Table) StoreRecord(rec interface{}) {

}
