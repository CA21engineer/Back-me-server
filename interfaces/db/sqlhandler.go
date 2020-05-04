package db

type SqlHandler interface {
	SelectInt(query string, args ...interface{}) (int64, error)
	SelectOne(holder interface{}, query string, args ...interface{}) error
	Select(holder interface{}, query string, args ...interface{}) ([]interface{}, error)
	Insert(list ...interface{}) error
	Update(list ...interface{}) (int64, error)
	Delete(list ...interface{}) (int64, error)
}
