# Mongofunc

mongofunc is a Go package that provides a fluent query builder for MongoDB. It allows constructing MongoDB find queries using a chained syntax.

The library aims to make querying MongoDB easy and intuitive for Go developers.

## Table of Contents
- [Introduction](#introduction)
- [Getting Started](#getting-started)
- [Moper: Building Queries](#moper-building-queries)
  - [Query](#query)
  - [Update](#update)
- [Mocom: MongoClient Commands](#mocom-mongoclient-commands)
  - [Connect with Mocom](#connect-with-mocom)
  - [Working with Collections](#working-with-collections) 
  - [Using Commands](#using-commands)
- [Transaction](#transaction)
  - [Nested Transaction](#nested-transaction)
  - [Tx()](#tx)
  - [TxOptimal()](#txoptimal)
- [License](#license)

## Content:


### Introduction

With mongofunc, you can write MongoDB find queries using chained method calls instead of constructing nested documents. This results in query code that is:

- Clean - Query logic reads clearly from top to bottom
- Composable - Combine and reuse query conditions easily
- Typed - Fully integrated with Go's type system

For example:

```go
query := moper.Query().
           Greater("age", 21).
           Equal("status", "active")
```

The above query filters for age greater than 21 and status equal to "active".

mongofunc handles converting the chained calls to proper MongoDB query filter documents under the hood.

### Getting Started

Install the package:

```sh
$ go get github.com/func25/mongofunc/v2
```

Import it in your code:

```go
import "github.com/func25/mongofunc/v2/moper"
```

### Moper: Building Queries

#### Query 

To start a new query:

```go
query := moper.Query()
```

Start chaining query methods to build your query, let's find all users name Sarah and age greater than 18:

```go
filter := query.Equal("name", "Sarah").
                Greater("age", 18)
```

The package provides functions for MongoDB queries to perform time-based comparisons on documents.

```go
filter := moper.Query().
                EqualGreaterTime("expired", time.Now())
```

#### Update

Let's take a complex update command as an example:

```go
update := moper.Query().
	Inc("age", 1).
	Push("hobbies", "guitar").
	Unset("addresses.0").
	Set("addresses.1", bson.M{
		"street": "123 Main St",
		"city":   "Springfield",
	})
```

Support simple aggregation:

```go
intArr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

matchStage := moper.Query().MatchD(moper.Query().InArray("damage", intArr))
groupStage := moper.Query().Group(
	moper.P{K: "_id", V: nil},
	moper.P{K: "total", V: moper.Query().Sum("damage")},
)

req := &mocom.AggregationRequest[Hero]{
	Pipeline: []moper.D{matchStage, groupStage},
	Options:  []*options.AggregateOptions{},
}
```

Let's skim through some query methods:

- Equal(fieldName, value): Adds an equality filter to the pipeline.
- NotEqual(fieldName, value): Adds a non-equality filter to the pipeline.
- InEll(fieldName, value...): Adds an "in" filter to the pipeline.
- NotInEll(fieldName, value...): Adds a "not in" filter to the pipeline.
- EqualLess(fieldName, value): Adds a "less than or equal" filter to the pipeline.
- EqualGreater(fieldName, value): Adds a "greater than or equal" filter to the pipeline.
- Less(fieldName, value): Adds a "less than" filter to the pipeline.
- Greater(fieldName, value): Adds a "greater than" filter to the pipeline.
- Exists(fieldName, exist): Adds an "exists" filter to the pipeline.
- Or(filters...): Adds an "or" filter to the pipeline
- ...

### Mocom: MongoClient Commands

mocom is a Go package that provides a simple MongoDB client wrapper to handle MongoDB connections and collections. It offers convenient functions to connect to a MongoDB server, set the default database, retrieve the client, and get collections for reading and writing.

#### Connect with mocom

Use the Connect function to establish a connection to the MongoDB server:

```go
ctx := context.Background()
uri := "mongodb://localhost:27017/?w=majority&retryWrites=false"
dbName := "your-database-name"

if err := mocom.Connect(ctx, uri, dbName); err != nil {
    log.Fatal("Failed to connect to MongoDB:", err)
}

// Now you can perform database operations using the mocom and moper package
```

#### Working with Collections

To use mocom command, your models should meet the mocom.Model interface{}

```go
// Get a collection for reading with read preference "Nearest"
readColl := mocom.CollRead("your-collection-name")

// Get a collection for writing with read preference "Primary"
writeColl := mocom.CollWrite("your-collection-name")
```

### Using commands

Mocom supports CRUD operations with moper commands, let's check a fulll example:

```go
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/func25/mongofunc/v2/mocom"
	"github.com/func25/mongofunc/v2/moper"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	mocom.ObjectID `json:",inline" bson:",inline"`
	Username       string    `json:"username" bson:"username"`
	Email          string    `json:"email" bson:"email"`
	Age            int       `json:"age" bson:"age"`
	CreatedAt      time.Time `json:"createdAt" bson:"createdAt"`
}

// This function is needed to satisfy the mocom.Model interface, so we can use it with mocom generic functions
func (u User) CollName() string {
	return "users"
}

func main() {
	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Second)
	defer cancel()

	uri := "mongodb://localhost:27017/?w=majority&retryWrites=false"
	dbName := "exampledb"

	err := mocom.Connect(ctx, uri, dbName)
	if err != nil {
		fmt.Println("Failed to connect to MongoDB:", err)
		return
	}

	// Perform CRUD operations with the User model
	user := User{
		Username:  "john_doe",
		Email:     "john@example.com",
		Age:       30,
		CreatedAt: time.Now(),
	}

	// Create a new user using mocom.Create[User] syntax
	insertedID, err := mocom.CreateT[User](ctx, user)
	if err != nil {
		fmt.Println("Failed to create user:", err)
		return
	}

	fmt.Println("User created with ID:", insertedID)

	// Find all users using moper package
	matchStage := moper.Query().MatchD(moper.Query().InArray("age", []int{30, 40}))
	groupStage := moper.Query().Group(moper.P{K: "_id", V: nil}, moper.P{K: "total", V: moper.Query().Sum("age")})

	// Aggregate the users with matching age and calculate the total age using mocom
	aggRequest := &mocom.AggregationRequest{
		Pipeline: []moper.D{matchStage, groupStage},
		Options:  []*options.AggregateOptions{},
	}

	aggResult, err := mocom.AggregateT[User](ctx, aggRequest)
	if err != nil {
		fmt.Println("Failed to aggregate users:", err)
		return
	}

	fmt.Println("Users aggregated:", aggResult)

	// Update the user's email using moper package
	filter := moper.Query().Equal("_id", insertedID)
	update := moper.Query().Set(moper.Pr("email", "john.doe@example.com"))
	_, err = mocom.UpdateOneT[User](ctx, filter, update)
	if err != nil {
		fmt.Println("Failed to update user:", err)
		return
	}

	fmt.Println("User updated")

	// Find a user by ID using mocom
	foundUser, err := mocom.FindOne[User](ctx, filter)
	if err != nil {
		fmt.Println("Failed to find user by ID:", err)
		return
	}

	fmt.Println("User found by ID:", foundUser)

	// Delete the user using mocom
	_, err = mocom.DeleteOneT[User](ctx, filter)
	if err != nil {
		fmt.Println("Failed to delete user:", err)
		return
	}

	fmt.Println("User deleted")
}
```

### Transaction

#### Nested transaction

mocom provides support for nested transactions to help you manage complex transactional workflows. If a nested transaction is detected, the Tx and TxOptimal functions will execute the transaction with the provided context.

#### Tx()

The Tx function executes a MongoDB transaction based on the provided configuration (cfg) and context (ctx). It allows you to perform a series of operations in a transactional manner, ensuring that either all operations succeed, or none of them are applied.

```go
result, err := mocom.Tx(ctx, TransactionConfig{
    SessConfig: options.Session(),
    Options:    options.Transaction(),
    Func: func(ctx mongo.SessionContext) (interface{}, error) {
        // Your transaction logic here
        // e.g., insert, update, or delete operations
        // e.g., perform multiple actions as a single atomic transaction
    },
})
```

### TxOptimal()

The TxOptimal function provides an optimal MongoDB transaction without read operations. It is suitable for scenarios where you only need to perform write operations, ensuring atomicity and consistency across multiple writes.

```go
result, err := mocom.TxOptimal(ctx, func(ctx mongo.SessionContext) (interface{}, error) {
    // Your transaction logic here
    // e.g., insert, update, or delete operations
    // e.g., write operations without read operations
})
```

### License

mongofunc is provided under the MIT License. Use it freely in your Go applications.

Let me know if you would like any changes to the README. I'm happy to help make it as useful as possible!