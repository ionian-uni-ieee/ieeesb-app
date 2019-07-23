package database

// Driver decribes what a database driver should look and behave like
type Driver interface {
	Connect(host string, port string, uname string, pass string, dbname string) error
	Disconnect()
	GetDatabaseType() string
	GetClient() interface{}
	GetDatabase() interface{}
}
