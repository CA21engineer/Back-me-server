package db

type SqlHandler interface {
	SelectOne(holder interface{}, query string, args ...interface{}) error
	Select(holder interface{}, query string, args ...interface{}) ([]interface{}, error)
	Insert(list ...interface{}) error
	Update(list ...interface{}) (int64, error)
	Delete(list ...interface{}) (int64, error)
}
