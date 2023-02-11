package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Name  string
	Phone string
}

func main() {
	// Connect to the MongoDB server
	session, err := mgo.Dial("mongodb://localhost:27017/")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Get a reference to the "people" collection
	collection := session.DB("test").C("people")

	// Create a person
	person := Person{"Ale", "+55 53 8116 9639"}
	err = collection.Insert(person)
	if err != nil {
		panic(err)
	}

	// Read a person
	result := Person{}
	err = collection.Find(bson.M{"name": "Ale"}).One(&result)
	if err != nil {
		panic(err)
	}
	fmt.Println("Phone:", result.Phone)

	// Update a person
	err = collection.Update(bson.M{"name": "Ale"}, bson.M{"$set": bson.M{"phone": "+86 138 1141 5597"}})
	if err != nil {
		panic(err)
	}

	// Delete a person
	err = collection.Remove(bson.M{"name": "Ale"})
	if err != nil {
		panic(err)
	}
}
