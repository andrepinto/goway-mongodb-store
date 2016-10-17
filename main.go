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
	prd := []product.Product_v1{

		product.Product_v1{
			Code: "customer",
			Name: "Customer Api",
			Version: "1",
			Routes: []product.Routes_v1{
				product.Routes_v1{
					ListenPath: "/api/facets",
					Verb: "GET",
					ServiceName: "authentication",
					Handlers: []string{"AUTHENTICATION", "METRICS"},
					Code:"auth_by_email",
					Roles: []string{},
					InjectData: []product.InjectData_v1{
						product.InjectData_v1{
							Where:"params",
							Code: "paramId",
							Value: "123456789",
						},
						product.InjectData_v1{
							Where:"header",
							Code: "headerId",
							Value: "9999999999",
						},
					},
					InjectGlobalData:true,
				},
				product.Routes_v1{
					ListenPath: "/api/facets",
					Verb: "POST",
					ServiceName: "authentication",
					Handlers: []string{"AUTHENTICATION", "METRICS"},
					Code:"auth_by_email",
					Roles: []string{},
					InjectData: []product.InjectData_v1{
						product.InjectData_v1{
							Where:"params",
							Code: "paramId",
							Value: "123456789",
						},
						product.InjectData_v1{
							Where:"header",
							Code: "headerId",
							Value: "9999999999",
						},
					},
					InjectGlobalData:true,
				},
			},
		},
		product.Product_v1{
			Code: "cockpit",
			Name: "cockpit Api",
			Version: "1",

			Routes: []product.Routes_v1{
				product.Routes_v1{
					ListenPath: "/auth/byemail",
					Verb: "GET",
					ServiceName: "authentication",
					Handlers: []string{"AUTHENTICATION", "METRICS"},
					Code:"auth_by_email",
					Roles: []string{},
				},
				product.Routes_v1{
					ListenPath: "/auth/renew",
					Verb: "GET",
					ServiceName: "authentication",
					Handlers:  []string{"AUTHENTICATION", "METRICS"},
					Code:"renew_token",
					Roles: []string{},
				},
			},
		},

	}


	c := l.Session.DB(l.DatabaseName).C(PRODUCT_COLLECTIONS)

	for _, k:= range prd{
		_ , err := c.UpsertId( k.Code, k )
		if err != nil {
			log.Println("Error creating Profile: ", err.Error())
		}
	}


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