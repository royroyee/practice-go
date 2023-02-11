package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Data struct {
	Timestamp time.Time
	Value     int
}

func storeDataInDB(session *mgo.Session) {
	// Get a reference to the "data" collection
	collection := session.DB("test").C("data")

	// Insert a new data point
	data := Data{Timestamp: time.Now(), Value: 42}
	err := collection.Insert(data)
	if err != nil {
		fmt.Println("Error inserting data:", err)
		return
	}
	fmt.Println("Data stored successfully")
}

func deleteOldDataFromDB(session *mgo.Session) {
	// Get a reference to the "data" collection
	collection := session.DB("test").C("data")

	// Delete all data points older than 24 hours
	now := time.Now()
	cutoff := now.Add(-24 * time.Hour)
	_, err := collection.RemoveAll(bson.M{"timestamp": bson.M{"$lt": cutoff}})
	if err != nil {
		fmt.Println("Error deleting old data:", err)
		return
	}
	fmt.Println("Old data deleted successfully")
}

func main() {
	// Connect to the MongoDB server
	session, err := mgo.Dial("mongodb://localhost:27017/")
	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
		return
	}
	defer session.Close()

	// Store data in the DB every 12 hours
	ticker := time.NewTicker(12 * time.Hour)
	go func() {
		for range ticker.C {
			storeDataInDB(session.Clone())
		}
	}()

	// Delete old data from the DB every 24 hours
	ticker2 := time.NewTicker(24 * time.Hour)
	go func() {
		for range ticker2.C {
			deleteOldDataFromDB(session.Clone())
		}
	}()

	// Wait indefinitely
	select {}
}
