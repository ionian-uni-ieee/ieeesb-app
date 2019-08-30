package mongo

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// DatabaseSession contains a database driver session with unexported variables and driver interface functions
type DatabaseSession struct {
	client   *mongo.Client
	database *mongo.Database
	host     string
	port     string
	uname    string
	pass     string
	dbname   string
}

// MakeDatabaseDriver builds a mongo database driver structure
func MakeDatabaseDriver() *DatabaseSession {
	d := &DatabaseSession{}

	return d
}

// GetDatabaseType returns the name of the DBMS that is used
func (d *DatabaseSession) GetDatabaseType() string {
	return "mongo"
}

// GetClient returns an instance of a database client
func (d *DatabaseSession) GetClient() interface{} {
	return d.client
}

// GetDatabase returns an instance of a database
func (d *DatabaseSession) GetDatabase() interface{} {
	return d.database
}

// GetCollection returns a database's collection
func (d *DatabaseSession) GetCollection(collectionName string) interface{} {
	db := d.GetDatabase().(*mongo.Database)
	collection := db.Collection(collectionName)
	return collection
}

// Connect connects to the mongodb database
func (d *DatabaseSession) Connect(host string, port string, uname string, pass string, dbname string) error {
	// Assign params to object
	d.host = host
	d.port = port
	d.uname = uname
	d.pass = pass
	d.dbname = dbname

	// Check if a database already exists
	if d.database != nil {
		return errors.New("A database session already exists")
	}

	// New Client
	clientOptions := &options.ClientOptions{}
	client, err := mongo.NewClient(clientOptions.ApplyURI(d.host + ":" + d.port))
	if err != nil {
		return err
	}
	d.client = client

	// Connect
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return err
	}

	// Verify connection topology
	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return errors.New("Couldn't ping to the database\nCheck your database service or your URI options (PORT, HOST)")
	}

	// Checking if database exists
	dbnames, err := client.ListDatabaseNames(context.Background(), &bson.M{})
	if err != nil {
		return err
	}
	found := false
	for _, dbname := range dbnames {
		if dbname == d.dbname {
			found = true
		}
	}
	if !found {
		return errors.New("Couldn't find database with name " + d.dbname)
	}

	// Assign database
	database := client.Database(d.dbname)
	d.database = database
	return err
}

// Disconnect disconnects from the database session
func (d *DatabaseSession) Disconnect() {
	d.client.Disconnect(context.Background())
	d.database = nil
}
