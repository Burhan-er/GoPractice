package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

func main() {

	type Gender string

	const (
		MALE   Gender = "male"
		FEMALE Gender = "female"
	)
	type Person struct {
		ID     string `json:"id"`
		Name   string `json:"name"`
		Gender Gender`json:"gender"`
		Age    string`json:"age"`
	}

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0})

	ping, err := client.Ping(context.Background()).Result()

	if err != nil {
		fmt.Printf("ping err: %s", err.Error())
	}

	fmt.Printf("ping: %s\n", ping)

	personID := uuid.NewString()

	person := Person{
		ID: personID,
		Name: "David Laid",
		Gender: MALE,
		Age: "18",
	}

	personJson, err := json.Marshal(person)

	if err != nil {
		fmt.Printf("json marshal err: %s\n",err.Error())
	}

	err = client.Set(context.Background(), "person",string(personJson),0).Err()

	if err != nil {
		fmt.Printf("client.Set error: %s\n",err.Error())
	}

	value, err := client.Get(context.Background(),"person").Result()
	if err != nil {
		fmt.Printf("client.Get Error: %s\n",err.Error())
	}

	fmt.Printf("valuefrom redis : %s",value)

}
