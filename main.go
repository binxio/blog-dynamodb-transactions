package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"log"
)

type Cat struct {
	ID      string
	Name    string
	Age     int
	OwnerID string
}

type Owner struct {
	ID     string
	Name   string
	Age    int
	CatIds []string
}

func CreateDynamoDBClient(region string) (*dynamodb.DynamoDB, error) {
	sess, err := session.NewSession()
	if err != nil {
		return nil, err
	}
	svc := dynamodb.New(sess, aws.NewConfig().WithRegion(region))
	return svc, nil
}

func main() {
	owner := Owner{
		ID:     "1",
		Name:   "Dennis",
		Age:    42,
		CatIds: []string{"1", "2"},
	}
	elsa := Cat{
		ID:      "1",
		Name:    "Elsa",
		Age:     16,
		OwnerID: "1",
	}
	tijger := Cat{
		ID:      "2",
		Name:    "Tijger",
		Age:     12,
		OwnerID: "1",
	}
	OwnerAV, err := dynamodbattribute.MarshalMap(owner)
	if err != nil {
		log.Fatal(err)
	}
	ElsaAV, err2 := dynamodbattribute.MarshalMap(elsa)
	if err2 != nil {
		log.Fatal(err)
	}
	TijgerAV, err3 := dynamodbattribute.MarshalMap(tijger)
	if err != nil {
		log.Fatal(err3)
	}
	svc, err := CreateDynamoDBClient("eu-west-1")
	if err != nil {
		log.Fatal(err)
	}
	_, err4 := svc.TransactWriteItems(&dynamodb.TransactWriteItemsInput{
		TransactItems: []*dynamodb.TransactWriteItem{
			{
				Put: &dynamodb.Put{
					TableName: aws.String("Owners"),
					Item:      OwnerAV,
				},
			},
			{
				Put: &dynamodb.Put{
					TableName: aws.String("Cats"),
					Item:      ElsaAV,
				},
			},
			{
				Put: &dynamodb.Put{
					TableName: aws.String("Cats"),
					Item:      TijgerAV,
				},
			},
		},
	})
	if err != nil {
		log.Fatal(err4)
	}
	res, err5 := svc.TransactGetItems(&dynamodb.TransactGetItemsInput{
		TransactItems: []*dynamodb.TransactGetItem{
			{
				Get: &dynamodb.Get{
					TableName: aws.String("Cats"),
					Key: map[string]*dynamodb.AttributeValue{
						"ID": {
							S: aws.String("1"),
						},
					},
				},
			},
			{
				Get: &dynamodb.Get{
					TableName: aws.String("Cats"),
					Key: map[string]*dynamodb.AttributeValue{
						"ID": {
							S: aws.String("2"),
						},
					},
				},
			},
			{
				Get: &dynamodb.Get{
					TableName: aws.String("Owners"),
					Key: map[string]*dynamodb.AttributeValue{
						"ID": {
							S: aws.String("1"),
						},
					},
				},
			},
		},

	})
	if err != nil {
		log.Fatal(err5)
	}

	for _, item := range res.Responses {
		fmt.Print(item.Item)
	}
}
