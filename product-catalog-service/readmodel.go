//mogo db read model
package main

import "go.mongodb.org/mongo-driver/mongo"

type readModel interface {
}

type MongoReadModel struct {
	collection *mongo.Collection
}
