package models

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
	"github.com/ltadrian/test-dynamo-db-api/db"
	"github.com/ltadrian/test-dynamo-db-api/forms"
)

type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Gender    string `json:"gender"`
	Age       int    `json:"age"`
	Time      int64  `json:"current_time"`
	UpdatedAt int64  `json:"updated_at"`
}

func (h User) FindUserByID(id string) (*User, error) {
	client := db.GetDB()

	params := &dynamodb.GetItemInput{
		Key: map[string]types.AttributeValue{
			"ID": &types.AttributeValueMemberS{Value: id},
		},
		TableName:      aws.String("User"),
		ConsistentRead: aws.Bool(true),
	}

	resp, err := client.GetItem(context.TODO(), params)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var user *User
	if err := attributevalue.UnmarshalMap(resp.Item, &user); err != nil {
		log.Println(err)
		return nil, err
	}
	return user, nil
}

func (h User) AddUser(userPayload forms.AddUser) (*User, error) {
	client := db.GetDB()
	id := uuid.New()
	user := User{
		ID:        id.String(),
		Name:      userPayload.Name,
		Gender:    userPayload.Gender,
		Age:       userPayload.Age,
		Time:      time.Now().UnixNano(),
		UpdatedAt: time.Now().UnixNano(),
	}

	itemMap, err := attributevalue.MarshalMap(user)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	params := &dynamodb.PutItemInput{
		Item:      itemMap,
		TableName: aws.String("User"),
	}

	output, err := client.PutItem(context.TODO(), params)
	if err != nil {
		return nil, err
	}
	log.Println(output)
	return &user, nil
}

// Example make attribute value map
// itemMap := map[string]types.AttributeValue{
// 	"ID":        &types.AttributeValueMemberS{Value: user.ID},
// 	"Name":      &types.AttributeValueMemberS{Value: user.Name},
// 	"Gender":    &types.AttributeValueMemberS{Value: user.Gender},
// 	"Age":       &types.AttributeValueMemberN{Value: strconv.Itoa(user.Age)},
// 	"Time":      &types.AttributeValueMemberN{Value: strconv.Itoa(int(user.Time))},
// 	"UpdatedAt": &types.AttributeValueMemberN{Value: strconv.Itoa(int(user.UpdatedAt))},
// }
