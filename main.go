package goway_mongodb_store

import (
	"github.com/andrepinto/goway/product"
	"gopkg.in/mgo.v2"
	"log"
	"gopkg.in/mgo.v2/bson"
)

type MongodbRepository struct {
	Session *mgo.Session
	DatabaseName string
}

type MongodbRepositoryOptions struct {
	Url string
	DatabaseName string
}

const (
	PRODUCT_COLLECTIONS = "products"
)

func NewMongodbRepository(options *MongodbRepositoryOptions) *MongodbRepository{

	if(len(options.DatabaseName)==0){
		panic("NO DATABASE")
	}

	session, err := mgo.Dial(options.Url)

	if err != nil {
		panic(err.Error())
	}

	return &MongodbRepository{
		Session: session,
		DatabaseName: options.DatabaseName,
	}
}


func(l *MongodbRepository) Create(){

}

func(l *MongodbRepository) GetAllProducts() []product.Product_v1{

	var products []product.Product_v1

	c := l.Session.DB(l.DatabaseName).C(PRODUCT_COLLECTIONS)

	err := c.Find(bson.M{}).All(&products)

	if err != nil {
		panic(err.Error())
	}

	return products
}

func(l *MongodbRepository) GetAllClients() []product.Client_v1{
	return nil
}

func(l *MongodbRepository) CreateProduct(product *product.Product_v1) (bool, *product.Product_v1){
	return true, nil
}
func(l *MongodbRepository) CreateClient(client *product.Client_v1) (bool, *product.Client_v1){
	return true, nil
}
