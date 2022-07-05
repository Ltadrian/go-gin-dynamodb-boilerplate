package db

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	localConfig "github.com/ltadrian/go-gin-dynamodb-boilerplate/config"
)

var db *dynamodb.Client

func Init() {
	c := localConfig.GetConfig()
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion(c.GetString("db.region")),
		config.WithEndpointResolverWithOptions(
			aws.EndpointResolverWithOptionsFunc(
				func(_, _ string, _ ...interface{}) (aws.Endpoint, error) {
					return aws.Endpoint{
						URL: c.GetString("db.endpoint"),
					}, nil
				}),
		),
	)

	if err != nil {
		panic(err)
	}

	db = dynamodb.NewFromConfig(cfg)

	tables, err := db.ListTables(context.TODO(), &dynamodb.ListTablesInput{})
	if err != nil {
		fmt.Printf("err: %v", err)
	} else {
		fmt.Printf("dynamodb tables: %v\n", tables.TableNames)
	}
}

func GetDB() *dynamodb.Client {
	return db
}
