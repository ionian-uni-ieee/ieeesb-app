package testing

// DatabaseSession contains a database driver session with unexported variables and driver interface functions
type DatabaseSession struct {
	Session *Session
}

// Session contains a set of collections/tables
// which simulate a database's storage
type Session struct {
	Collections map[string]*Collection
}

// Collection contains a set of columns that represent a map
// of fields with their corresponding data
type Collection struct {
	Columns map[string][]interface{}
}

// MakeDatabaseDriver builds a mongo database driver structure
func MakeDatabaseDriver() *DatabaseSession {
	d := &DatabaseSession{
		Session: &Session{
			Collections: make(map[string]*Collection),
		},
	}

	return d
}

// Connect is an empty function to fill the Driver interface's requirements
func (d *DatabaseSession) Connect(host string, port string, uname string, pass string, dbname string) error {
	return nil
}

// Disconnect is an empty function to fill the interface's requirements
func (d *DatabaseSession) Disconnect() {

}

// GetDatabaseType describes what type this database driver is
// returns "testing"
func (d *DatabaseSession) GetDatabaseType() string {
	return "testing"
}

// GetClient returns an controlling interface for in-memory
// database controlling
func (d *DatabaseSession) GetClient() interface{} {
	return nil
}

// GetDatabase returns the database
func (d *DatabaseSession) GetDatabase() interface{} {
	return d.Session
}

// GetCollection returns a database's collection
func (d *DatabaseSession) GetCollection(collectionName string) interface{} {
	collection := d.Session.Collections[collectionName]

	if collection == nil {
		d.Session.Collections[collectionName] = &Collection{}
		d.Session.Collections[collectionName].Columns = make(map[string][]interface{})
		collection = d.Session.Collections[collectionName]
	}

	return collection
}
