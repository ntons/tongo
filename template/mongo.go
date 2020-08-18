package template

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"
)

type M map[string]interface{}

func FromMongo(
	ctx context.Context, dburi string, collections ...string) (
	_ M, err error) {
	// get database name from uri
	var cs connstring.ConnString
	if cs, err = connstring.ParseAndValidate(dburi); err != nil {
		return
	}
	database := cs.Database
	// connect mongodb
	cli, err := mongo.NewClient(options.Client().ApplyURI(dburi))
	if err != nil {
		return
	}
	if err = cli.Connect(ctx); err != nil {
		return
	}
	db := cli.Database(database)
	// query collection names
	fmt.Println(len(collections))
	if len(collections) == 0 {
		if collections, err = db.ListCollectionNames(ctx, bson.D{}); err != nil {
			return
		}
	}
	fmt.Println(len(collections))
	// query documents
	tables := M{}
	for _, name := range collections {
		var cur *mongo.Cursor
		if cur, err = db.Collection(name).Find(ctx, bson.D{}); err != nil {
			return
		}
		table := []M{}
		for cur.Next(ctx) {
			doc := M{}
			if err = cur.Decode(&doc); err != nil {
				return
			}
			table = append(table, doc)
		}
		tables[name] = table
	}
	return tables, nil
}
