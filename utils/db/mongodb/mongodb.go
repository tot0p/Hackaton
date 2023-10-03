package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	client *mongo.Client
}

// NewMongoDB creates a new MongoDB connection
func NewMongoDB(uri string) (*MongoDB, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}
	return &MongoDB{
		client: client,
	}, nil
}

// NewDatabase creates a new MongoDB database
func (m *MongoDB) NewDatabase(db string) *mongo.Database {
	return m.client.Database(db)
}

// NewCollection creates a new MongoDB collection
func (m *MongoDB) NewCollection(db, collection string) *mongo.Collection {
	return m.client.Database(db).Collection(collection)
}

// GetClient returns the MongoDB client
func (m *MongoDB) GetClient() *mongo.Client {
	return m.client
}

// GetCollection returns a MongoDB collection
func (m *MongoDB) GetCollection(db, collection string) *mongo.Collection {
	return m.client.Database(db).Collection(collection)
}

// GetAllData returns all data from a MongoDB collection
func (m *MongoDB) GetAllData(db, collection string) (*mongo.Cursor, error) {
	return m.GetCollection(db, collection).Find(context.Background(), nil)
}

// Disconnect closes the MongoDB connection
func (m *MongoDB) Disconnect() error {
	return m.client.Disconnect(context.Background())
}

// Ping pings the MongoDB connection
func (m *MongoDB) Ping() error {
	return m.client.Ping(context.Background(), nil)
}

// InsertOne inserts one document into a MongoDB collection
func (m *MongoDB) InsertOne(db, collection string, document interface{}) (*mongo.InsertOneResult, error) {
	return m.GetCollection(db, collection).InsertOne(context.Background(), document)
}

// InsertMany inserts many documents into a MongoDB collection
func (m *MongoDB) InsertMany(db, collection string, documents []interface{}) (*mongo.InsertManyResult, error) {
	return m.GetCollection(db, collection).InsertMany(context.Background(), documents)
}

// FindOne finds one document in a MongoDB collection
func (m *MongoDB) FindOne(db, collection string, filter interface{}) *mongo.SingleResult {
	return m.GetCollection(db, collection).FindOne(context.Background(), filter)
}

// Find finds documents in a MongoDB collection
func (m *MongoDB) Find(db, collection string, filter interface{}) (*mongo.Cursor, error) {
	return m.GetCollection(db, collection).Find(context.Background(), filter)
}

// UpdateOne updates one document in a MongoDB collection
func (m *MongoDB) UpdateOne(db, collection string, filter, update interface{}) (*mongo.UpdateResult, error) {
	return m.GetCollection(db, collection).UpdateOne(context.Background(), filter, update)
}

// UpdateMany updates many documents in a MongoDB collection
func (m *MongoDB) UpdateMany(db, collection string, filter, update interface{}) (*mongo.UpdateResult, error) {
	return m.GetCollection(db, collection).UpdateMany(context.Background(), filter, update)
}

// DeleteOne deletes one document in a MongoDB collection
func (m *MongoDB) DeleteOne(db, collection string, filter interface{}) (*mongo.DeleteResult, error) {
	return m.GetCollection(db, collection).DeleteOne(context.Background(), filter)
}

// DeleteMany deletes many documents in a MongoDB collection
func (m *MongoDB) DeleteMany(db, collection string, filter interface{}) (*mongo.DeleteResult, error) {
	return m.GetCollection(db, collection).DeleteMany(context.Background(), filter)
}

// CountDocuments counts documents in a MongoDB collection
func (m *MongoDB) CountDocuments(db, collection string, filter interface{}) (int64, error) {
	return m.GetCollection(db, collection).CountDocuments(context.Background(), filter)
}

// Aggregate aggregates documents in a MongoDB collection
func (m *MongoDB) Aggregate(db, collection string, pipeline interface{}) (*mongo.Cursor, error) {
	return m.GetCollection(db, collection).Aggregate(context.Background(), pipeline)
}

// Distinct finds distinct values for a field in a MongoDB collection
func (m *MongoDB) Distinct(db, collection, fieldName string, filter interface{}) ([]interface{}, error) {
	return m.GetCollection(db, collection).Distinct(context.Background(), fieldName, filter)
}

// BulkWrite performs a bulk write in a MongoDB collection
func (m *MongoDB) BulkWrite(db, collection string, models []mongo.WriteModel) (*mongo.BulkWriteResult, error) {
	return m.GetCollection(db, collection).BulkWrite(context.Background(), models)
}

// Watch watches a MongoDB collection
func (m *MongoDB) Watch(db, collection string, pipeline interface{}) (*mongo.ChangeStream, error) {
	return m.GetCollection(db, collection).Watch(context.Background(), pipeline)
}

// Indexes returns the indexes of a MongoDB collection
func (m *MongoDB) Indexes(db, collection string) mongo.IndexView {
	return m.GetCollection(db, collection).Indexes()
}
