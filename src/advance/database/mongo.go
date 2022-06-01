package database

import (
	"Go-Fundamentals/src/webapp/entity"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getDBConfig() *entity.DBConfig {
	config := &entity.DBConfig{}
	config.HostName = "localhost"
	config.Port = "27017"
	config.DBName = "target"
	config.UserName = ""
	config.Password = ""
	return config
}

// MongoDB mongo database
type MongoDB struct {
}

// GetConn gives mongo connection object
func (db MongoDB) GetConn(ctx context.Context) (*entity.DBConn, error) {
	config := getDBConfig()
	client, err := db.getMongoClient(config)
	if err != nil {
		return nil, fmt.Errorf("mongo db client fetch issue %v ", err)
	}
	err = client.Connect(ctx)
	if err != nil {
		return nil, fmt.Errorf("mongo db client connect issue %v ", err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("mongo db client ping issue %v ", err)
	}
	mconn := &entity.DBConn{}
	mconn.Client = client
	mconn.Config = config
	mconn.Config.Password = ""
	return mconn, nil
}

// Get read collections from db
func (db MongoDB) GetCollection(conn entity.DBConn, collName string) *mongo.Collection {
	coll := conn.Client.Database(conn.Config.DBName).Collection(collName)
	return coll
}

// Insert put data inside db
// func (db MongoDB) GetAll(ctx context.Context, conn entity.DBConn, req entity.DBRequest) error {
// 	collection := db.GetCollection(conn, req.Collection)
// 	collection.Find(ctx, nil)
// 	return err
// }

// Insert put data inside db
func (db MongoDB) Insert(ctx context.Context, conn entity.DBConn, req entity.DBRequest) error {
	collection := db.GetCollection(conn, req.Collection)
	_, err := collection.InsertMany(ctx, req.Data)
	return err
}

// // Update updaate data from db
// func (db MongoDB) Update(ctx context.Context, conn entity.DBConn, req entity.DBRequest) error {
// 	collection := db.GetCollection(conn, req.Collection)
// 	res, err := collection.UpdateOne(ctx, req.Filter, req.Rows[0], req.Opts)
// 	return err
// }

// // Delete delete data from db
func (db MongoDB) Delete(ctx context.Context, conn entity.DBConn, req entity.DBRequest) error {
	collection := db.GetCollection(conn, req.Collection)
	return collection.Drop(ctx)
}

// Close close db connection
func (db MongoDB) Close(ctx context.Context, conn entity.DBConn) error {
	return conn.Client.Disconnect(ctx)
}

func (db MongoDB) getMongoClient(config *entity.DBConfig) (*mongo.Client, error) {
	return mongo.NewClient(options.Client().ApplyURI(db.getURI(config)))
}

func (db MongoDB) getURI(c *entity.DBConfig) string {
	// connString := fmt.Sprintf("mongodb://%v:%v@%v:%v/?authSource=%v",c.UserName, c.Password, c.HostName, c.Port, c.DBName)
	connString := fmt.Sprintf("mongodb://%v:%v/?authSource=%v", c.HostName, c.Port, c.DBName)
	return connString

}
