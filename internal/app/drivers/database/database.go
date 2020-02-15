package database

// Driver decribes what a database driver should look and behave like
type Driver interface {
	Connect(origin string, uname string, pass string, dbname string) error
	Disconnect()
	StartSession() (interface{}, error)
	GetClient() interface{}
	GetCollection(collectionName string) interface{}
	GetDatabase() interface{}
	GetDatabaseType() string
}
